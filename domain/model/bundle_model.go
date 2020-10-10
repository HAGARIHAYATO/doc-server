package model

type Bundle struct {
	ID int64
	Title string
	UserID int64
	BundledDocs []*BundledDoc
}

type BundledDoc struct {
	ID int64
	DocID int64
	BundleID int64
	Doc
}