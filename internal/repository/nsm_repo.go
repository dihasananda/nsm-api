package repository

import (
	"nsm-api/internal/entities"

	"gorm.io/gorm"
)

type NSMRepository struct {
	DB *gorm.DB
}

func NewNSMRepository(db *gorm.DB) *NSMRepository {
	return &NSMRepository{DB: db}
}

func (r *NSMRepository) GetAll() ([]entities.NSM, error) {
	var nsms []entities.NSM
	result := r.DB.Find(&nsms)
	return nsms, result.Error
}
