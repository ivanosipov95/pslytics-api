package db

import "time"

type Poster struct {
	ID        int32 `gorm:"primary_key" sql:"AUTO_INCREMENT"`
	CreatedAt time.Time
	Product   Product `gorm:"ForeignKey:ProductID"`
	ProductID string
	URL       string
}

type PosterMgr interface {
	CreatePoster(poster *Poster) error
	EnsurePosterExists(poster *Poster) (*Poster, error)
}

func (mgr *AppDatabaseMgr) CreatePoster(poster *Poster) error {
	panic("not implemented")
}
func (mgr *AppDatabaseMgr) EnsurePosterExists(poster *Poster) (*Poster, error) {
	panic("not implemented")
}
