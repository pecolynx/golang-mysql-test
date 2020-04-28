package gateway

import (
	"github.com/jinzhu/gorm"

	"github.com/pecolynx/golang-mysql-test/pkg/application"
	"github.com/pecolynx/golang-mysql-test/pkg/domain"
)

type repositoryFactory struct {
}

func NewRepositoryFactory() application.RepositoryFactory {
	return &repositoryFactory{}
}

func (r *repositoryFactory) NewUserRepository(db *gorm.DB) domain.UserRepository {
	return NewUserRepository(db)
}
