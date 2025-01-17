package responses

import (
	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"github.com/recovery-flow/organization-storage/resources"
)

func OrganizationResponse(organization models.Organization) models.Organization{
	return resources.Organization{
		Data: resources.OrganizationData{
			Id: organization.ID.String(),
			Type: resources.OrganizationType,
			Attributes: resources.OrganizationDataAttributes{
				Name: organization.Name,
				Logo: organization.Logo,
				Verified: organization.Logo,
				Desc: organization.Desc,
				Sort: string(organization.Sort),
				Country: organization.Country,
				City: organization.City,
				Links: organization.Links
			}
		}
	}
}
