package handlers

import (
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/organization-storage/internal/config"
	"github.com/recovery-flow/organization-storage/internal/service/requests"
	"github.com/recovery-flow/roles"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func OrganizationUploadLogo(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewUploadImage(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	defer req.File.Close()

	initiatorId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		server.Logger.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	orgId, err := primitive.ObjectIDFromHex(chi.URLParam(r, "organization_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse organization id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
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
			if roles.CompareRolesTeam(roles.TeamRole(emp.Role), roles.RoleTeamAdmin) > -1 {
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

	yes := true
	uploadParams := uploader.UploadParams{
		Folder:       "organization/logo",
		PublicID:     organization.ID.String(),
		Overwrite:    &yes,
		ResourceType: "image",
	}
	uploadResult, err := server.Storage.Upload.Upload(r.Context(), req.File, uploadParams)
	if err != nil {
		log.Errorf("Failed to upload avatar to Cloudinary: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to upload avatar"))
		return
	}

	stmt := map[string]any{
		"avatar": uploadResult.SecureURL,
	}

	_, err = server.MongoDB.Organization.FilterById(orgId).UpdateOne(r.Context(), stmt)
	if err != nil {
		log.WithError(err).Error("Failed to update organization")
		httpkit.RenderErr(w, problems.InternalError("Failed to update organization"))
		return
	}

	httpkit.Render(w, http.StatusAccepted)
}
