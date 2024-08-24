package controller

import (
	"go-ecom/controller/product_controller"
	"go-ecom/service"
)

type Controller struct {
	ProductController product_controller.ProductController
}

func NewController(svc service.Service) Controller {

	return Controller{
		ProductController: product_controller.NewProductController(svc.ProductService),
	}

}
