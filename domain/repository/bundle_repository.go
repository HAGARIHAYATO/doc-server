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
}

type BundleOption struct {
	UserID int64
	Limit int
	Offset int
}

func NewBundleRepository(Conn *sql.DB) BundleRepository {
	return &bundleRepository{Conn}
}

func (r bundleRepository) Fetch(options *BundleOption) ([]*model.Bundle, error) {
	var bundles []*model.Bundle
	rows, err := r.conn.Query("SELECT * FROM bundles;")
	for rows.Next() {
		bundle := &model.Bundle{}
		//err = rows.Scan(&doc.ID, &doc.Title, &doc.Text)
		if err == nil {
			bundles = append(bundles, bundle)
		}
	}
	return bundles, err
}
