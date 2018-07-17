package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/objque/pslytics-api/pkg/db"
)

func getProduct(w http.ResponseWriter, r *http.Request) {
	body, _ := json.Marshal(db.Product{
		ID:          "EP4139-CUSA01400_00-MAMA02GP40000002",
		Name:        "Magicka 2: Special Edition",
		ReleaseDate: time.Now().UTC(),
		Rate: &db.Rate{
			Total: 17,
			Value: 3.18,
		},
		Poster: "https://store.playstation.com/store/api/chihiro/00_09_000/container/RU/ru/19/EP4139-CUSA01400_00-MAMA02GP40000002/1531810662000/image",
		Price:  5999,
		Discounts: []*db.Discount{
			{
				IsPlus:     true,
				Value:      1399,
				Percentage: 64,
				Since:      time.Now().UTC(),
				Till:       time.Now().UTC().Add(time.Hour),
			},
		},
	})

	w.Write(body)
}
