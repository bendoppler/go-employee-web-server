package utils

import (
	"go-employee-web-server/internal/models"
	"strings"
)

func FilterEmployees(employees []models.Employee, searchTerm string) []models.Employee {
	if searchTerm == "" {
		return employees
	}
	var filteredEmployees []models.Employee
	for _, employee := range employees {
		if strings.Contains(strings.ToLower(employee.Name), strings.ToLower(searchTerm)) {
			filteredEmployees = append(filteredEmployees, employee)
		}
	}
	return filteredEmployees
}
