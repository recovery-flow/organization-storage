package nosql

import (
	"github.com/recovery-flow/organization-storage/internal/data/nosql/repositories"
)

type Repo struct {
	Organization repositories.Organization
}

func NewRepositoryNoSql(uri, dbName string) (*Repo, error) {
	return &Repo{
		Organization: _
	}, nil
}
