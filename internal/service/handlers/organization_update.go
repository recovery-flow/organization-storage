package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/organization-storage/internal/config"
	"github.com/recovery-flow/organization-storage/internal/service/requests"
	"github.com/recovery-flow/organization-storage/internal/service/responses"
	"github.com/recovery-flow/roles"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func OrganizationUpdate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewOrganizationUpdate(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if chi.URLParam(r, "organization_id") != req.Data.Id {
		log.Error("Request id does not match URL id")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"id": validation.NewError("id", "Request id does not match URL id"),
		})...)
		return
	}

	orgId, err := primitive.ObjectIDFromHex(req.Data.Id)
	if err != nil {
		log.WithError(err).Error("Failed to parse organization id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	initiatorId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	filters := make(map[string]any)
	filters["_id"] = orgId

	participant, err := server.MongoDB.Organization.New().Filter(filters).Participants().Filter(map[string]any{
		"user_id": initiatorId,
	}).Get(r.Context())
	if err != nil {
		log.WithError(err).Error("Failed to update initiative")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	if roles.CompareRolesOrg(participant.Role, roles.RoleOrgModer) <= 0 {
		log.Error("User has no rights to update initiative")
		httpkit.RenderErr(w, problems.Forbidden("User has no rights to update initiative"))
		return
	}

	stmt := make(map[string]any)
	if req.Data.Attributes.Name != nil {
		stmt["name"] = *req.Data.Attributes.Name
		stmt["verified"] = false
	}
	if req.Data.Attributes.Sort != nil {
		stmt["sort"] = *req.Data.Attributes.Sort
		stmt["verified"] = false
	}
	if req.Data.Attributes.Country != nil {
		stmt["country"] = *req.Data.Attributes.Country
		stmt["verified"] = false
	}
	if req.Data.Attributes.City != nil {
		stmt["city"] = *req.Data.Attributes.City
		stmt["verified"] = false
	}
	if req.Data.Attributes.Desc != nil {
		stmt["desc"] = *req.Data.Attributes.Desc
	}

	res, err := server.MongoDB.Organization.New().Filter(filters).UpdateOne(r.Context(), stmt)
	if err != nil {
		log.WithError(err).Error("Failed to update organization")
		httpkit.RenderErr(w, problems.InternalError("Failed to update organization"))
		return
	}

	log.Infof("Organization updated %s", initiatorId)
	httpkit.Render(w, responses.Organization(*res))
}
