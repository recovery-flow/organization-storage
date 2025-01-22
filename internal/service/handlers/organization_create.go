package handlers

import (
	"net/http"
	"strings"

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
)

func OrganizationCreate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewOrganizationCreate(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	name := req.Data.Attributes.Name
	desc := req.Data.Attributes.Desc
	sort := req.Data.Attributes.Sort
	city := req.Data.Attributes.City
	country := req.Data.Attributes.Country
	owner := req.Data.Attributes.Owner

	userId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		server.Logger.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	ownerEmp := models.Employee{
		UserID:      userId,
		FirstName:   owner.FirstName,
		SecondName:  owner.SecondName,
		ThirdName:   owner.ThirdName,
		DisplayName: owner.DisplayName,
		Position:    owner.Position,
		Desc:        owner.Desc,
		Verified:    false,
		Role:        roles.RoleOrgOwner,
	}

	sortOrg, err := models.StringToSortOfOrg(sort)
	if err != nil {
		log.WithError(err).Error("Failed to convert sort of organization")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Organization := models.Organization{
		Name:      name,
		Verified:  false,
		Desc:      desc,
		Country:   country,
		City:      city,
		Sort:      sortOrg,
		Employees: []models.Employee{ownerEmp},
	}

	res, err := server.MongoDB.Organization.Insert(r.Context(), Organization)
	if err != nil {
		if strings.HasPrefix(err.Error(), "failed to insert organization") {
			httpkit.RenderErr(w, problems.InternalError("Failed to insert organization"))
			return
		}
		log.Infof("Failed to create organization: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	log.Infof("Organization created: %v by %s", res.ID, userId)
	httpkit.Render(w, responses.Organization(*res))
}
