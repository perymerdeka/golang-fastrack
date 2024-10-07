package main

import (
	"fmt"
	"myapp/products"
)

func main() {
	factory := products.Product{}

	product := factory.New()
	fmt.Printf("Product %v created at: %v", product.CreatedAt.UTC(), product.ProductName)
}