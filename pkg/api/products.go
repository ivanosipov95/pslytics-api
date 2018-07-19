package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/objque/pslytics-api/pkg/db"
)

func getProduct(w http.ResponseWriter, r *http.Request) {
	product, err := db.DbMgr.GetProductByID(chi.URLParam(r, "id"))
	if err != nil {
		RaiseInternalIfError(err)
	}

	body, _ := json.Marshal(&product)
	w.Write(body)
}
