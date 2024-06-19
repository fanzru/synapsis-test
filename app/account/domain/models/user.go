package models

import (
	"synapsis-test/app/product/domain/models"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int64          `json:"id" gorm:"id"`
	Username  string         `json:"username" gorm:"username"`
	Password  string         `json:"-" gorm:"password"`
	CreatedAt time.Time      `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"deleted_at"`
	Orders    []models.Order
	Carts     []models.Cart
}

type UserWithoutPassword struct {
	Id        int64          `json:"id" gorm:"id"`
	Username  string         `json:"username" gorm:"username"`
	Status    string         `json:"status" gorm:"status"`
	CreatedAt time.Time      `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"deleted_at"`
}
