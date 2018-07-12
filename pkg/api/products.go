package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/objque/pslytics-api/pkg/db"
)

func getProduct(w http.ResponseWriter, r *http.Request) {
	body, _ := json.Marshal(db.Product{
		ID:          chi.URLParam(r, "id"),
		Name:        strings.ToUpper(chi.URLParam(r, "id")),
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
	})

	w.Write(body)
}
