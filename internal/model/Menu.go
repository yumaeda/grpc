package model

type Menu struct {
	ID          string
	SortOrder   int64
	Category    int64
	SubCategory int64
	Region      int64
	Name        string
	NameJpn     string
	Price       int64
	IsMinPrice  int64
	IsHidden    int64
}
