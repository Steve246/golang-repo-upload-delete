package manager

import (
	"golang-upload-download/repository"
)

type RepositoryManager interface {
	// ProductRepo Disini kumpulan semua repo dalam 1 project yang dibuat
	ProductRepo() repository.ProductRepository
	FileRepo() repository.FileRepository
}

type repositoryManager struct {
	infra InfraManager
}

func (r *repositoryManager) FileRepo() repository.FileRepository {
	return repository.NewFileRepository(r.infra.FilePath())
}

func (r *repositoryManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra InfraManager) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
