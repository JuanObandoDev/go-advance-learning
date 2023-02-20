package repository

import (
	"context"

	"github.com/JuanObandoDeveloper/Rest/models"
)

type UserRespository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id int64) (*models.User, error)
	Close() error
}

var implementation UserRespository

func SetRepository(repo UserRespository) {
	implementation = repo
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id int64) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func Close() error {
	return implementation.Close()
}
