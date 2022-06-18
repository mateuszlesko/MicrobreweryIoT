package models

type category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	SubCategory *category `json:"subCat"`
	CreateOn    string    `json:"-"`
	UpdateOn    string    `json:"-"`
}
