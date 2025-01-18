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

func EmployeeCreate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewEmployeeCreate(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if chi.URLParam(r, "organization_id") != req.Data.Attributes.OrgId {
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("organization_id does not match request organization_id"))...)
		return
	}
	employeeIdStr := req.Data.Attributes.UserId
	firstName := req.Data.Attributes.FirstName
	secondName := req.Data.Attributes.SecondName
	thirdName := req.Data.Attributes.ThirdName
	displayName := req.Data.Attributes.DisplayName
	position := req.Data.Attributes.Position
	desc := req.Data.Attributes.Desc
	role := req.Data.Attributes.Role

	employeeId, err := uuid.Parse(employeeIdStr)
	if err != nil {
		log.WithError(err).Error("Failed to parse employee id")
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

	organization, err := server.MongoDB.Organization.FilterById(orgId).Get(r.Context())
	if err != nil {
		log.WithError(err).Error("Failed to get organization")
		httpkit.RenderErr(w, problems.InternalError("Failed to get organization"))
		return
	}

	for _, emp := range organization.Employees {
		if emp.UserID == initiatorId {
			if roles.CompareRolesOrg(emp.Role, roles.RoleOrgModer) > -1 &&
				roles.CompareRolesOrg(emp.Role, roles.OrgRole(role)) > -1 &&
				roles.OrgRole(role) != roles.RoleOrgOwner {
				err = roles.ErrorNoPermission
			}
			break
		}
	}

	if err != nil {
		log.WithError(err).Error("Failed to find initiator user")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	employee, err := server.MongoDB.Organization.FilterById(orgId).Employees().Insert(r.Context(), models.Employee{
		UserID:      employeeId,
		FirstName:   firstName,
		SecondName:  secondName,
		ThirdName:   thirdName,
		DisplayName: displayName,
		Position:    position,
		Desc:        desc,
		Verified:    false,
		Role:        roles.OrgRole(role),

		CreatedAt: time.Now(),
	})
	if err != nil {
		log.WithError(err).Error("Failed to insert employee")
		httpkit.RenderErr(w, problems.InternalError("Failed to insert employee"))
		return
	}

	httpkit.Render(w, responses.Employee(*employee))
}
