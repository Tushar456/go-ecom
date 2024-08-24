package router

import (
	"go-ecom/controller/product_controller"
	"net/http"

	"github.com/go-chi/chi"
)

var r *chi.Mux

func ProductRouter(pc product_controller.ProductController) *chi.Mux {

	r = chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		r.Post("/", pc.CreateProduct)
		r.Get("/", pc.ListProducts)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", pc.GetProduct)
			r.Delete("/", pc.DeleteProduct)
		})
	})

	return r

}

func Start(addr string) error {

	return http.ListenAndServe(addr, r)

}
