package responses

import (
	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"github.com/recovery-flow/organization-storage/resources"
)

func Organization(organization models.Organization) resources.Organization {
	var orgStamp []string
	if organization.Status != nil {
		for _, el := range organization.Status.Stamp {
			orgStamp = append(orgStamp, string(el))
		}
	}

	var links resources.OrganizationLinks
	if organization.Links != nil {
		if organization.Links.TikTok != nil {
			links.Tiktok = organization.Links.TikTok
		}
		if organization.Links.Facebook != nil {
			links.Facebook = organization.Links.Facebook
		}
		if organization.Links.Twitter != nil {
			links.Twitter = organization.Links.Twitter
		}
		if organization.Links.Instagram != nil {
			links.Instagram = organization.Links.Instagram
		}
		if organization.Links.Linkedin != nil {
			links.Linkedin = organization.Links.Linkedin
		}
		if organization.Links.Telegram != nil {
			links.Telegram = organization.Links.Telegram
		}
		if organization.Links.Discord != nil {
			links.Discord = organization.Links.Discord
		}
		if organization.Links.Website != nil {
			links.Website = organization.Links.Website
		}
	}

	var complicatedStatus resources.ComplicatedStatus
	if organization.Status != nil {
		complicatedStatus = resources.ComplicatedStatus{
			State: string(organization.Status.State),
			Stamp: orgStamp,
			From:  organization.Status.From,
		}
	}

	verOrg := "false"
	if organization.Verified {
		verOrg = "true"
	}

	return resources.Organization{
		Data: resources.OrganizationData{
			Id:   organization.ID.String(),
			Type: resources.OrganizationType,
			Attributes: resources.OrganizationDataAttributes{
				Name:              organization.Name,
				Logo:              organization.Logo,
				Verified:          verOrg,
				Desc:              organization.Desc,
				Sort:              string(organization.Sort),
				Country:           organization.Country,
				City:              organization.City,
				Links:             &links,
				ComplicatedStatus: &complicatedStatus,
			},
			Relationships: resources.OrganizationDataRelationships{
				Participants: resources.OrganizationDataRelationshipsParticipants{
					Links: resources.OrganizationDataRelationshipsParticipantsLinks{
						Self:    "." + resources.LinkGetOrg + organization.ID.Hex() + "/participant",
						Related: "." + resources.LinkGetOrg + organization.ID.Hex() + "/participant?page[size]=10&page[number]=1",
					},
				},
			},
		},
	}
}
