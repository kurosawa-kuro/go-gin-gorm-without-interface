package services

import (
	"go-gin-gorm-without-interface/internal/models"

	"gorm.io/gorm"
)

type MicropostService struct {
	db *gorm.DB
}

func NewMicropostService(db *gorm.DB) *MicropostService {
	return &MicropostService{db: db}
}

func (s *MicropostService) Create(micropost *models.Micropost) error {
	return s.db.Create(micropost).Error
}

func (s *MicropostService) FindByID(id string) (*models.Micropost, error) {
	var micropost models.Micropost
	err := s.db.First(&micropost, id).Error
	if err != nil {
		return nil, err
	}
	return &micropost, nil
}

func (s *MicropostService) GetAll() ([]models.Micropost, error) {
	var microposts []models.Micropost
	err := s.db.Find(&microposts).Error
	return microposts, err
}

func (s *MicropostService) Update(micropost *models.Micropost, newMicropost models.Micropost) error {
	return s.db.Model(micropost).Updates(newMicropost).Error
}

func (s *MicropostService) Delete(micropost *models.Micropost) error {
	return s.db.Delete(micropost).Error
}

// 必要に応じて追加のビジネスロジックメソッドをここに実装
func (s *MicropostService) ExistsWithTitle(title string) (bool, error) {
	var count int64
	err := s.db.Model(&models.Micropost{}).Where("title = ?", title).Count(&count).Error
	return count > 0, err
}
