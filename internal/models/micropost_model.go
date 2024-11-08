package models

type Micropost struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
}
