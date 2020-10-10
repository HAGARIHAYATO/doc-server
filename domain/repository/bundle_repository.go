package repository

import (
	"database/sql"
	"doc-server/domain/model"
)

type bundleRepository struct{
	conn *sql.DB
}

type BundleRepository interface {
	Fetch(options *BundleOption) ([]*model.Bundle, error)
	FetchByUserID(userId int64) ([]*model.Bundle, error)
	FetchBundledDoc(bundleID int64) ([]*model.BundledDoc, error)
	FetchByID(id int64) (*model.Bundle, error)
}

type BundleOption struct {
	UserID int64
	Limit int
	Offset int
}

func NewBundleRepository(Conn *sql.DB) BundleRepository {
	return &bundleRepository{Conn}
}

func (r *bundleRepository) Fetch(options *BundleOption) ([]*model.Bundle, error) {
	var bundles []*model.Bundle
	rows, err := r.conn.Query("SELECT id, title, user_id FROM bundles;")
	if rows == nil { return nil, err }
	for rows.Next() {
		bundle := &model.Bundle{}
		err = rows.Scan(&bundle.ID, &bundle.Title, &bundle.UserID)

		if err != nil {
			return nil, err
		}

		bundledDocs, err := r.FetchBundledDoc(bundle.ID)

		if err == nil {
			bundle.BundledDocs = bundledDocs
			bundles = append(bundles, bundle)
		}
	}
	return bundles, err
}

func (r *bundleRepository) FetchByUserID(userId int64) ([]*model.Bundle, error) {
	var bundles []*model.Bundle
	rows, err := r.conn.Query("SELECT id, title, user_id FROM bundles WHERE user_id=$1;", userId)
	if rows == nil { return nil, err }
	for rows.Next() {
		bundle := &model.Bundle{}
		err = rows.Scan(&bundle.ID, &bundle.Title, &bundle.UserID)

		if err != nil {
			return nil, err
		}

		bundledDocs, err := r.FetchBundledDoc(bundle.ID)

		if err == nil {
			bundle.BundledDocs = bundledDocs
			bundles = append(bundles, bundle)
		}
	}
	return bundles, err
}

func (r *bundleRepository) FetchBundledDoc(bundleID int64) ([]*model.BundledDoc, error) {
	var bundledDocs []*model.BundledDoc
	rows, err := r.conn.Query("SELECT bundleddocs.id, bundleddocs.doc_id, bundleddocs.bundle_id, docs.id, docs.title, docs.text, docs.user_id FROM bundleddocs INNER JOIN docs ON bundleddocs.doc_id=docs.id WHERE bundleddocs.bundle_id=$1;", bundleID)
	if rows == nil { return nil, err }
	for rows.Next() {
		bundledDoc := &model.BundledDoc{}
		err = rows.Scan(&bundledDoc.ID, &bundledDoc.DocID, &bundledDoc.BundleID, &bundledDoc.Doc.ID, &bundledDoc.Doc.Title, &bundledDoc.Doc.Text, &bundledDoc.Doc.UserID)
		if err == nil {
			bundledDocs = append(bundledDocs, bundledDoc)
		}
	}
	return bundledDocs, nil
}

func (r *bundleRepository) FetchByID(id int64) (*model.Bundle, error) {
	bundle := &model.Bundle{}
	err := r.conn.QueryRow("SELECT id, title, user_id FROM bundles WHERE id=$1;", id).Scan(&bundle.ID, &bundle.Title, &bundle.UserID)
	if err != nil { return nil, err }

	bundledDocs, err := r.FetchBundledDoc(bundle.ID)

	if err == nil {
		bundle.BundledDocs = bundledDocs
	}

	return bundle, err
}