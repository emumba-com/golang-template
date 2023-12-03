package db

import (
	"github.com/gofrs/uuid"
	"gogintemplate/models/v1"
)

const (
	UsersTable = "users"
)

func (p *PGStore) CreateUser(user models.Users) (models.Users, error) {
	var newUser models.Users

	result := p.db.Table(UsersTable).Create(&user).Scan(&newUser)

	return newUser, result.Error
}

func (p *PGStore) GetUser(userID uuid.UUID) (models.Users, error) {
	var user models.Users

	result := p.db.Where("user_id = ?", userID).Find(&user)

	return user, result.Error
}
