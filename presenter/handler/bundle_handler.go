package handler

import "doc-server/usecase"

type(
	bundleHandler struct {
		usecase.BundleUseCase
	}
	BundleHandler interface {

	}
)

func NewBundleHandler(u usecase.BundleUseCase) BundleHandler {
	return &bundleHandler{u}
}
