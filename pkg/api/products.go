package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/objque/pslytics-api/pkg/db"
)

func listProducts(w http.ResponseWriter, r *http.Request) {
	products := []db.Product{
		{
			ID:          "god-of-war",
			Name:        "God of War",
			ReleaseDate: time.Now().UTC(),
			Rates: []*db.Rate{
				{
					Total: 100,
					Value: 4.5,
				},
			},
			Posters: []*db.Poster{
				{
					URL: "http://cdn.pic/game/url",
				},
			},
			Prices: []*db.Price{
				{
					IsPlus: false,
					Value:  5999,
				},
				{
					IsPlus: true,
					Value:  3599,
				},
			},
		}, {
			ID:          "fifa",
			Name:        "FIFA18",
			ReleaseDate: time.Now().UTC(),
			Rates: []*db.Rate{
				{
					Total: 146,
					Value: 3.9,
				},
			},
			Posters: []*db.Poster{
				{
					URL: "http://cdn.pic/another-game/url",
				},
			},
			Prices: []*db.Price{
				{
					IsPlus: false,
					Value:  6999,
				},
				{
					IsPlus: true,
					Value:  2500,
				},
			},
		},
	}

	body, _ := json.Marshal(products)
	w.Write(body)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	body, _ := json.Marshal(db.Product{
		ID:          chi.URLParam(r, "id"),
		Name:        "God of War",
		ReleaseDate: time.Now().UTC(),
	})

	w.Write(body)
}
