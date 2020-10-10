package handler

import (
	"doc-server/domain/model"
	"doc-server/usecase"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"log"
	"net/http"
)

type(
	bundleHandler struct {
		usecase.BundleUseCase
		usecase.UserUseCase
	}
	BundleHandler interface {
		BundleIndex(w http.ResponseWriter, r *http.Request)
		BundleShow(w http.ResponseWriter, r *http.Request)
	}
)

func NewBundleHandler(uu usecase.BundleUseCase, ub usecase.UserUseCase) BundleHandler {
	return &bundleHandler{uu, ub}
}

func (b *bundleHandler) BundleIndex(w http.ResponseWriter, r *http.Request) {
	type response struct {
		User *model.User
		Bundles []*model.Bundle
	}

	uid := cast.ToInt64(chi.URLParam(r, "user_id"))

	user, err := b.UserUseCase.GetUserByID(uid)
	if err != nil {
		log.Fatal(err)
	}

	bundles, err := b.BundleUseCase.GetBundlesByUserID(uid)
	if err != nil {
		log.Fatal(err)
	}

	resp := &response{
		user,
		bundles,
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

func (b *bundleHandler) BundleShow(w http.ResponseWriter, r *http.Request) {
	type response struct {
		User *model.User
		Bundle *model.Bundle
	}

	id := cast.ToInt64(chi.URLParam(r, "id"))
	uid := cast.ToInt64(chi.URLParam(r, "user_id"))

	user, err := b.UserUseCase.GetUserByID(uid)
	if err != nil {
		log.Fatal(err)
	}

	bundle, err := b.BundleUseCase.GetBundleByID(id)
	if err != nil {
		log.Fatal(err)
	}

	resp := &response{
		user,
		bundle,
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