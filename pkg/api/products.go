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
			ID:          chi.URLParam(r, "id"),
			Name:        "God of War",
			ReleaseDate: time.Now().UTC(),
		}, {
			ID:          chi.URLParam(r, "id"),
			Name:        "God of War",
			ReleaseDate: time.Now().UTC(),
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
