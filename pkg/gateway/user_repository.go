package gateway

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pecolynx/golang-mysql-test/pkg/domain"
)

type userRepository struct {
	db *gorm.DB
}

type userEntity struct {
	ID        uint
	Name      string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *userEntity) TableName() string {
	return "user"
}

func (u *userEntity) toModel() domain.User {
	return domain.NewUser(u.ID, u.Name, u.Age, u.CreatedAt, u.UpdatedAt)
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) AddUser(user *domain.UserAddParameter) (uint, error) {
	userEntity := userEntity{
		Name: user.Name,
		Age:  user.Age,
	}
	// INSERT INTO `user` (`name`,`age`,`created_at`,`updated_at`) VALUES (?,?,?,?)
	if result := u.db.Create(&userEntity); result.Error != nil {
		return 0, result.Error
	}

	return userEntity.ID, nil
}

func (u *userRepository) UpdateUser(id uint, user *domain.UserUpdateParameter) (bool, error) {
	userEntity, err := u.findUserByID(id)
	if err != nil {
		return false, err
	}
	userEntity.Name = user.Name
	userEntity.Age = user.Age
	// UPDATE `user` SET `name` = ?, `age` = ?, `created_at` = ?, `updated_at` = ?  WHERE `user`.`id` = ?
	result := u.db.Save(&userEntity)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (u *userRepository) RemoveUser(id uint) (bool, error) {
	// DELETE FROM `user`  WHERE (`id` = ?)
	result := u.db.Where("`id` = ?", id).Delete(&userEntity{})
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (u *userRepository) FindUserByID(id uint) (domain.User, error) {
	userEntity, err := u.findUserByID(id)
	if err != nil {
		return nil, err
	}
	return userEntity.toModel(), nil
}

func (u *userRepository) FindUsers(limit, offset int) ([]domain.User, error) {
	return nil, nil
}

func (u *userRepository) FindUnderageUsers(limit, offset int) ([]domain.User, error) {
	var userEntities []userEntity
	// 	SELECT * FROM `user`  WHERE (`age` < ?) LIMIT {limit} OFFSET {offset}
	result := u.db.Where("`age` < ?", 20).Limit(limit).Offset(offset).Find(&userEntities)
	if result.Error != nil {
		return nil, result.Error
	}

	users := make([]domain.User, len(userEntities))
	for i, u := range userEntities {
		users[i] = u.toModel()
	}
	return users, nil
}

func (u *userRepository) findUserByID(id uint) (*userEntity, error) {
	userEntity := userEntity{}
	// SELECT * FROM `user`  WHERE (`user`.`id` = ?) ORDER BY `user`.`id` ASC LIMIT 1
	if result := u.db.First(&userEntity, id); result.Error != nil {
		if result.RecordNotFound() {
			return nil, domain.NewUserNotFoundError(id)
		}
		return nil, result.Error
	}
	return &userEntity, nil
}
