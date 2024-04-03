package model

type Book struct {
	Id    int    `gorm:"type:int;primary_key"`
	Title string `gorm:"not null" json:"title, omitempty"`
}
