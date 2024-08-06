package services

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
	weberrors "github.com/piheta/sept/errors"

	"github.com/piheta/sept/models"
	"github.com/piheta/sept/repos"
)

type UserService struct {
	userRepo    *repos.UserRepo
	authService *AuthService
}

func NewUserService(userRepo *repos.UserRepo, authService *AuthService) *UserService {
	return &UserService{
		userRepo:    userRepo,
		authService: authService,
	}
}

func (es *UserService) GetUsers(jwt *models.JWT) (*[]models.User, error) {
	return es.userRepo.GetUsers()
}

func (es *UserService) GetUser(jwt *models.JWT, id uuid.UUID) (*models.User, error) {
	return es.userRepo.GetUser(id)
}

func (es *UserService) CreateUser(user *models.User) error {
	// Validate submitted required fields
	if err := es.validateUser(user); err != nil {
		return weberrors.NewError(400, err.Error())
	}

	// Check if the user already exists
	existingUser, err := es.userRepo.GetUserByEmail(user.Email)
	if err != nil && err.Error() != "user not found" {
		return weberrors.NewError(500, "failed to check existing user")
	}

	if existingUser != nil {
		return weberrors.NewError(409, "user with this email already exists")
	}

	// Hash password and create user
	hash, err := es.authService.HashPassword(user.Password)
	if err != nil {
		return weberrors.NewError(500, "failed to hash password")
	}
	user.Password = hash

	if err := es.userRepo.CreateUser(user); err != nil {
		return weberrors.NewError(500, "failed to create user")
	}

	return nil
}

func (es *UserService) UpdateUser(jwt *models.JWT, id uuid.UUID, updatedUser *models.UpdateUser) error {
	existingUser, err := es.userRepo.GetUser(id)
	if err != nil {
		return err
	}

	if updatedUser.Name != nil {
		existingUser.Name = *updatedUser.Name
	}
	if updatedUser.Email != nil {
		existingUser.Email = *updatedUser.Email
	}
	if updatedUser.Password != nil {
		existingUser.Password = *updatedUser.Password
	}

	// Validate submitted required fields
	if err := es.validateUser(existingUser); err != nil {
		return weberrors.NewError(400, err.Error())
	}

	// Hash password if its submitted
	if updatedUser.Password != nil {
		hash, err := es.authService.HashPassword(*updatedUser.Password)
		if err != nil {
			return weberrors.NewError(500, "failed to hash password")
		}
		existingUser.Password = hash
	}

	return es.userRepo.UpdateUser(existingUser)
}

func (es *UserService) DeleteUser(jwt *models.JWT, id uuid.UUID) error {
	return es.userRepo.DeleteUser(id)
}

//
// HELPER METHODS
//

func (es *UserService) validateUser(user *models.User) error {
	if user == nil {
		return fmt.Errorf("user is required")
	}

	if len(user.Name) == 0 {
		return fmt.Errorf("name is required")
	}

	if err := validateEmail(&user.Email); err != nil {
		return err
	}

	if len(user.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	return nil
}

func validateEmail(email *string) error {
	if email == nil || *email == "" {
		return fmt.Errorf("email is required")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(*email) {
		return fmt.Errorf("invalid email format")
	}
	return nil
}
