package usecase

import (
	"doc-server/domain/model"
	"doc-server/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type (
	userUsecase struct {
		repository.UserRepository
	}
	UserUseCase interface {
		GetUsers(limit int, offset int) ([]*model.User, error)
		CreateUser(user *model.User) (*model.User, error)
		GetUserByEmail(email string) (*model.User, error)
		GetUserByID(id int64) (*model.User, error)
	}
)

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUsecase{r}
}

func (u userUsecase) GetUsers(limit int, offset int) ([]*model.User, error) {
	options := &repository.UserOption{Limit: limit, Offset: offset}
	return u.UserRepository.Fetch(options)
}

func (u userUsecase) GetUserByEmail(email string) (*model.User, error) {
	return u.UserRepository.FetchByEmail(email)
}

func (u userUsecase) GetUserByID(id int64) (*model.User, error) {
	return u.UserRepository.FetchByID(id)
}

func (u userUsecase) CreateUser(user *model.User) (*model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return nil, err
	}

	user.Password = string(hash)
	return u.UserRepository.Create(user)
}