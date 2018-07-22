package fetcher

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/objque/pslytics-api/pkg/config"
	"github.com/objque/pslytics-api/pkg/db"
	"github.com/stretchr/testify/assert"
)

var (
	server *httptest.Server
	mux    *http.ServeMux
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	db.DbMgr = db.NewFakeDatabaseMgr()
	config.Config = &config.AppConfig{
		ProxyURL: server.URL,
	}
}

func teardown() {
	db.DbMgr.DropAllTables()
	db.DbMgr.Close()
}

func TestFetcher_Fetch(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	product_id := "EP4139-CUSA01400_00-MAMA02GP40000002"
	db.DbMgr.CreateProduct(&db.Product{
		ID:       product_id,
		Name:     "Magicka 2",
		Released: time.Now().UTC(),
	})
	mux.HandleFunc("/resolve/EP4139-CUSA01400_00-MAMA02GP40000002", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
          "released": "2015-05-26T00:00:00Z",
          "rate": {
            "total": 100,
            "value": 4.9
          },
          "discounts": [
            {
              "is_plus": false,
              "till": "2018-07-25T22:59:00Z",
              "percentage": 64,
              "since": "2018-07-11T00:00:00Z",
              "value": 499
            }
          ],
          "name": "Magicka 2: Special Edition ",
          "poster": "https://store.playstation.com/store/api/chihiro/00_09_000/container/RU/ru/19/EP4139-CUSA01400_00-MAMA02GP40000002/1531810662000/image",
          "price": 1399,
          "type": "game",
          "id": "EP4139-CUSA01400_00-MAMA02GP40000002"
        }`))
	})

	// action
	fetch()

	// assert
	poster, err := db.DbMgr.GetPosterForProduct(product_id)
	assert.NoError(t, err)
	assert.Equal(t, product_id, poster.ProductID)
	assert.Contains(t, poster.URL, product_id)

	rate, err := db.DbMgr.GetRateForProduct(product_id)
	assert.NoError(t, err)
	assert.Equal(t, product_id, rate.ProductID)
	assert.Equal(t, int64(100), rate.Total)
	assert.Equal(t, 4.9, rate.Value)

	price, err := db.DbMgr.GetPriceForProduct(product_id)
	assert.NoError(t, err)
	assert.Equal(t, product_id, price.ProductID)
	assert.Equal(t, int64(1399), price.Value)

	discount, err := db.DbMgr.GetDiscountForProduct(product_id, false)
	assert.NoError(t, err)
	assert.Equal(t, product_id, discount.ProductID)
	assert.False(t, discount.IsPlus)
	assert.Equal(t, int64(64), discount.Percentage)
}

func TestFetcher_Internal_IsMustFetch_FirstRun(t *testing.T) {
	// first run means that no records in last_fetches
	setup()
	defer teardown()

	// action
	must := isMustFetch()

	// assert
	assert.True(t, must)
}

func TestFetcher_Internal_IsMustFetch_ReloadApp_AfterFetching(t *testing.T) {
	// fetch was successful and someone restart the app
	setup()
	defer teardown()

	// arrange
	db.DbMgr.SetLastFetch(time.Now().UTC())

	// action
	must := isMustFetch()

	// assert
	assert.False(t, must)
}

func TestFetcher_Internal_IsMustFetch_ReloadApp_AfterOldestFetching(t *testing.T) {
	// fetch was successful some times ago and someone restart the app
	setup()
	defer teardown()

	// arrange
	db.DbMgr.SetLastFetch(time.Now().UTC().Truncate(time.Hour * 48))

	// action
	must := isMustFetch()

	// assert
	assert.True(t, must)
}
