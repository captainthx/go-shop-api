package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// define the order status enum

type orderStatus string
type Role string

const (
	Admin    Role = "admin"
	Customer Role = "customer"
)

const (
	Pending orderStatus = "pending"
	Success orderStatus = "success"
	Cancel  orderStatus = "cancel"
)

func (o *orderStatus) Sacn(value interface{}) error {
	*o = orderStatus(value.([]byte))
	return nil
}

func (o orderStatus) Value() (interface{}, error) {
	return string(o), nil
}

func (r *Role) Sacn(value interface{}) error {
	*r = Role(value.([]byte))
	return nil
}

func (r Role) Value() (interface{}, error) {
	return string(r), nil
}

// all the domain models are defined here
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(30);not null"`
	Username string `json:"username" gorm:"type:varchar(20);not null;unique"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email" gorm:"type:varchar(50);not null;unique"`
	Role     Role   `json:"role" gorm:"type:ENUM('admin', 'customer');default:'customer'"`
	Orders   []Order
}

type Product struct {
	gorm.Model
	Name     string  `json:"name" gorm:"type:varchar(100);not null"`
	Price    float64 `json:"amount" gorm:"type:decimal(7,2); not null"`
	Quantity int     `json:"quantity" gorm:"not null"`
}

type ProductImage struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	URL       string `json:"url" gorm:"type:varchar(255);not null"`
}

type Order struct {
	ID            uuid.UUID   `json:"id" grom:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Status        orderStatus `json:"status" gorm:"type:ENUM('pending', 'success', 'cancel');default:'pending'"`
	TotalPay      float64     `json:"total_pay" gorm:"type:decimal(7,2);"`
	UserID        uint        `json:"user_id"`
	ProductId     uint        `json:"product_id"`
	TransactionId uint        `json:"transaction_id"`
	CreateAt      time.Time   `json:"create_at"`
	UpdateAt      time.Time   `json:"update_at"`
	OrderItems    []OrderItem
	Transaction   Transaction
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price" gorm:"type:decimal(7,2);"`
}

type Transaction struct {
	gorm.Model
	OrderID uint      `json:"order_id"`
	Amount  float64   `json:"amount" gorm:"type:decimal(7,2);"`
	Status  string    `json:"status"`
	PayTime time.Time `json:"pay_time"`
}
