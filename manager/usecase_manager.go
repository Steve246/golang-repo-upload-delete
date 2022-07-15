package manager

import "golang-upload-download/usecase"

type UseCaseManager interface {
	CreateProductUseCase() usecase.CreateProductUseCase
	ListProductUseCase() usecase.ListProductUseCase
	FindProductUseCase() usecase.FindProductUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) FindProductUseCase() usecase.FindProductUseCase {
	return usecase.NewFindProductUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) CreateProductUseCase() usecase.CreateProductUseCase {
	return usecase.NewCreateProductUseCase(u.repoManager.ProductRepo(), u.repoManager.FileRepo())
}

func (u *useCaseManager) ListProductUseCase() usecase.ListProductUseCase {
	return usecase.NewListProductUseCase(u.repoManager.ProductRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
