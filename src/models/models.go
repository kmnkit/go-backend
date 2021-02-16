package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Product 상품 구조체
type Product struct {
	gorm.Model
	Image       string  `json:"img"`
	ImagAlt     string  `json:"imgalt" gorm:"column:imgalt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"` // sql.NullFloat64
	ProductName string  `gorm:"column:productname" json:"productname"`
	Description string
}

// TableName 상품 테이블명 리턴
func (Product) TableName() string {
	return "products"
}

// Customer 고객 구조체
type Customer struct {
	gorm.Model
	Name      string  `json:"name"`
	FirstName string  `gorm:"column:firstname" json:"firstname"`
	LastName  string  `gorm:"column:lastname" json:"lastname"`
	Email     string  `gorm:"column:email" json:"email"`
	Pass      string  `json:"password"`
	LoggedIn  bool    `gorm:"column:loggedin" json:"loggedin"`
	Orders    []Order `json:"orders"`
}

// TableName 고객 테이블 리턴
func (Customer) TableName() string {
	return "customers"
}

// Order 주문 구조체
type Order struct {
	gorm.Model
	Product
	Customer
	CustomerID   int       `json:"customer_id"`
	ProductID    int       `json:"product_id"`
	Price        float64   `gorm:"column:price" json:"sell_price"`
	PurchaseDate time.Time `gorm:"column:purchase_date" json:"purchase_date"`
}

// TableName 주문 테이블명 리턴
func (Order) TableName() string {
	return "orders"
}
