package responses

import (
	"fmt"
	"net/url"

	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"github.com/recovery-flow/organization-storage/resources"
)

func ParticipantCollection(participants []models.Participant, baseURL string, queryParams url.Values, totalItems, pageSize, pageNumber int64) resources.ParticipantCollection {
	var res []resources.ParticipantData
	for _, el := range participants {
		res = append(res, Participant(el).Data)
	}

	links := resources.Links{
		Self:     generatePaginationLink(baseURL, queryParams, pageNumber, pageSize),
		Previous: generatePaginationLink(baseURL, queryParams, pageNumber-1, pageSize),
		Next:     generatePaginationLink(baseURL, queryParams, pageNumber+1, pageSize),
	}

	if pageNumber <= 1 {
		links.Previous = nil
	}

	totalPages := (totalItems + pageSize - 1) / pageSize
	if pageNumber >= totalPages {
		links.Next = nil
	}

	return resources.ParticipantCollection{
		Data:  res,
		Links: links,
	}
}

func generatePaginationLink(baseURL string, queryParams url.Values, pageNumber, pageSize int64) *string {
	if pageNumber < 1 {
		return nil
	}

	queryParams.Set("page[number]", fmt.Sprintf("%d", pageNumber))
	queryParams.Set("page[size]", fmt.Sprintf("%d", pageSize))
	res := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())
	return &res
}
