package models

import "github.com/gofrs/uuid"

type Response struct {
	Status string      `json:"status"`
	Errors string      `json:"errors"`
	Data   interface{} `json:"data"`
}

type Users struct {
	UserID    uuid.UUID `json:"userID" gorm:"column:user_id;type:uuid;primaryKey;default:(-)" example:"b251379e-01a1-11ec-82d6-a312e0239c7b"`
	FirstName string    `json:"firstName" gorm:"column:first_name" example:"John"`
	LastName  string    `json:"lastName" gorm:"column:last_name" example:"Doe"`
	Email     string    `json:"email" gorm:"column:email" example:"john.doe@email.com"`
}
