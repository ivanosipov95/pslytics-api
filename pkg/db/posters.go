package db

import (
	"encoding/json"
	"time"
)

type Poster struct {
	ID        int32     `gorm:"primary_key" sql:"AUTO_INCREMENT" json:"-"`
	CreatedAt time.Time `json:"-"`
	ProductID string    `json:"-"`
	URL       string    `json:"url"`
}

type PosterMgr interface {
	CreatePoster(poster *Poster) error
	GetPosterForProduct(id string) (*Poster, error)
	EnsurePosterExists(poster *Poster) error
}

func (mgr *AppDatabaseMgr) CreatePoster(poster *Poster) error {
	return mgr.db.Create(poster).Error
}

func (mgr *AppDatabaseMgr) GetPosterForProduct(id string) (*Poster, error) {
	poster := Poster{}
	if err := mgr.db.Where("product_id = ?", id).Last(&poster).Error; err != nil {
		return nil, err
	}
	return &poster, nil
}

func (mgr *AppDatabaseMgr) EnsurePosterExists(poster *Poster) error {
	if _, err := mgr.GetPosterForProduct(poster.ProductID); err != nil {
		return mgr.CreatePoster(poster)
	}
	return nil
}

func (p *Poster) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.URL)
}

func (p *Poster) UnmarshalJSON(data []byte) error {
	p.URL = string(data)
	return nil
}
