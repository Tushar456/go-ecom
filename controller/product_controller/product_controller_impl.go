package product_controller

import (
	"context"
	"encoding/json"
	"go-ecom/model"
	"go-ecom/service/product_service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type ProductControllerImpl struct {
	ctx context.Context
	ps  product_service.ProductService
}

func NewProductController(ps product_service.ProductService) ProductController {

	return &ProductControllerImpl{
		ctx: context.Background(),
		ps:  ps,
	}

}

func (pc *ProductControllerImpl) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var p ProductReq

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "error while decoding", http.StatusBadRequest)
		return
	}

	product, err := pc.ps.CreateProduct(context.Background(), toStorerProduct(p))
	if err != nil {
		http.Error(w, "error creating product", http.StatusInternalServerError)
		return
	}

	res := toProductRes(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

}

func (pc *ProductControllerImpl) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "error parsing ID", http.StatusBadRequest)
		return
	}

	product, err := pc.ps.GetProduct(context.Background(), i)
	if err != nil {
		http.Error(w, "error getting product", http.StatusInternalServerError)
		return
	}

	res := toProductRes(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

func (pc *ProductControllerImpl) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := pc.ps.ListProducts(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var resp []ProductRes

	for _, pr := range products {
		resp = append(resp, toProductRes(&pr))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func (pc *ProductControllerImpl) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "error parsing ID", http.StatusBadRequest)
		return
	}

	var p ProductReq
	if err = json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "error while decoding", http.StatusBadRequest)
		return
	}

	product, err := pc.ps.GetProduct(pc.ctx, i)
	if err != nil {
		http.Error(w, "error getting product", http.StatusInternalServerError)
		return
	}
	res := toProductRes(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

func (pc *ProductControllerImpl) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "error parsing ID", http.StatusBadRequest)
		return
	}

	err = pc.ps.DeleteProduct(context.Background(), i)
	if err != nil {
		http.Error(w, "error deleting product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func toStorerProduct(p ProductReq) *model.Product {

	return &model.Product{
		Name:         p.Name,
		Image:        p.Image,
		Category:     p.Category,
		Description:  p.Description,
		Rating:       p.Rating,
		NumReviews:   p.NumReviews,
		Price:        p.Price,
		CountInStock: p.CountInStock,
	}

}

func toProductRes(p *model.Product) ProductRes {
	return ProductRes{
		ID:           p.ID,
		Name:         p.Name,
		Image:        p.Image,
		Category:     p.Category,
		Description:  p.Description,
		Rating:       p.Rating,
		NumReviews:   p.NumReviews,
		Price:        p.Price,
		CountInStock: p.CountInStock,
	}

}
