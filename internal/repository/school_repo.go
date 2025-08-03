package repository

import (
	"nsm-api/internal/entities"

	"gorm.io/gorm"
)

type SchoolRepository struct {
	DB *gorm.DB
}

func NewSchoolRepository(db *gorm.DB) *SchoolRepository {
	return &SchoolRepository{DB: db}
}

func (r *SchoolRepository) GetAll() ([]entities.School, error) {
	var school []entities.School
	result := r.DB.Find(&school)
	return school, result.Error
}
