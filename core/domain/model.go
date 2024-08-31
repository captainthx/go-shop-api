package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// define the order status enum

type OrderStatus string
type Role string

const (
	Pending OrderStatus = "pending"
	Success OrderStatus = "success"
	Cancel  OrderStatus = "cancel"
)

const (
	Admin    Role = "admin"
	Customer Role = "customer"
)

func (r *OrderStatus) Scan(value interface{}) error {
	if value == nil {
		*r = ""
		return nil
	}

	switch v := value.(type) {
	case []byte:
		*r = OrderStatus(string(v))
	case string:
		*r = OrderStatus(v)
	default:
		return fmt.Errorf("unsupported type for orderStatus: %T", value)
	}
	return nil
}

func (o OrderStatus) Value() (interface{}, error) {
	return string(o), nil
}

func (r *Role) Scan(value interface{}) error {
	if value == nil {
		*r = ""
		return nil
	}

	switch v := value.(type) {
	case []byte:
		*r = Role(string(v))
	case string:
		*r = Role(v)
	default:
		return fmt.Errorf("unsupported type for Role: %T", value)
	}
	return nil
}

func (r Role) Value() (interface{}, error) {
	return string(r), nil
}

// all the domain models are defined here
type User struct {
	gorm.Model
	Name      string      `json:"name" gorm:"type:varchar(30);not null"`
	Username  string      `json:"username" gorm:"type:varchar(20);not null;unique"`
	Password  string      `json:"password" gorm:"type:varchar(255);not null"`
	Avatar    string      `json:"avatar"`
	Email     string      `json:"email" gorm:"type:varchar(50);not null;unique"`
	Role      Role        `json:"role" gorm:"type:ENUM('admin', 'customer');default:'customer'"`
	Orders    []Order     `gorm:"foreignKey:UserID"`
	CartItem  []CartItem  `gorm:"foreignKey:UserID"`
	OrderItem []OrderItem `gorm:"foreignKey:UserID"`
}

type Category struct {
	gorm.Model
	Name     string    `json:"name" gorm:"type:varchar(100);not null"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}

type Product struct {
	gorm.Model
	Name         string         `json:"name" gorm:"type:varchar(100);not null"`
	Price        float64        `json:"price" gorm:"type:decimal(7,2); not null"`
	Quantity     int            `json:"quantity" gorm:"not null"`
	CategoryID   uint           `json:"category_id"`
	ProductImage []ProductImage `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;"`
}

type ProductImage struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	URL       string `json:"url" gorm:"type:varchar(255);not null"`
}

type CartItem struct {
	gorm.Model
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type Order struct {
	gorm.Model
	OrderNumber uuid.UUID   `json:"order_number" gorm:"type:char(36);unique"`
	Status      OrderStatus `json:"status" gorm:"type:ENUM('pending', 'success', 'cancel');default:'pending'"`
	TotalPay    float64     `json:"total_pay" gorm:"type:decimal(7,2);"`
	UserID      uint        `json:"user_id"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Transaction Transaction `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderNumber uuid.UUID `json:"order_number"`
	UserID      uint      `json:"user_id"`
	ProductID   uint      `json:"product_id"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
}

type Transaction struct {
	gorm.Model
	OrderID uuid.UUID `json:"order_id" `
	Amount  float64   `json:"amount" gorm:"type:decimal(7,2);"`
	Status  string    `json:"status"`
	PayTime time.Time `json:"pay_time"`
}
