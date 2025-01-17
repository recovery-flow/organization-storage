package responses

import (
	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"github.com/recovery-flow/organization-storage/resources"
)

func OrganizationResponse(organization models.Organization) resources.Organization {
	var orgStamp []string
	for _, el := range organization.Status.Stamp {
		orgStamp = append(orgStamp, string(el))
	}
	var employees []resources.Employee
	for _, el := range organization.Employees {
		ver := "false"
		if el.Verified {
			ver = "verified"
		}
		employees = append(employees, resources.Employee{
			Data: resources.EmployeeData{
				Id:   el.UserID.String(),
				Type: resources.EmployeeType,
				Attributes: resources.EmployeeDataAttributes{
					FirstName:   el.FirstName,
					SecondName:  el.SecondName,
					ThirdName:   el.ThirdName,
					DisplayName: el.DisplayName,
					Position:    el.Position,
					Verified:    ver,
					Desc:        el.Desc,
					Role:        string(el.Role),
					UpdatedAt:   el.UpdatedAt,
					CreatedAt:   el.CreatedAt,
				},
			},
		})
	}
	return resources.Organization{
		Data: resources.OrganizationData{
			Id:   organization.ID.String(),
			Type: resources.OrganizationType,
			Attributes: resources.OrganizationDataAttributes{
				Name:     organization.Name,
				Logo:     organization.Logo,
				Verified: organization.Logo,
				Desc:     organization.Desc,
				Sort:     string(organization.Sort),
				Country:  organization.Country,
				City:     organization.City,
				Links: &resources.Links{
					Twitter:   organization.Links.Twitter,
					Instagram: organization.Links.Instagram,
					Facebook:  organization.Links.Facebook,
					Tiktok:    organization.Links.TikTok,
					Linkedin:  organization.Links.Linkedin,
					Telegram:  organization.Links.Telegram,
					Discord:   organization.Links.Discord,
				},
				ComplicatedStatus: &resources.ComplicatedStatus{
					State: string(organization.Status.State),
					Stamp: orgStamp,
					From:  organization.Status.From,
				},
			},
		},
		Included: employees,
	}
}
