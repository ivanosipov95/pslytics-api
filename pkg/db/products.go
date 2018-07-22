package db

import (
	"html/template"
	"time"
)

type Product struct {
	CreatedAt time.Time   `json:"-"`
	ID        string      `json:"id" gorm:"primary_key"`
	Name      string      `json:"name"`
	Released  time.Time   `json:"released"`
	Rate      *Rate       `json:"rate"`
	Discounts []*Discount `json:"discounts"`
	Price     Price       `json:"price"`
	Poster    Poster      `json:"poster"`
}

type ProductMgr interface {
	GetProductByID(id string) (*Product, error)
	GetAllProducts() ([]*Product, error)
	CreateProduct(product *Product) error
	EnsureProductExists(product *Product) (*Product, error)
	GetAllProductsWithActiveDiscounts() ([]*Product, error)
	SearchProductsByName(name string) ([]*Product, error)
}

func (mgr *AppDatabaseMgr) GetAllProducts() ([]*Product, error) {
	var products = make([]*Product, 0)
	return products, mgr.db.Find(&products).Error
}

func (mgr *AppDatabaseMgr) CreateProduct(product *Product) error {
	return mgr.db.Create(product).Error
}

func (mgr *AppDatabaseMgr) EnsureProductExists(product *Product) (*Product, error) {
	panic("not implemented")
}

func (mgr *AppDatabaseMgr) GetAllProductsWithActiveDiscounts() ([]*Product, error) {
	products := []*Product{}
	return products, mgr.db.
		Preload("Discounts").
		Preload("Rate").
		Preload("Price").
		Preload("Poster").
		Limit(50).
		Find(&products).Error
}

func (mgr *AppDatabaseMgr) GetProductByID(id string) (*Product, error) {
	product := Product{ID: id}
	if err := mgr.db.First(&product).Error; err != nil {
		return nil, err
	}

	return &product, mgr.db.
		Preload("Discounts").
		Preload("Rate").
		Preload("Price").
		Preload("Poster").
		Find(&product).Error
}

func (mgr *AppDatabaseMgr) SearchProductsByName(name string) ([]*Product, error) {
	name = template.HTMLEscapeString(name)
	name = "%" + name + "%"
	products := []*Product{}
	if err := mgr.db.
		Where("name LIKE ?", name).
		Preload("Poster").
		Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
