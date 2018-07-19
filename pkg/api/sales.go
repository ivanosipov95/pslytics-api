package api

import (
	"encoding/json"
	"net/http"

	"github.com/objque/pslytics-api/pkg/db"
	"github.com/objque/pslytics-api/pkg/log"
)

func listSales(w http.ResponseWriter, r *http.Request) {
	products, err := db.DbMgr.GetAllProductsWithActiveDiscounts()
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(products)
	w.Write(body)
}
