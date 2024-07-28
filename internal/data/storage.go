package data

import "go-employee-web-server/internal/models"

type Storage interface {
	SaveEmployees(employees []models.Employee) error
	LoadEmployees(searchTerm string, page int, pageSize int) ([]models.Employee, error)
}
