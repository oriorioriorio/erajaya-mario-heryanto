package model

import "time"

type Product struct {
	ID          int64     `json:"-"`
	Name        string    `json:"name" validate:"required"`
	Price       float64   `json:"price" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Quantity    int64     `json:"quantity" validate:"required"`
	PublishAt   time.Time `json:"publish_at"`
}

type GetProductParams struct {
	Sort  []string
	Limit uint64
}
