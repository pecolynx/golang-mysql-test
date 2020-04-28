package administrator

import (
	"github.com/jinzhu/gorm"

	"github.com/pecolynx/golang-mysql-test/pkg/application"
	"github.com/pecolynx/golang-mysql-test/pkg/domain"
)

type UserService interface {
	FindUserByID(id uint) (domain.User, error)
}

type userService struct {
	repository        application.Repository
	repositoryFactory application.RepositoryFactory
}

func NewUserService(repository application.Repository, repositoryFactory application.RepositoryFactory) UserService {
	return &userService{
		repository:        repository,
		repositoryFactory: repositoryFactory,
	}
}

func (u *userService) FindUserByID(id uint) (domain.User, error) {
	db := u.repository.NewDB().DB()
	userRepository := u.repositoryFactory.NewUserRepository(db)
	administrator := domain.NewAdministrator(userRepository)
	return administrator.FindUserByID(id)
}

func (u *userService) FindUserByID2(id uint) (domain.User, error) {
	var user domain.User
	var err error
	err = u.repository.NewTX().Transaction(func(db *gorm.DB) error {
		userRepository := u.repositoryFactory.NewUserRepository(db)
		administrator := domain.NewAdministrator(userRepository)
		user, err = administrator.FindUserByID(id)
		return err
	})
	return user, err
}
