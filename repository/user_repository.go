package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

/*9.userのrepositoryを作成*/
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

/*11.userのrepositoryを実装*/
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}
/*12.GetUserByEmailの実装*/
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return err
	}
	return nil
}
/*13.CreateUserの実装*/
func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
