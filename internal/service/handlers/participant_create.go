package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/organization-storage/internal/config"
	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"github.com/recovery-flow/organization-storage/internal/service/requests"
	"github.com/recovery-flow/organization-storage/internal/service/responses"
	"github.com/recovery-flow/roles"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ParticipantCreate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.ParticipantCreate(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if chi.URLParam(r, "organization_id") != req.Data.Attributes.OrgId {
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("organization_id does not match request organization_id"))...)
		return
	}

	participantIdStr := req.Data.Attributes.UserId
	firstName := req.Data.Attributes.FirstName
	secondName := req.Data.Attributes.SecondName
	thirdName := req.Data.Attributes.ThirdName
	displayName := req.Data.Attributes.DisplayName
	position := req.Data.Attributes.Position
	desc := req.Data.Attributes.Desc
	role := req.Data.Attributes.Role

	participantId, err := uuid.Parse(participantIdStr)
	if err != nil {
		log.WithError(err).Error("Failed to parse participant id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	orgId, err := primitive.ObjectIDFromHex(chi.URLParam(r, "organization_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse organization id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	initiatorId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		server.Logger.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	filters := make(map[string]any)
	filters["_id"] = orgId

	organization, err := server.MongoDB.Organization.New().Filter(filters).Get(r.Context())
	if err != nil {
		log.WithError(err).Error("Failed to get organization")
		httpkit.RenderErr(w, problems.InternalError("Failed to get organization"))
		return
	}

	newEmpRole, err := roles.StringToRoleOrg(role)
	if err != nil {
		log.WithError(err).Error("Failed to parse role")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	for _, emp := range organization.Participants {
		if emp.UserID == initiatorId {
			if roles.CompareRolesOrg(emp.Role, roles.RoleOrgModer) < 0 ||
				roles.CompareRolesOrg(emp.Role, newEmpRole) < 0 ||
				newEmpRole == roles.RoleOrgOwner {
				log.Error("User haven't enough rights to create participant")
				httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
				return
			}
			break
		}
	}

	for _, emp := range organization.Participants {
		log.Infof("emp.UserID: %v, participantId: %v", emp.UserID, participantId)
		if emp.UserID == participantId {
			log.WithError(err).Error("Failed to find initiator user")
			httpkit.RenderErr(w, problems.Conflict("User already exists in organization"))
			return
		}
	}

	participant, err := server.MongoDB.Organization.New().Filter(filters).
		Participants().Create(r.Context(), models.Participant{
		UserID:      participantId,
		FirstName:   firstName,
		SecondName:  secondName,
		ThirdName:   thirdName,
		DisplayName: displayName,
		Position:    position,
		Desc:        desc,
		Verified:    false,
		Role:        newEmpRole,

		CreatedAt: time.Now().UTC(),
	})

	if err != nil {
		log.WithError(err).Error("Failed to insert participant")
		httpkit.RenderErr(w, problems.InternalError("Failed to insert participant"))
		return
	}

	log.Infof("Create participant test")

	httpkit.Render(w, responses.Participant(*participant))
}
