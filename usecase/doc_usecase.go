package usecase

import (
	"doc-server/domain/model"
	"doc-server/domain/repository"
)

type (
	docUsecase struct {
		repository.DocRepository
	}
	DocUseCase interface {
		GetDocs(limit int, offset int) ([]*model.Doc, error)
		GetDoc(id int64) (*model.Doc, error)
		CreateDoc(doc *model.Doc) (*model.Doc, error)
	}
)

func NewDocUseCase(r repository.DocRepository) DocUseCase {
	return &docUsecase{r}
}

func (u docUsecase) GetDocs(limit int, offset int) ([]*model.Doc, error) {
	options := &repository.DocOption{Limit: limit, Offset: offset}
	return u.DocRepository.Fetch(options)
}

func (u docUsecase) GetDoc(id int64) (*model.Doc, error) {
	return u.DocRepository.FetchByID(id)
}

func (u docUsecase) CreateDoc(doc *model.Doc) (*model.Doc, error) {
	return u.DocRepository.Create(doc)
}


