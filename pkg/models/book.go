package models

type Book struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  float64 `json:"price"`
}
