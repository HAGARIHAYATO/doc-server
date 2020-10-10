package repository

import (
	"database/sql"
	"doc-server/domain/model"
)

type docRepository struct{
	conn *sql.DB
}

type DocOption struct {
	IDs []int64
	Limit int
	Offset int
}

type DocRepository interface {
	Fetch(options *DocOption) ([]*model.Doc, error)
	FetchByID(id int64) (*model.Doc, error)
	FetchByUserID(userId int64) ([]*model.Doc, error)
	Create(doc *model.Doc) (*model.Doc, error)
}

func NewDocRepository(Conn *sql.DB) DocRepository {
	return &docRepository{Conn}
}

func (r *docRepository) Fetch(options *DocOption) ([]*model.Doc, error) {
	var docs []*model.Doc
	rows, err := r.conn.Query("SELECT id, title, text, user_id FROM docs;")
	if rows == nil { return nil, err }
	for rows.Next() {
		doc := &model.Doc{}
		err = rows.Scan(&doc.ID, &doc.Title, &doc.Text, &doc.UserID)
		if err == nil {
			docs = append(docs, doc)
		}
	}
	return docs, err
}

func (r *docRepository) FetchByUserID(userId int64) ([]*model.Doc, error) {
	var docs []*model.Doc
	rows, err := r.conn.Query("SELECT id, title, text, user_id FROM docs WHERE user_id=$1;", userId)
	if rows == nil { return nil, err }
	for rows.Next() {
		doc := &model.Doc{}
		err = rows.Scan(&doc.ID, &doc.Title, &doc.Text, &doc.UserID)
		if err == nil {
			docs = append(docs, doc)
		}
	}
	return docs, err
}

func (r *docRepository) FetchByID(id int64) (*model.Doc, error) {
	doc := &model.Doc{}
	rows := r.conn.QueryRow("SELECT id, title, text, user_id FROM docs WHERE id = $1;", id)
	err := rows.Scan(&doc.ID, &doc.Title, &doc.Text, &doc.UserID)
	return doc, err
}

func (r *docRepository) Create(doc *model.Doc) (*model.Doc, error) {
	stmt, err := r.conn.Prepare("INSERT INTO docs(title, text, user_id) VALUES($1, $2, $3) RETURNING id;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var docID int64
	err = stmt.QueryRow(doc.Title, doc.Text).Scan(&docID)
	doc.ID = docID
	return doc, err
}




