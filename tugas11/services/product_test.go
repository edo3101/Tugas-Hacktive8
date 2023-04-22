package services

import (
	"testing"

	"tesjwt.go/models"
	"tesjwt.go/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductServiceFindAllSvcNotFound(t *testing.T) {
	productRepo := &repository.ProductRepoMock{Mock: mock.Mock{}}
	productSvcMock := productSvc{productRepo: productRepo}

	productRepo.Mock.On("FindAll").Return([]models.Product{}, "product not found")

	product, err := productSvcMock.FindAllSvc()

	assert.Equal(t, product, []models.Product{}, "product is an empty slice of struct")
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductSeriveFindAllSvcFound(t *testing.T) {
	productRepo := &repository.ProductRepoMock{Mock: mock.Mock{}}
	productSvcMock := productSvc{productRepo: productRepo}

	productsData := []models.Product{
		{
			Title:       "product 1",
			Description: "description of product 1",
			UserID:      1,
		},
		{
			Title:       "product 2",
			Description: "description of product 2",
			UserID:      1,
		},
	}

	productRepo.Mock.On("FindAll").Return(productsData, nil)

	products, err := productSvcMock.FindAllSvc()

	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, productsData[0].Title, products[0].Title, "product found is 'product 1'")
	assert.Equal(t, productsData[1].Title, products[1].Title, "product found is 'product 2'")
}

func TestProductServiceFindByIdSvcNotFound(t *testing.T) {
	productRepo := &repository.ProductRepoMock{Mock: mock.Mock{}}
	productSvcMock := productSvc{productRepo: productRepo}

	productRepo.Mock.On("FindById", 1).Return(models.Product{}, "product not found")

	product, err := productSvcMock.FindByIdSvc(1)

	assert.Equal(t, product, models.Product{}, "product is an empty struct")
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductServiceFindByIdSvcFound(t *testing.T) {
	productRepo := &repository.ProductRepoMock{Mock: mock.Mock{}}
	productSvcMock := productSvc{productRepo: productRepo}

	productOne := models.Product{
		Title:       "product 1",
		Description: "description of product 1",
		UserID:      1,
	}

	productRepo.Mock.On("FindById", 1).Return(productOne, nil)

	product, err := productSvcMock.FindByIdSvc(1)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, productOne.Title, product.Title, "product found is 'product 1'")
	assert.Equal(t, productOne, product, "product found with id '1'")
}
