package usecase

import (
	"doc-server/domain/model"
	"doc-server/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type (
	userUsecase struct {
		repository.UserRepository
		repository.DocRepository
	}
	UserUseCase interface {
		GetUsers(limit int, offset int) ([]*model.User, error)
		CreateUser(user *model.User) (*model.User, error)
		GetUserByEmail(email string) (*model.User, error)
		GetUserByID(id int64) (*model.User, error)
	}
)

func NewUserUseCase(r repository.UserRepository, d repository.DocRepository) UserUseCase {
	return &userUsecase{r, d}
}

func (u userUsecase) GetUsers(limit int, offset int) ([]*model.User, error) {
	options := &repository.UserOption{Limit: limit, Offset: offset}
	return u.UserRepository.Fetch(options)
}

func (u userUsecase) GetUserByEmail(email string) (*model.User, error) {
	user, err := u.UserRepository.FetchByEmail(email)
	if err != nil {
		return nil, err
	}
	docs, err := u.DocRepository.FetchByUserID(user.ID)
	user.Docs = docs
	return user, err
}

func (u userUsecase) GetUserByID(id int64) (*model.User, error) {
	user, err := u.UserRepository.FetchByID(id)
	if err != nil {
		return nil, err
	}
	docs, err := u.DocRepository.FetchByUserID(user.ID)
	user.Docs = docs
	return user, err
}

func (u userUsecase) CreateUser(user *model.User) (*model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return nil, err
	}

	user.Password = string(hash)
	return u.UserRepository.Create(user)
}