package responses

import (
	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"github.com/recovery-flow/organization-storage/resources"
)

func ParticipantCollection(participants []models.Participant) resources.ParticipantCollection {
	var res []resources.Participant
	for _, el := range participants {
		ver := "false"
		if el.Verified {
			ver = "true"
		}
		res = append(res, resources.Participant{
			Data: resources.ParticipantData{
				Id:   el.UserID.String(),
				Type: resources.ParticipantType,
				Attributes: resources.ParticipantDataAttributes{
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

	return resources.ParticipantCollection{
		Data: res,
	}
}
