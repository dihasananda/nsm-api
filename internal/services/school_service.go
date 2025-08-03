package services

import (
	"nsm-api/internal/entities"
	"nsm-api/internal/repository"
)

type SchoolService struct {
	repo *repository.SchoolRepository
}

func NewSchoolService(repo *repository.SchoolRepository) *SchoolService {
	return &SchoolService{repo: repo}
}

func (s *SchoolService) GetAllSchools() ([]entities.School, error) {
	return s.repo.GetAll()
}
