package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/organization-storage/internal/config"
	"github.com/recovery-flow/organization-storage/internal/service/responses"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EmployeeByUserID(w http.ResponseWriter, r *http.Request) {
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

	userId, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse user id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	filtersOrg := make(map[string]any)
	filtersOrg["_id"] = orgId

	res, err := server.MongoDB.Organization.New().Filter(filtersOrg).Employees().
		Filter(map[string]any{
			"user_id": userId,
		}).
		Select(r.Context())
	if err != nil {
		log.WithError(err).Error("Failed to update organization")
		httpkit.RenderErr(w, problems.InternalError("Failed to update organization"))
		return
	}

	httpkit.Render(w, responses.Employee(res[0]))
}
