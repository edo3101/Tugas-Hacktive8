package repository

import (
	"errors"

	"tesjwt.go/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepoMock struct {
	Mock mock.Mock
}

func (p *ProductRepoMock) FindAll() (products []models.Product, err error) {
	args := p.Mock.Called()
	if args.Get(1) != nil {
		return []models.Product{}, errors.New("product not found")
	}

	products = args.Get(0).([]models.Product)
	return products, nil
}

func (p *ProductRepoMock) FindById(productId int) (product models.Product, err error) {
	args := p.Mock.Called(productId)
	if args.Get(1) != nil {
		return product, errors.New("product not found")
	}

	product = args.Get(0).(models.Product)
	return product, nil
}
