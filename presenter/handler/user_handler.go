package handler

import (
	"doc-server/domain/model"
	"doc-server/usecase"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
)

type(
	userHandler struct {
		usecase.UserUseCase
	}
	UserHandler interface {
		UserCreate(w http.ResponseWriter, r *http.Request)
		UserSession(w http.ResponseWriter, r *http.Request)
		VerifyAccess(w http.ResponseWriter, r *http.Request)
	}
    JWTResponse struct {
    	ID int64
		Token string
	}
	SessionRequest struct {
		Email string
		Password string
	}
)

func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{u}
}

func (h *userHandler) UserCreate(w http.ResponseWriter, r *http.Request) {
	user := &model.User{
		Name: r.FormValue("name"),
		Email: r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	user, err := h.UserUseCase.CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}

	token, err := user.CreateToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	resp := &JWTResponse{
		ID: user.ID,
		Token: token,
	}

	res, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	_ ,err = w.Write(res)
	if err != nil {
		log.Fatal(err)
	}
}


// sessionはフォームでｎはなくjsonで送る
func (h *userHandler) UserSession(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var req SessionRequest
	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal(err)
	}

	user, err := h.UserUseCase.GetUserByEmail(req.Email)

	hashedPassword := user.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := user.CreateToken()
	if err != nil {
		log.Fatal(err)
	}

	resp := &JWTResponse{
		ID: user.ID,
		Token: token,
	}

	res, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	_ ,err = w.Write(res)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *userHandler) VerifyAccess(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	token = strings.Trim(token, "Bearer%20")

	email, err := model.DecodeToken(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	user, err := h.UserUseCase.GetUserByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	_ ,err = w.Write(res)
	if err != nil {
		log.Fatal(err)
	}
}

