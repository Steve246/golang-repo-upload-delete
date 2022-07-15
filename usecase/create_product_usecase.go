package usecase

import (
	"fmt"
	"golang-upload-download/model"
	"golang-upload-download/repository"
	"mime/multipart"
)

type CreateProductUseCase interface {
	CreateProduct(newProduct *model.Product, file multipart.File, fileExt string) error
}

type createProductUseCase struct {
	repo     repository.ProductRepository
	fileRepo repository.FileRepository
}

func (c *createProductUseCase) CreateProduct(newProduct *model.Product, file multipart.File, fileExt string) error {
	fileName := fmt.Sprintf("img-%s.%s", newProduct.ProductId, fileExt) 
	//-> struktur penamaan utk masukin cth img-c001.jpg | img-coo2.jpeg
	fileLocation, err := c.fileRepo.Save(file, fileName)
	if err != nil {
		return err
	} //gak bisa rollback karena bukand ataase

	newProduct.ImgPath = fileLocation

	newProduct.UrlPath = fmt.Sprintf("/product/image/%s", newProduct.ProductId)
	err = c.repo.Add(newProduct)
	if err != nil {
		return err
	} //bikin file repository untuk delete

	return nil
}

func NewCreateProductUseCase(repo repository.ProductRepository, fileRepo repository.FileRepository) CreateProductUseCase {
	return &createProductUseCase{
		repo:     repo,
		fileRepo: fileRepo,
	}
}
