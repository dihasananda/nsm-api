package repository

import (
	"nsm-api/internal/entities"

	"gorm.io/gorm"
)

type SchoolRepository struct {
	db *gorm.DB
}

func NewSchoolRepository(db *gorm.DB) *SchoolRepository {
	return &SchoolRepository{db: db}
}

func (r *SchoolRepository) GetAll() ([]entities.School, error) {
	var schools []entities.School
	if err := r.db.Find(&schools).Error; err != nil {
		return nil, err
	}
	return schools, nil
}
