package handlers

import (
	"net/http"

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

	initiatorId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	name := req.Data.Attributes.Name
	desc := req.Data.Attributes.Desc
	sort := req.Data.Attributes.Sort
	country := req.Data.Attributes.Country
	city := req.Data.Attributes.City

	orgId, err := primitive.ObjectIDFromHex(req.Data.Id)
	if err != nil {
		log.WithError(err).Error("Failed to parse organization id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
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

	for _, emp := range organization.Participants {
		if emp.UserID == initiatorId {
			if roles.CompareRolesOrg(emp.Role, roles.RoleOrgAdmin) < 0 {
				err = roles.ErrorNoPermission
			}
			break
		}
	}

	if err != nil {
		log.WithError(err).Error("User does not have permission to update organization")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	stmt := make(map[string]any)
	if name != nil {
		stmt["name"] = *name
		stmt["verified"] = false
	}
	if sort != nil {
		stmt["sort"] = *sort
		stmt["verified"] = false
	}
	if country != nil {
		stmt["country"] = *country
		stmt["verified"] = false
	}
	if city != nil {
		stmt["city"] = *city
		stmt["verified"] = false
	}
	if desc != nil {
		stmt["desc"] = *desc
	}

	res, err := server.MongoDB.Organization.New().Filter(filters).UpdateOne(r.Context(), stmt)
	if err != nil {
		log.WithError(err).Error("Failed to update organization")
		httpkit.RenderErr(w, problems.InternalError("Failed to update organization"))
		return
	}

	log.Infof("Organization update test")

	log.Infof("Organization updated %s", initiatorId)
	httpkit.Render(w, responses.Organization(*res))
}
