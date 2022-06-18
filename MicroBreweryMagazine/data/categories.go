package data

type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	SubCategory *Category `json:"subCat"`
	CreateOn    string    `json:"-"`
	UpdateOn    string    `json:"-"`
}
