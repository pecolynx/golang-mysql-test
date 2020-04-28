package application

import (
	"github.com/jinzhu/gorm"

	"github.com/pecolynx/golang-mysql-test/pkg/domain"
)

type RepositoryFactory interface {
	NewUserRepository(db *gorm.DB) domain.UserRepository
}
