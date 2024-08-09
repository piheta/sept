package repos

import (
	"github.com/google/uuid"
	weberrors "github.com/piheta/sept/login-server/errors"
	weberrormapper "github.com/piheta/sept/login-server/errors/mappers"

	"github.com/piheta/sept/login-server/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (er *UserRepo) GetUsers() (*[]models.User, error) {
	var users []models.User
	result := er.db.Find(&users)

	if err := weberrormapper.MapGormError("users", result.Error); err != nil {
		return nil, err
	}

	return &users, nil
}

func (er *UserRepo) GetUser(id uuid.UUID) (*models.User, error) {
	var user models.User
	result := er.db.Where("id = ?", id).First(&user)

	if err := weberrormapper.MapGormError("user", result.Error); err != nil {
		return nil, err
	}

	return &user, nil
}

func (er *UserRepo) CreateUser(user *models.User) error {
	user.ID = uuid.New()
	result := er.db.Create(user)

	if err := weberrormapper.MapGormError("user", result.Error); err != nil {
		return err
	}

	return nil
}

func (er *UserRepo) UpdateUser(user *models.User) error {
	result := er.db.Save(user)

	if err := weberrormapper.MapGormError("user", result.Error); err != nil {
		return err
	}

	return nil
}

func (er *UserRepo) DeleteUser(id uuid.UUID) error {
	result := er.db.Where("id = ?", id).Delete(&models.User{})

	if err := weberrormapper.MapGormError("user", result.Error); err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return weberrors.NewError(404, "user not found")
	}

	return nil
}

// AUTH
func (er *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := er.db.Where("email = ?", email).First(&user)

	if err := weberrormapper.MapGormError("user", result.Error); err != nil {
		return nil, err
	}

	return &user, nil
}
