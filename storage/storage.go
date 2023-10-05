package storage

import (
	"context"
	"main/models"
)

type StoregeI interface {
	User() UserI
}

type UserI interface {
	Create(context.Context, models.CreateUser) (string, error)
	Update(context.Context, string, models.UpdateUser) (string, error)
	Get(context.Context, models.IdRequest) (models.User, error)
	GetAll(context.Context, models.GetAllUserRequest) (models.GetAllUser, error)
	Delete(context.Context, models.IdRequest) (string, error)
	GetByLoging(context.Context, models.GetByLoginReq) (models.GetByLoginRes, error)
}
