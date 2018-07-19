package psn

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/objque/pslytics-api/pkg/config"
	"github.com/stretchr/testify/assert"
)

var (
	server *httptest.Server
	mux    *http.ServeMux
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	config.Config = &config.AppConfig{
		ProxyURL: server.URL,
	}
}

func teardown() {
	server.Close()
}

func TestResolve(t *testing.T) {
	// arrange
	setup()
	defer teardown()
	mux.HandleFunc("/resolve/EP4139-CUSA01400_00-MAMA02GP40000002", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
          "released": "2015-05-26T00:00:00Z",
          "rate": {
            "total": 17,
            "value": 3.18
          },
          "name": "Magicka 2: Special Edition ",
          "discounts": [
            {
              "is_plus": false,
              "till": "2018-07-25T22:59:00Z",
              "percentage": 64,
              "since": "2018-07-11T00:00:00Z",
              "value": 499
            },
            {
              "is_plus": false,
              "till": "2018-07-25T22:59:00Z",
              "percentage": 64,
              "since": "2018-07-11T00:00:00Z",
              "value": 499
            }
          ],
          "poster": "https://store.playstation.com/store/api/chihiro/00_09_000/container/RU/ru/19/EP4139-CUSA01400_00-MAMA02GP40000002/1531810662000/image",
          "price": 1399,
          "type": "game",
          "id": "EP4139-CUSA01400_00-MAMA02GP40000002"
        }`))
	})

	// action
	product, err := Resolve("EP4139-CUSA01400_00-MAMA02GP40000002")

	// assert
	assert.NoError(t, err)
	assert.Equal(t, "EP4139-CUSA01400_00-MAMA02GP40000002", product.ID)
	assert.Equal(t, "Magicka 2: Special Edition ", product.Name)
	assert.Equal(t, int64(1399), product.PriceValue)
	assert.Len(t, product.Discounts, 2)
	assert.Equal(t, product.Rate.Value, 3.18)
	assert.Equal(t, product.Rate.Total, int64(17))
	assert.Contains(t, product.PosterURL, product.ID)
}
