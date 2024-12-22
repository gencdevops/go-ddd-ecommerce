package domain

import (
	"ddd-kata-golang/ecommerce/pkg/errors"

	"github.com/google/uuid"
)

type ProductID struct {
	value string
}

func NewProductID() ProductID {
	return ProductID{value: uuid.New().String()}
}

func (p ProductID) String() string {
	return p.value
}

type CategoryID struct {
	Value string
}

func NewCategoryID(value string) CategoryID {
	return CategoryID{Value: value}
}

func (c CategoryID) String() string {
	return c.Value
}

type Product struct {
	id       ProductID
	name     string
	price    float64
	category CategoryID
}

func NewProduct(name string, price float64, categoryID CategoryID) (*Product, error) {
	if name == "" {
		return nil, errors.ErrInvalidProductName
	}
	if price <= 0 {
		return nil, errors.ErrInvalidPrice
	}

	return &Product{
		id:       NewProductID(),
		name:     name,
		price:    price,
		category: categoryID,
	}, nil
}

func (p *Product) ID() ProductID {
	return p.id
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Price() float64 {
	return p.price
}

func (p *Product) Category() CategoryID {
	return p.category
}
