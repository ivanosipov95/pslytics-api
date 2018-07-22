package fetcher

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/objque/pslytics-api/pkg/config"
	"github.com/objque/pslytics-api/pkg/db"
	"github.com/objque/pslytics-api/pkg/log"
	"github.com/objque/pslytics-api/pkg/psn"
	"github.com/pkg/errors"
)

func fetch() error {
	// load all products from the db
	products, err := db.DbMgr.GetAllProducts()
	if err != nil {
		return errors.Wrap(err, "can't load products from the db")
	}

	// load actual product info from the db
	for _, product := range products {
		actual, err := psn.Resolve(product.ID)
		if err != nil {
			return errors.Wrapf(err, "can't load product '%s' via proxy", product.ID)
		}

		if err := db.DbMgr.EnsurePosterExists(&db.Poster{
			ProductID: product.ID,
			URL:       actual.Poster.URL,
		}); err != nil {
			log.Error("#1", err)
		}

		if err := db.DbMgr.EnsureRateExists(&db.Rate{
			Date:      time.Now().UTC(),
			ProductID: product.ID,
			Total:     actual.Rate.Total,
			Value:     actual.Rate.Value,
		}); err != nil {
			log.Error("#2", err)
		}

		if err := db.DbMgr.EnsurePriceExists(&db.Price{
			Date:      time.Now().UTC(),
			ProductID: product.ID,
			Value:     actual.Price.Value,
		}); err != nil {
			log.Error("#3", err)
		}

		for _, discount := range actual.Discounts {
			discount.ProductID = product.ID
			if err := db.DbMgr.EnsureDiscountExists(discount); err != nil {
				log.Error("#4", err)
			}
		}
	}
	return nil
}

func isMustFetch() bool {
	last, err := db.DbMgr.GetLastFetch()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true
		}

		log.Error(err)
		return false
	}
	return calcDiffHours(last.Date) > config.Config.Fetching.CountOfSkippedHoursToFetch
}

func Run() {
	for {
		if isMustFetch() {
			now := time.Now().UTC()
			log.Infof("Start fetching stage for '%s'...", now.String())
			if err := fetch(); err != nil {
				log.Error(err)
			} else {
				log.Infof("Finish fetching stage '%s'...", time.Now().UTC().String())
				db.DbMgr.SetLastFetch(now)
			}
		}

		time.Sleep(time.Hour)
	}
}
