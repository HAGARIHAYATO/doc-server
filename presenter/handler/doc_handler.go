package handler

import (
	"doc-server/domain/model"
	"doc-server/usecase"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"log"
	"net/http"
	"time"
)

type docHandler struct {
	usecase.DocUseCase
}

type DocHandler interface {
	GetAllDocs(w http.ResponseWriter, r *http.Request)
	GetShowDoc(w http.ResponseWriter, r *http.Request)
	DocCreate(w http.ResponseWriter, r *http.Request)
}

func NewDocHandler(u usecase.DocUseCase) DocHandler {
	return &docHandler{u}
}

func (d *docHandler) GetAllDocs(w http.ResponseWriter, r *http.Request) {
	limit := chi.URLParam(r, "limit")
	offset := chi.URLParam(r, "offset")
	docs, err := d.DocUseCase.GetDocs(cast.ToInt(limit), cast.ToInt(offset))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(docs)
	res, err := json.Marshal(docs)
	if err != nil {
		log.Fatal(err)
	}
	_ ,err = w.Write(res)
	if err != nil {
		log.Fatal(err)
	}
}

func (d *docHandler) GetShowDoc(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	doc, err := d.DocUseCase.GetDoc(cast.ToInt64(id))
	if err != nil {
		log.Fatal(err)
	}
	res, err := json.Marshal(doc)
	if err != nil {
		log.Fatal(err)
	}
	_ ,err = w.Write(res)
	if err != nil {
		log.Fatal(err)
	}
}

func (d *docHandler) DocCreate(w http.ResponseWriter, r *http.Request) {
	doc := &model.Doc{
		Title: r.FormValue("title"),
		Text: r.FormValue("text"),
		UserID: cast.ToInt64(r.FormValue("user_id")),
		CreatedAt: time.Now(),
	}
	doc, err := d.DocUseCase.CreateDoc(doc)
	if err != nil {
		log.Fatal(err)
	}
	res, err := json.Marshal(doc)
	if err != nil {
		log.Fatal(err)
	}
	_ ,err = w.Write(res)
	if err != nil {
		log.Fatal(err)
	}
}
