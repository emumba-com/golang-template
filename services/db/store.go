package db

import (
	"github.com/gofrs/uuid"
	"gogintemplate/models/v1"
	"gorm.io/gorm"
)

type Store interface {
	CreateUser(user models.Users) (models.Users, error)
	GetUser(userID uuid.UUID) (models.Users, error)
}

type PGStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) Store {
	return &PGStore{
		db: db,
	}
}

var _ Store = (*PGStore)(nil)
