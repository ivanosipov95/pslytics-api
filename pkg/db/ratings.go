package db

import "time"

type Rate struct {
	ID        int32     `gorm:"primary_key" sql:"AUTO_INCREMENT" json:"-"`
	Date      time.Time `json:"-"`
	ProductID string    `json:"-"`
	Total     int64     `json:"total"`
	Value     float64   `json:"value"`
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
