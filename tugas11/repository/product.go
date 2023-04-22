package repository

import (
	"gorm.io/gorm"
	"tesjwt.go/models"
)

type ProductRepository interface {
	FindAll() (product []models.Product, err error)
	FindById(productId int) (product models.Product, err error)
}

type productRepo struct {
	db *gorm.DB
}

func ProductRepo(db *gorm.DB) ProductRepository {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) FindAll() (products []models.Product, err error) {
	err = p.db.Debug().Find(&products).Error
	return
}

func (p *productRepo) FindById(productId int) (product models.Product, err error) {
	err = p.db.Debug().First(&product, productId).Error
	return
}
