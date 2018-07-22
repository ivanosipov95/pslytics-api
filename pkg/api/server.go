package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func StartAPIServer(ip string, port int) error {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/sales", func(r chi.Router) {
		r.Get("/", listSales)
	})

	r.Route("/products", func(r chi.Router) {
		r.Get("/{id}", getProduct)
		r.Post("/search", searchProduct)
	})

	return http.ListenAndServe(fmt.Sprintf("%s:%d", ip, port), r)
}
