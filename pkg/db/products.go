package db

import "time"

type Product struct {
	CreatedAt time.Time   `json:"-"`
	ID        string      `json:"id" gorm:"primary_key"`
	Name      string      `json:"name"`
	PosterURL string      `json:"poster" gorm:"-"`
	Released  time.Time   `json:"released"`
	Rate      *Rate       `json:"rate"`
	Discounts []*Discount `json:"discounts"`
	Price     Price
	Poster    Poster
}

type ProductMgr interface {
	GetProductByID(id string) (*Product, error)
	GetAllProducts() ([]*Product, error)
	CreateProduct(product *Product) error
	EnsureProductExists(product *Product) (*Product, error)
	GetAllProductsWithActiveDiscounts() ([]*Product, error)
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
