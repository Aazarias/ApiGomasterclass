package db

import "github.com/Aazarias/ApiGomasterclass/model"

type Storage struct {
	Games StorageGames
	User  StorageUser
}

type StorageGames interface {
	GetByID(id string) (*model.Games, error)
	GetAll() ([]model.Games, error)
	DeleteByID(id string) error
	Create(u *model.Games) (*model.Games, error)
	Update(id string, data map[string]interface{}) (*model.Games, error)
}

type StorageUser interface {
	GetByID(id string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetAll() ([]model.User, error)
	DeleteByID(id string) error
	Create(u *model.User) (*model.User, error)
	Update(id string, data map[string]interface{}) (*model.User, error)
}
