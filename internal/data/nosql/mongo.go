package nosql

import (
	"github.com/recovery-flow/organization-storage/internal/data/nosql/repositories"
)

type Repo struct {
	Organization repositories.Organization
}

func NewRepositoryNoSql(uri, dbName string) (*Repo, error) {
	orgRepo, err := repositories.NewOrganization(uri, dbName, "organizations")
	if err != nil {
		return nil, err
	}
	return &Repo{
		Organization: orgRepo,
	}, nil
}
