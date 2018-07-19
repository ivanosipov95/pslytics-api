package fetcher

import (
	"time"

	"github.com/objque/pslytics-api/pkg/db"
	"github.com/objque/pslytics-api/pkg/log"
	"github.com/objque/pslytics-api/pkg/psn"
)

func fetch() {
	// load all products from the db
	products, err := db.DbMgr.GetAllProducts()
	if err != nil {
		log.Error("can't load products from the db", err)
		return
	}

	// load actual product info from the db
	for _, product := range products {
		actual, err := psn.Resolve(product.ID)
		if err != nil {
			log.Error("can't load product via proxy", err)
			continue
		}

		db.DbMgr.EnsurePosterExists(&db.Poster{
			ProductID: product.ID,
			URL:       actual.PosterURL,
		})

		db.DbMgr.EnsureRateExists(&db.Rate{
			Date:      time.Now().UTC(),
			ProductID: product.ID,
			Total:     actual.Rate.Total,
			Value:     actual.Rate.Value,
		})

		db.DbMgr.EnsurePriceExists(&db.Price{
			Date:      time.Now().UTC(),
			ProductID: product.ID,
			Value:     actual.PriceValue,
		})

		for _, discount := range actual.Discounts {
			discount.ProductID = product.ID
			db.DbMgr.EnsureDiscountExists(discount)
		}
	}
}

func Run() {
	// replace with the cron package
	for {
		fetch()
		time.Sleep(time.Hour * 8)
	}
}
