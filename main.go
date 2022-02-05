package main

import (
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/requests"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/dependencyInjector"
)

func main2() {
	// mongoDb.Setup()
	// product := product.Product{
	// 	Name:        "Pasta",
	// 	Description: "Pasta Vermicelli",
	// 	Price:       25.0,
	// 	Quantity:    100,
	// }
	// fmt.Println("Hello")
	// mongoDb.Save(product)

	productController := dependencyInjector.ContainerBuilder{}.GetProductController()
	_, _ = productController.Create(new(requests.CreateProductRequest))

}
