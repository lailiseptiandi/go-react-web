package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/lailiseptiandi/go-web-app/app/dtos"
	"github.com/lailiseptiandi/go-web-app/app/models"
	"github.com/lailiseptiandi/go-web-app/app/repository"
	"github.com/lailiseptiandi/go-web-app/app/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(string) (models.User, error)
	CreateUser(dtos.UserCreate) (primitive.ObjectID, error)
}

type userService struct {
	userRepo repository.UserRepository
	validate *validator.Validate
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo,
		validate: validator.New(),
	}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) GetUserByID(id string) (models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) CreateUser(user dtos.UserCreate) (primitive.ObjectID, error) {
	if err := s.validate.Struct(user); err != nil {
		return primitive.NilObjectID, err
	}

	pass, _ := utils.HashPassword(user.Password)
	createUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: pass,
	}
	return s.userRepo.Create(createUser)
}
