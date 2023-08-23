package library

import (
	"context"
	"sort"
	"strings"
	"time"

	"github.com/marioheryanto/erajaya/go-app/helper"
	"github.com/marioheryanto/erajaya/go-app/model"
	"github.com/marioheryanto/erajaya/go-app/repository"
)

type ProductLibrary struct {
	repo      repository.ProductRepositoryInterface
	validator *helper.Validator
}

type ProductLibraryInterface interface {
	AddProduct(ctx context.Context, product model.Product) error
	GetProduct(ctx context.Context, sort []string) ([]model.Product, error)
}

func NewProductLibrary(repo repository.ProductRepositoryInterface, validator *helper.Validator) ProductLibraryInterface {
	return &ProductLibrary{
		repo:      repo,
		validator: validator,
	}
}

func (l *ProductLibrary) AddProduct(ctx context.Context, Product model.Product) error {
	// validation
	err := l.validator.ValidateStruct(Product)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	_, err = l.repo.CreateProduct(ctx, Product)

	return err
}

func (l *ProductLibrary) GetProduct(ctx context.Context, sortBy []string) ([]model.Product, error) {
	products, err := l.repo.GetProducts(ctx)
	if err != nil {
		return products, err
	}

	if len(sortBy) > 0 {
		sort.Slice(products, func(i, j int) bool {
			switch strings.ToLower(strings.ReplaceAll(sortBy[0], " ", "")) {
			case "termurah":
				return products[i].Price < products[j].Price
			case "termahal":
				return products[i].Price > products[j].Price
			case "name(a-z)":
				return strings.ToLower(products[i].Name) < strings.ToLower(products[j].Name)
			case "name(z-a)":
				return strings.ToLower(products[i].Name) > strings.ToLower(products[j].Name)
			}

			return products[i].PublishAt.After(products[j].PublishAt)
		})
	}

	return products, err
}
