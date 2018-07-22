package db

import "time"

type Discount struct {
	ID         int64      `gorm:"primary_key" sql:"AUTO_INCREMENT" json:"-"`
	CreatedAt  time.Time  `json:"-"`
	ProductID  string     `json:"-"`
	IsPlus     bool       `json:"is_plus"`
	Value      int64      `json:"value"`
	Percentage int64      `json:"percentage"`
	Since      *time.Time `json:"since"`
	Till       *time.Time `json:"till"`
}

type DiscountMgr interface {
	CreateDiscount(discount *Discount) error
	GetDiscountForProduct(id string, plus bool) (*Discount, error)
	EnsureDiscountExists(discount *Discount) error
}

func (mgr *AppDatabaseMgr) GetDiscountForProduct(id string, plus bool) (*Discount, error) {
	discount := Discount{}
	if err := mgr.db.Where("product_id = ? and is_plus = ?", id, plus).Last(&discount).Error; err != nil {
		return nil, err
	}
	return &discount, nil
}

func (mgr *AppDatabaseMgr) CreateDiscount(discount *Discount) error {
	return mgr.db.Create(discount).Error
}

func (mgr *AppDatabaseMgr) EnsureDiscountExists(discount *Discount) error {
	dbDiscount, err := mgr.GetDiscountForProduct(discount.ProductID, discount.IsPlus)
	if err != nil {
		return mgr.CreateDiscount(discount)
	}

	if dbDiscount.Value == discount.Value {
		return nil
	}
	return mgr.CreateDiscount(discount)
}
