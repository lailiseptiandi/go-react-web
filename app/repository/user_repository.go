package repository

import (
	"context"
	"time"

	"github.com/lailiseptiandi/go-web-app/app/config"
	"github.com/lailiseptiandi/go-web-app/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(id string) (models.User, error)
	Create(user models.User) (primitive.ObjectID, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindAll() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []models.User

	cursor, err := config.UserCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) FindByID(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	err = config.UserCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Create(user models.User) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := config.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}
