package api

import "go-employee-web-server/internal/models"

type APIClient interface {
	FetchEmployees() ([]models.Employee, error)
}
