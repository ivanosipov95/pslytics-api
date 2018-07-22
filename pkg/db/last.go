package db

import (
	"time"
)

type LastFetch struct {
	Date time.Time
}

type LastFetchMgr interface {
	GetLastFetch() (*LastFetch, error)
	SetLastFetch(time time.Time) error
}

func (mgr *AppDatabaseMgr) GetLastFetch() (*LastFetch, error) {
	poster := LastFetch{}
	if err := mgr.db.Last(&poster).Error; err != nil {
		return nil, err
	}
	return &poster, nil
}

func (mgr *AppDatabaseMgr) SetLastFetch(time time.Time) error {
	last, err := mgr.GetLastFetch()
	if err != nil {
		return mgr.db.Create(LastFetch{
			Date: time,
		}).Error
	}
	return mgr.db.Model(last).Update("date", time).Error
}
