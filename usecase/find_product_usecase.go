package usecase

import (
	"golang-upload-download/model"
	"golang-upload-download/repository"
)

type FindProductUseCase interface {
	ById(id string) (model.Product, error)
}

type findProductUseCase struct {
	repo repository.ProductRepository
}

func (l *findProductUseCase) ById(id string) (model.Product, error) {
	return l.repo.FindById(id)
}

func NewFindProductUseCase(repo repository.ProductRepository) FindProductUseCase {
	return &findProductUseCase{repo: repo}
}
