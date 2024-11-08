package models

import "gorm.io/gorm"

type Micropost struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
}

func (m *Micropost) Create(db *gorm.DB) error {
	return db.Create(m).Error
}

func (m *Micropost) FindByID(db *gorm.DB, id string) error {
	return db.First(m, id).Error
}

func GetAllMicroposts(db *gorm.DB) ([]Micropost, error) {
	var microposts []Micropost
	err := db.Find(&microposts).Error
	return microposts, err
}

func (m *Micropost) Update(db *gorm.DB, newMicropost Micropost) error {
	return db.Model(m).Updates(newMicropost).Error
}

func (m *Micropost) Delete(db *gorm.DB) error {
	return db.Delete(m).Error
}
