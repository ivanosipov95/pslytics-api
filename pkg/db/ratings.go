package db

import "time"

type Rate struct {
	ID        int32 `gorm:"primary_key" sql:"AUTO_INCREMENT"`
	Date      time.Time
	ProductID string
	Product   Product `gorm:"ForeignKey:ProductID"`
	Total     int64
	Value     float64
}

type RateMgr interface {
	CreateRate(rate *Rate) error
	EnsureRateExists(rate *Rate) (*Rate, error)
}

func (mgr *AppDatabaseMgr) CreateRate(rate *Rate) error {
	panic("not implemented")
}

func (mgr *AppDatabaseMgr) EnsureRateExists(rate *Rate) (*Rate, error) {
	panic("not implemented")
}
