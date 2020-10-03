package interactor

import (
	"database/sql"
	"doc-server/domain/repository"
	"doc-server/presenter/handler"
	"doc-server/usecase"
)

type (
	interactor struct {
		conn *sql.DB
	}
	Interactor interface {
		NewRepository() *Repository
		NewUsecase(r *Repository) *UseCase
		NewHandler(u *UseCase) *Handler
	}
)

func NewInteractor(conn *sql.DB) Interactor {
	return &interactor{conn}
}

type Repository struct {
	repository.DocRepository
	repository.BundleRepository
	repository.UserRepository
}

func (i *interactor) NewRepository() *Repository {
	r := &Repository{}
	r.DocRepository = repository.NewDocRepository(i.conn)
	r.BundleRepository = repository.NewBundleRepository(i.conn)
	r.UserRepository = repository.NewUserRepository(i.conn)
	return r
}

type UseCase struct {
	usecase.DocUseCase
	usecase.UserUseCase
	usecase.BundleUseCase
}

func (i *interactor) NewUsecase(r *Repository) *UseCase {
	u := &UseCase{}
	u.DocUseCase = usecase.NewDocUseCase(r.DocRepository)
	u.UserUseCase = usecase.NewUserUseCase(r.UserRepository)
	u.BundleUseCase = usecase.NewBundleUseCase(r.BundleRepository)
	return u
}

type Handler struct {
	handler.DocHandler
	handler.UserHandler
	handler.BundleHandler
}

func (i *interactor) NewHandler(u *UseCase) *Handler {
	h := &Handler{}
	h.DocHandler = handler.NewDocHandler(u.DocUseCase)
	h.UserHandler = handler.NewUserHandler(u.UserUseCase)
	h.BundleHandler = handler.NewBundleHandler(u.BundleUseCase)
	return h
}