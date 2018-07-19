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
	GetRateForProduct(id string) (*Rate, error)
	EnsureRateExists(rate *Rate) error
}

func (mgr *AppDatabaseMgr) GetRateForProduct(id string) (*Rate, error) {
	rate := Rate{}
	if err := mgr.db.Where("product_id = ?", id).Last(&rate).Error; err != nil {
		return nil, err
	}
	return &rate, nil
}

func (mgr *AppDatabaseMgr) CreateRate(rate *Rate) error {
	return mgr.db.Create(rate).Error
}

func (mgr *AppDatabaseMgr) EnsureRateExists(rate *Rate) error {
	dbRate, err := mgr.GetRateForProduct(rate.ProductID)
	if err != nil {
		// for newest games
		return mgr.CreateRate(rate)
	}

	// rating was already saved for today
	if dbRate.Date.After(time.Now().Truncate(time.Hour * 24)) {
		mgr.db.Where("id = ", dbRate.ID).Updates(map[string]interface{}{
			"total": rate.Total,
			"value": rate.Value,
		})
		return nil
	}
	return mgr.CreateRate(rate)
}
