package data

import "go-employee-web-server/internal/models"

type Storage interface {
	SaveEmployees(employees []models.Employee) error
	LoadEmployees() ([]models.Employee, error)
}
