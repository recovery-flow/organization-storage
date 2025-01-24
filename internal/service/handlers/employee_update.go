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
	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"github.com/recovery-flow/organization-storage/internal/service/requests"
	"github.com/recovery-flow/organization-storage/internal/service/responses"
	"github.com/recovery-flow/roles"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EmployeeUpdate(w http.ResponseWriter, r *http.Request) {
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
	updatedId, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse request: ")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	orgId, err := primitive.ObjectIDFromHex(chi.URLParam(r, "organization_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse organization id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	filtersForOrg := make(map[string]any)
	filtersForOrg["_id"] = orgId

	organization, err := server.MongoDB.Organization.New().Filter(filtersForOrg).Get(r.Context())
	if err != nil {
		log.WithError(err).Error("Failed to get organization")
		httpkit.RenderErr(w, problems.InternalError("Failed to get organization"))
		return
	}

	stmt := make(map[string]any)
	if req.Data.Attributes.FirstName != nil {
		stmt["first_name"] = req.Data.Attributes.FirstName
		stmt["verified"] = false
	}
	if req.Data.Attributes.SecondName != nil {
		stmt["second_name"] = req.Data.Attributes.SecondName
		stmt["verified"] = false
	}
	if req.Data.Attributes.ThirdName != nil {
		stmt["third_name"] = req.Data.Attributes.ThirdName
		stmt["verified"] = false
	}
	if req.Data.Attributes.DisplayName != nil {
		stmt["display_name"] = req.Data.Attributes.DisplayName
	}
	if req.Data.Attributes.Position != nil {
		stmt["position"] = req.Data.Attributes.Position
	}
	if req.Data.Attributes.Desc != nil {
		stmt["desc"] = req.Data.Attributes.Desc
	}

	role := req.Data.Attributes.Role
	if role != nil {
		newEmpRole, err := roles.StringToRoleOrg(*role)
		if err != nil {
			log.WithError(err).Error("Failed to parse role")
			httpkit.RenderErr(w, problems.BadRequest(err)...)
			return
		}
		var owner models.Employee
		for _, emp := range organization.Employees {
			if emp.UserID == initiatorId {
				if roles.CompareRolesOrg(emp.Role, roles.RoleOrgModer) < 0 ||
					roles.CompareRolesOrg(emp.Role, newEmpRole) < 0 ||
					newEmpRole == roles.RoleOrgOwner {
					log.Error("User haven't enough rights to create employee")
					httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
					return
				}
				if emp.Role == roles.RoleOrgOwner {
					owner = emp
				}
				break
			}
		}
		if owner.UserID == updatedId {
			if newEmpRole != roles.RoleOrgOwner {
				log.Error("Owner can't change his role")
				httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
				return
			}
		}
		stmt["role"] = newEmpRole
	} else {
		for _, emp := range organization.Employees {
			if emp.UserID == initiatorId {
				if roles.CompareRolesOrg(emp.Role, roles.RoleOrgModer) < 0 {
					log.Error("User haven't enough rights to create employee")
					httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
					return
				}
				break
			}
		}
	}

	filterForEmp := make(map[string]any)
	filterForEmp["user_id"] = updatedId

	err = server.MongoDB.Organization.New().Filter(filtersForOrg).
		Employees().
		Filter(filterForEmp).
		UpdateOne(r.Context(), stmt)
	if err != nil {
		log.WithError(err).Error("Failed to update employee")
		httpkit.RenderErr(w, problems.InternalError("Failed to update employee"))
		return
	}

	res, err := server.MongoDB.Organization.New().Filter(filtersForOrg).Employees().Filter(filterForEmp).Get(r.Context())
	if err != nil {
		log.WithError(err).Error("Failed to get updated employee")
		httpkit.RenderErr(w, problems.InternalError("Failed to get updated employee"))
		return
	}

	httpkit.Render(w, responses.Employee(*res))
}
