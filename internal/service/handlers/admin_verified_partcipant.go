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
	"github.com/recovery-flow/organization-storage/internal/service/responses"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AdminVerifiedParticipant(w http.ResponseWriter, r *http.Request) {
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
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"organization_id": validation.NewError("organization_id", "Invalid organization id"),
		})...)
		return
	}

	userId, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse user id")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"user_id": validation.NewError("user_id", "Invalid user id"),
		})...)
		return
	}

	value := chi.URLParam(r, "value")
	if value != "true" && value != "false" {
		log.Error("Failed to parse organization verification value")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"value": validation.NewError("value", "Invalid value"),
		})...)
		return
	}

	filters := make(map[string]any)
	filters["_id"] = orgId

	participants, err := server.MongoDB.Organization.New().Filter(map[string]any{"_id": orgId}).
		Participants().Filter(map[string]any{"user_id": userId}).
		UpdateOne(r.Context(), map[string]any{"verified": value})
	if err != nil {
		log.WithError(err).Error("Failed to update organization")
		httpkit.RenderErr(w, problems.InternalError("Failed to update organization"))
		return
	}

	httpkit.Render(w, responses.Participant(*participants))
}
