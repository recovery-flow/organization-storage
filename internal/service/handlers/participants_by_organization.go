package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/organization-storage/internal/config"
	"github.com/recovery-flow/organization-storage/internal/service/responses"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ParticipantsByOrganization(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	orgId, err := primitive.ObjectIDFromHex(chi.URLParam(r, "organization_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse organization id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	filters := make(map[string]any)
	filters["_id"] = orgId

	participants, err := server.MongoDB.Organization.New().Filter(filters).Participants().Select(r.Context())
	if err != nil {
		log.WithError(err).Error("Failed to get participants")
		httpkit.RenderErr(w, problems.InternalError("Failed to get participants"))
		return
	}

	httpkit.Render(w, responses.ParticipantCollection(participants))
}
