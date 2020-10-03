package usecase

import (
	"doc-server/domain/model"
	"doc-server/domain/repository"
)

type (
	bundleUsecase struct {
		repository.BundleRepository
	}
	BundleUseCase interface {
		GetBundles(limit int, offset int) ([]*model.Bundle, error)
	}
)

func NewBundleUseCase(r repository.BundleRepository) BundleUseCase {
	return &bundleUsecase{r}
}

func (u bundleUsecase)GetBundles(limit int, offset int) ([]*model.Bundle, error) {
	options := &repository.BundleOption{Limit: limit, Offset: offset}
	return u.BundleRepository.Fetch(options)
}