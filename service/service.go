package service

import (
	"go-ecom/repository"
	"go-ecom/service/product_service"
)

type Service struct {
	ProductService product_service.ProductService
}

func NewService(repo repository.Repository) Service {

	return Service{
		ProductService: product_service.NerProductServiceImpl(repo.ProductRepo),
	}

}
