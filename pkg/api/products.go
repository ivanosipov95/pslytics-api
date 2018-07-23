package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/objque/pslytics-api/pkg/db"
	"github.com/objque/pslytics-api/pkg/log"
	"github.com/pkg/errors"
)

func getProduct(w http.ResponseWriter, r *http.Request) {
	product, err := db.DbMgr.GetProductByID(chi.URLParam(r, "id"))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		RaiseInternalIfError(err)
	}

	body, _ := json.Marshal(&product)
	w.Write(body)
}

func searchProduct(w http.ResponseWriter, r *http.Request) {
	product := db.Product{}
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		log.Error(errors.Wrap(err, "can't decode body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	products, err := db.DbMgr.SearchProductsByName(product.Name)
	if err != nil {
		RaiseInternalIfError(err)
	}

	body, _ := json.Marshal(&products)
	w.Write(body)
}
