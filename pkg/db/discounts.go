package db

import "time"

type Discount struct {
	ID         int64     `gorm:"primary_key" sql:"AUTO_INCREMENT" json:"-"`
	CreatedAt  time.Time `json:"-"`
	ProductID  string    `json:"-"`
	IsPlus     bool      `json:"is_plus"`
	Value      int64     `json:"value"`
	Percentage int64     `json:"percentage"`
	Since      time.Time `json:"since"`
	Till       time.Time `json:"till"`
}

type DiscountMgr interface {
	CreateDiscount(discount *Discount) error
	EnsureDiscountExists(discount *Discount) (*Discount, error)
}

func (mgr *AppDatabaseMgr) CreateDiscount(discount *Discount) error {
	panic("not implemented")
}

func (mgr *AppDatabaseMgr) EnsureDiscountExists(discount *Discount) (*Discount, error) {
	panic("not implemented")
}
