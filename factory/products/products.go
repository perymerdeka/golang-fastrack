package products

import "time"

type Product struct {
	ProductName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) New () *Product {
	product := Product {
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ProductName: "Product 1",
	}

	return &product
}