package repository

import (
	"golang-upload-download/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Add(newpProduct *model.Product) error
	Retrieve() ([]model.Product, error)
	FindById(id string) (model.Product, error) //ambil gambar berdasarkan id
}

type productRepository struct {
	db *gorm.DB
}

func (p *productRepository) Retrieve() ([]model.Product, error) {
	//nambain select
	var products []model.Product
	err := p.db.Select("product_id", "product_name", "is_status", "url_path").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productRepository) FindById(id string) (model.Product, error) {
	var product model.Product
	err := p.db.First(&product, "product_id=?", id).Error
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}
func (p *productRepository) Add(newpProduct *model.Product) error {
	err := p.db.Create(&newpProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}
