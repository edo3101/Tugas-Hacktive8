package services

import (
	"log"

	"tesjwt.go/models"
	"tesjwt.go/repository"
)

type ProductSvcInterface interface {
	FindAllSvc() (products []models.Product, err error)
	FindByIdSvc(productId int) (product models.Product, err error)
}

type productSvc struct {
	productRepo repository.ProductRepository
}

func ProductSvc(productRepo repository.ProductRepository) ProductSvcInterface {
	return &productSvc{
		productRepo: productRepo,
	}
}

func (p *productSvc) FindAllSvc() (products []models.Product, err error) {
	if products, err = p.productRepo.FindAll(); err != nil {
		log.Print("error finding product")
	}
	return
}

func (p *productSvc) FindByIdSvc(productId int) (product models.Product, err error) {
	if product, err = p.productRepo.FindById(productId); err != nil {
		log.Print("product not found")
	}
	return
}
