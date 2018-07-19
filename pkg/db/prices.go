package db

import "time"

type Price struct {
	ID        int64     `gorm:"primary_key" sql:"AUTO_INCREMENT" json:"-"`
	Date      time.Time `json:"-"`
	ProductID string    `json:"-"`
	Value     int64     `json:"value"`
}

type PriceMgr interface {
	CreatePrice(price *Price) error
	GetPriceForProduct(id string) (*Price, error)
	EnsurePriceExists(price *Price) error
}

func (mgr *AppDatabaseMgr) GetPriceForProduct(id string) (*Price, error) {
	price := Price{}
	if err := mgr.db.Where("product_id = ?", id).Last(&price).Error; err != nil {
		return nil, err
	}
	return &price, nil
}

func (mgr *AppDatabaseMgr) CreatePrice(price *Price) error {
	return mgr.db.Create(price).Error
}

func (mgr *AppDatabaseMgr) EnsurePriceExists(price *Price) error {
	dbPrice, err := mgr.GetPriceForProduct(price.ProductID)
	if err != nil {
		return mgr.CreatePrice(price)
	}

	if price.Value == dbPrice.Value {
		return nil
	}
	return mgr.CreatePrice(price)
}
