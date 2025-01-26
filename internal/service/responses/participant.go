package responses

import (
	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"github.com/recovery-flow/organization-storage/resources"
)

func Participant(participant models.Participant) resources.Participant {
	ver := "false"
	if participant.Verified {
		ver = "verified"
	}
	return resources.Participant{
		Data: resources.ParticipantData{
			Id:   participant.UserID.String(),
			Type: resources.ParticipantType,
			Attributes: resources.ParticipantDataAttributes{
				FirstName:   participant.FirstName,
				SecondName:  participant.SecondName,
				ThirdName:   participant.ThirdName,
				DisplayName: participant.DisplayName,
				Position:    participant.Position,
				Verified:    ver,
				Desc:        participant.Desc,
				Role:        string(participant.Role),
				CreatedAt:   participant.CreatedAt,
			},
		},
	}
}
