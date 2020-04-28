package domain

type Administrator interface {
	AddUser(user *UserAddParameter) (uint, error)
	FindUserByID(id uint) (User, error)
}

type administrator struct {
	userRepository UserRepository
}

func NewAdministrator(userRepository UserRepository) Administrator {
	return &administrator{
		userRepository: userRepository,
	}
}

func (a *administrator) AddUser(user *UserAddParameter) (uint, error) {
	return 0, nil
}

func (a *administrator) FindUserByID(id uint) (User, error) {
	return a.userRepository.FindUserByID(id)
}
