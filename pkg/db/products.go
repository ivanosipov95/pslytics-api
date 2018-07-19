package db

import "time"

type Product struct {
	CreatedAt time.Time   `json:"-"`
	ID        string      `json:"id" gorm:"primary_key"`
	Name      string      `json:"name"`
	Poster    string      `json:"poster" gorm:"-"`
	Price     int64       `json:"price" gorm:"-"`
	Released  time.Time   `json:"released"`
	Rate      *Rate       `json:"rate" gorm:"-"`
	Discounts []*Discount `json:"discounts" gorm:"-"`
}

type ProductMgr interface {
	GetAllProducts() ([]*Product, error)
	CreateProduct(product *Product) error
	EnsureProductExists(product *Product) (*Product, error)
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
