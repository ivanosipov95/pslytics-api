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
          "name": "Magicka 2: Special Edition ",
          "poster": "https://store.playstation.com/store/api/chihiro/00_09_000/container/RU/ru/19/EP4139-CUSA01400_00-MAMA02GP40000002/1531810662000/image",
          "prices": {
            "plus-user": {
              "discount-percentage": 64,
              "strikethrough-price": {
                "display": "RUB 1.399",
                "value": 139900
              },
              "upsell-price": null,
              "is-plus": false,
              "availability": {
                "end-date": "2018-07-25T22:59:00Z",
                "start-date": "2018-07-11T00:00:00Z"
              },
              "actual-price": {
                "display": "RUB 499",
                "value": 49900
              }
            },
            "non-plus-user": {
              "discount-percentage": 64,
              "strikethrough-price": {
                "display": "RUB 1.399",
                "value": 139900
              },
              "upsell-price": null,
              "is-plus": false,
              "availability": {
                "end-date": "2018-07-25T22:59:00Z",
                "start-date": "2018-07-11T00:00:00Z"
              },
              "actual-price": {
                "display": "RUB 499",
                "value": 49900
              }
            }
          },
          "score": {
            "total": 17,
            "score": 3.18
          },
          "released": "2015-05-26T00:00:00Z",
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
}
