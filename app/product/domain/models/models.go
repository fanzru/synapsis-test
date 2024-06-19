package models

import (
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
	Orders    []Order
	Carts     []Cart
}

type Category struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"unique;not null"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}

func (Category) TableName() string {
	return "category"
}

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	Price       float64  `gorm:"not null"`
	Stock       int      `gorm:"not null"`
	CategoryID  uint     `gorm:"not null"`
	Category    Category `gorm:"foreignKey:CategoryID"`
}

func (Product) TableName() string {
	return "product"
}

type Order struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"`
	TotalAmount float64   `gorm:"not null"`
	Status      string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
	User        User      `gorm:"foreignKey:UserID"`
	OrderItems  []OrderItem
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`
	Order     Order   `gorm:"foreignKey:OrderID"`
	Product   Product `gorm:"foreignKey:ProductID"`
}

type Cart struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	User      User      `gorm:"foreignKey:UserID"`
	CartItems []CartItem
}

type CartItem struct {
	ID        uint    `gorm:"primaryKey"`
	CartID    uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Cart      Cart    `gorm:"foreignKey:CartID"`
	Product   Product `gorm:"foreignKey:ProductID"`
}
