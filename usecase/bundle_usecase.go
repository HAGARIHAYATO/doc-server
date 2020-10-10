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
		GetBundlesByUserID(userId int64) ([]*model.Bundle, error)
		GetBundleByID(id int64) (*model.Bundle, error)
	}
)

func NewBundleUseCase(r repository.BundleRepository) BundleUseCase {
	return &bundleUsecase{r}
}

func (u *bundleUsecase)GetBundles(limit int, offset int) ([]*model.Bundle, error) {
	options := &repository.BundleOption{Limit: limit, Offset: offset}
	return u.BundleRepository.Fetch(options)
}

func (u *bundleUsecase)GetBundlesByUserID(userId int64) ([]*model.Bundle, error) {
	return u.BundleRepository.FetchByUserID(userId)
}

func (u *bundleUsecase)GetBundleByID(id int64) (*model.Bundle, error) {
	return u.BundleRepository.FetchByID(id)
}