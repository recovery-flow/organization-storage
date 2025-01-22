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

func OrganizationLinksUpdate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewOrganizationLinksUpdate(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	twitter := req.Data.Attributes.Twitter
	instagram := req.Data.Attributes.Instagram
	facebook := req.Data.Attributes.Facebook
	tiktok := req.Data.Attributes.Tiktok
	linkedin := req.Data.Attributes.Linkedin
	telegram := req.Data.Attributes.Telegram
	discord := req.Data.Attributes.Discord
	website := req.Data.Attributes.Website

	initiatorId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		server.Logger.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	orgId, err := primitive.ObjectIDFromHex(req.Data.Id)
	if err != nil {
		log.WithError(err).Error("Failed to parse organization id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	filters := make(map[string]any)
	filters["_id"] = orgId

	organization, err := server.MongoDB.Organization.Filter(filters).Get(r.Context())
	if err != nil {
		log.WithError(err).Error("Failed to get organization")
		httpkit.RenderErr(w, problems.InternalError("Failed to get organization"))
		return
	}

	for _, emp := range organization.Employees {
		if emp.UserID == initiatorId {
			if roles.CompareRolesOrg(emp.Role, roles.RoleOrgModer) > -1 {
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

	err = server.MongoDB.Organization.Links().UpdateOne(r.Context(), map[string]any{
		"twitter":   twitter,
		"instagram": instagram,
		"facebook":  facebook,
		"tiktok":    tiktok,
		"linkedin":  linkedin,
		"telegram":  telegram,
		"discord":   discord,
		"website":   website,
	})

	if err != nil {
		log.WithError(err).Error("Failed to update organization links")
		httpkit.RenderErr(w, problems.InternalError("Failed to update organization links"))
		return
	}

	organization, err = server.MongoDB.Organization.Filter(filters).Get(r.Context())
	if err != nil {
		log.WithError(err).Error("Failed to get organization")
		httpkit.RenderErr(w, problems.InternalError("Failed to get organization"))
		return
	}

	httpkit.Render(w, responses.Organization(*organization))
}
