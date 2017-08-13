package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name          string `json:"name"`
	PrimaryCardID string `json:"primary_card_id"`
}
