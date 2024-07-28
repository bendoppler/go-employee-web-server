package utils

import (
	"strings"

	"go-employee-web-server/internal/models"
)

func FilterEmployees(employees []models.Employee, searchTerm string) []models.Employee {
	if searchTerm == "" {
		return employees
	}
	var result []models.Employee

	for _, e := range employees {
		if ContainsIgnoreCase(e.Name, searchTerm) {
			result = append(result, e)
		}
	}
	return result
}

func ContainsIgnoreCase(str, word string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(word))
}
