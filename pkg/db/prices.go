package db

import "time"

type Price struct {
	ID        int64 `gorm:"primary_key" sql:"AUTO_INCREMENT"`
	Date      time.Time
	ProductID string
	Product   Product `gorm:"ForeignKey:ProductID"`
	IsPlus    bool
	Value     int64
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
