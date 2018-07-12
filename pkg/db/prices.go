package db

import "time"

type Price struct {
	ID        int64     `gorm:"primary_key" sql:"AUTO_INCREMENT" json:"-"`
	Date      time.Time `json:"-"`
	ProductID string    `json:"-"`
	Product   Product   `gorm:"ForeignKey:ProductID" json:"-"`
	IsPlus    bool      `json:"is_plus"`
	Value     int64     `json:"value"`
}

type PriceMgr interface {
	CreatePrice(price *Price) error
	EnsurePriceExists(price *Price) (*Price, error)
}

func (mgr *AppDatabaseMgr) CreatePrice(price *Price) error {
	panic("not implemented")
}

func (mgr *AppDatabaseMgr) EnsurePriceExists(price *Price) (*Price, error) {
	panic("not implemented")
}
