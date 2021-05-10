package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// Seller represents seller data
	Seller struct {
		SellerID 	uint32 `json:"seller_id" bson:"seller_id"`
		Name 		string `json:"name" bson:"name"`
		Address 	string `json:"address" bson:"address"`
		PhoneNumber string `json:"phone-number" bson:"phone-number"`
		OpenHour 	int64  `json:"open_hour" bson:"open_hour"`
		ClosedHour 	int64  `json:"closed_hour" bson:"closed_hour"`
	}

	// SellerBody body for buyer
	SellerBody struct {
		Name        string `json:"name"`
		Address     string `json:"address"`
		PhoneNumber string `json:"phone_number"`
		OpenHour    int64  `json:"open_hour"`
		ClosedHour  int64  `json:"closed_hour"`
	}

	// Customer for customer table
	Customer struct {
		CustomerID uint32 `json:"customer_id"`
		Name       string `json:"customer_name" bson:"customer_name"`
		Address    string `json:"customer_address" bson:"customer_address"`
	}

	// CustomerBody body for Customer
	CustomerBody struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	// Menu represents menu data
	Menu struct {
		MenuID      uint32    `json:"menu_id" bson:"menu_id"`
		SellerID    uint32    `json:"seller_id" bson:"seller_id"`
		Seller      *Seller   `json:"seller" bson:"seller"`
		Name        string    `json:"name" bson:"name"`
		Description string    `bson:"description" json:"description"`
		Price       float32   `json:"price" bson:"price"`
		Calorie     float32   `json:"calorie" bson:"calorie"`
		ImageURL    string    `json:"image_url" bson:"image_url"`
	}

	// MenuBody for receiving body grom json
	MenuBody struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float32 `json:"price"`
		Calorie     float32 `json:"calorie"`
		ImageURL    string  `json:"image_url"`
		SellerID    uint32  `json:"id_seller"`
	}

	// Order represents order
	Order struct {
		ID              uint32        `bson:"order_id" json:"order_id"`
		CustomerID      uint32        `bson:"customer_id" json:"customer_id"`
		SellerID        uint32        `bson:"seller_id" json:"seller_id"`
		SourceAddress   string        `bson:"source_address" json:"source_address"`
		DeliveryAddress string        `bson:"delivery_address" json:"delivery_address"`
		OrderDetails    []OrderDetail `bson:"order_details" json:"order_details"`
		TotalPrice      float32       `bson:"total_price" json:"total_price"`
		Status          string        `bson:"status" json:"status"`
	}

	// OrderDetail will detail the order
	OrderDetail struct {
		MenuID     uint32  `json:"menu_id" bson:"menu_id"`
		Menu       *Menu   `json:"menu" bson:"menu"`
		Quantity   uint32  `json:"quantity" bson:"quantity"`
		TotalPrice float32 `json:"total_price" bson:"total_price"`
	}

	// MenuDetail for body
	MenuDetail struct {
		ProductID uint32 `json:"menu_id"`
		Quantity  uint32 `json:"quantity"`
	}

	// OrderBody body from json
	OrderBody struct {
		CustomerID uint32       `json:"customer_id"`
		SellerID uint32          `json:"seller_id"`
		Menus      []MenuDetail `json:"menus"`
	}

	// Counter represents counter for
	Counter struct {
		CustomerID uint32 `bson:"customer_id" json:"customer_id"`
		MenuID     uint32 `bson:"menu_id" json:"menu_id"`
		SellerID   uint32 `bson:"seller_id" json:"seller_id"`
		OrderID    uint32 `bson:"order_id" json:"order_id"`
	}
)

type (
	// SellerRepository represents repo functions for Seller
	SellerRepository interface {
		Store(seller *Seller) (primitive.ObjectID, error)
		GetAll() ([]Seller, error)
		GetByID(id uint32) (*Seller, error)
		GetByOID(oid primitive.ObjectID) (*Seller, error)
		UpdateArbitrary(id uint32, key string, value interface{}) error
	}

	// CustomerRepository for repo
	CustomerRepository interface {
		Store(customer *Customer) (primitive.ObjectID, error)
		GetAll() ([]Customer, error)
		GetByID(id uint32) (*Customer, error)
		GetByOID(oid primitive.ObjectID) (*Customer, error)
		UpdateArbitrary(id uint32, key string, value interface{}) error
	}

	// MenuRepository represents repo functions for Menu
	MenuRepository interface {
		Store(menu *Menu) (primitive.ObjectID, error)
		GetAll() ([]Menu, error)
		GetBySellerID(sellerID uint32) ([]Menu, error)
		GetByID(id uint32) (*Menu, error)
		GetByOID(oid primitive.ObjectID) (*Menu, error)
		UpdateArbitrary(id uint32, key string, value interface{}) error
	}

	// OrderRepository reprresents repo functions for order
	OrderRepository interface {
		Store(order *Order) (primitive.ObjectID, error)
		GetAll() ([]Order, error)
		GetByID(id uint32) (*Order, error)
		GetByOID(oid primitive.ObjectID) (*Order, error)
		UpdateArbitrary(id uint32, key string, value interface{}) error
		GetBySellerID(sellerID uint32) ([]Order, error)
		GetByBuyerID(buyerID uint32) ([]Order, error)
		GetByBuyerIDAndStatus(buyerID uint32, status string) ([]Order, error)
		GetBySellerIDAndStatus(sellerID uint32, status string) ([]Order, error)
		GetByStatus(status string) ([]Order, error)
	}

	// CounterRepository repo for counter
	CounterRepository interface {
		Get(collectionName string, identifier string) (uint32, error)
	}
)

type (
	// SellerUsecase for Seller usecase
	SellerUsecase interface {
		CreateSeller(seller SellerBody) (uint32, error)
		GetAll() ([]Seller, error)
		GetByID(id uint32) (*Seller, error)
	}

	// CustomerUsecase for usecase
	CustomerUsecase interface {
		CreateCustomer(customer CustomerBody) (uint32, error)
		GetAll() ([]Customer, error)
		GetByID(id uint32) (*Customer, error)
	}

	// MenuUsecase usecase for Menu
	MenuUsecase interface {
		CreateMenu(Menu MenuBody) (uint32, error)
		GetAll() ([]Menu, error)
		GetBySellerID(SellerID uint32) ([]Menu, error)
		GetByID(id uint32) (*Menu, error)
	}

	// OrderUsecase usecase for order
	OrderUsecase interface {
		CreateOrder(order OrderBody) (uint32, error)
		AcceptOrder(id uint32) error
		GetAll() ([]Order, error)
		GetByID(id uint32) (*Order, error)
		GetBySellerID(sellerID uint32) ([]Order, error)
		GetByBuyerID(buyerID uint32) ([]Order, error)
		GetByBuyerIDAndStatus(buyerID uint32, status string) ([]Order, error)
		GetBySellerIDAndStatus(sellerID uint32, status string) ([]Order, error)
		GetByStatus(status string) ([]Order, error)
	}
)