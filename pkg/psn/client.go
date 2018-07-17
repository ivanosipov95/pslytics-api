package psn

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/objque/pslytics-api/pkg/config"
	"github.com/objque/pslytics-api/pkg/db"
)

func Resolve(id string) (*db.Product, error) {
	resp, err := http.Get(fmt.Sprintf("%s/resolve/%s", config.Config.ProxyURL, id))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	product := db.Product{}
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return nil, err
	}
	return &product, nil
}
