package db

import "time"

type Product struct {
	CreatedAt   time.Time
	ID          string `gorm:"primary_key"`
	Name        string
	ReleaseDate time.Time
}

type ProductMgr interface {
	CreateProduct(product *Product) error
	EnsureProductExists(product *Product) (*Product, error)
}

func (mgr *AppDatabaseMgr) CreateProduct(product *Product) error {
	panic("not implemented")
}

func (mgr *AppDatabaseMgr) EnsureProductExists(product *Product) (*Product, error) {
	panic("not implemented")
}
