package main

import (
	"go-ecom/controller"
	"go-ecom/db"
	"go-ecom/repository"
	"go-ecom/router"
	"go-ecom/service"
	"log"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {

		log.Fatal("error opening database: ", err)
		return
	}

	defer db.Close()

	log.Println("connectd to db: " + db.GetDb().DriverName())

	repo := repository.NewRepository(db.GetDb())

	srv := service.NewService(repo)

	controller := controller.NewController(srv)

	router.ProductRouter(controller.ProductController)

	router.Start(":8080")

}
