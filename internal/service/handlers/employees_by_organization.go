package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/organization-storage/internal/config"
	"github.com/recovery-flow/organization-storage/internal/service/requests"
	"github.com/recovery-flow/organization-storage/internal/service/responses"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EmployeesByOrganization(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewEmployeeUpdate(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	initiatorId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		server.Logger.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	if chi.URLParam(r, "organization_id") != req.Data.Attributes.OrgId {
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("organization_id does not match request organization_id"))...)
		return
	}

	if chi.URLParam(r, "user_id") != req.Data.Id {
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("user_id does not match request user_id"))...)
		return
	}

	orgId, err := primitive.ObjectIDFromHex(chi.URLParam(r, "organization_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse organization id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	employees, err := server.MongoDB.Organization.FilterById(orgId).Employees().Select(r.Context())
	if err != nil {
		log.WithError(err).Error("Failed to get employees")
		httpkit.RenderErr(w, problems.InternalError("Failed to get employees"))
		return
	}

	httpkit.Render(w, responses.EmployeesCollection(employees))

}
