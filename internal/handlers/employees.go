package handlers

import (
	"go-employee-web-server/internal/api"
	"go-employee-web-server/internal/data"
	"go-employee-web-server/internal/models"
	"go-employee-web-server/internal/utils"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var (
	once      sync.Once
	mu        sync.RWMutex
	employees []models.Employee
	pageSize  = 10
)

func initializeEmployees(storage data.Storage, apiClient api.APIClient) {
	var err error
	employees, err = apiClient.FetchEmployees()
	if err != nil {
		log.Fatalf("Error fetching employees: %v", err)
		return
	}
	err = storage.SaveEmployees(employees)
	if err != nil {
		log.Fatalf("Error saving employees to file: %v", err)
		return
	}
}

func EmployeesHandler(storage data.Storage, apiClient api.APIClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the path is root or not
		if r.URL.Path != "/" {
			http.Error(w, "404 Page Not Found", http.StatusNotFound)
			return
		}
		once.Do(
			func() {
				mu.Lock()
				defer mu.Unlock()
				initializeEmployees(storage, apiClient)
			},
		)

		mu.RLock()
		defer mu.RUnlock()

		// Parse search term
		searchTerm := r.URL.Query().Get("search")
		filteredEmployees := utils.FilterEmployees(employees, searchTerm)

		// Parse page number
		pageStr := r.URL.Query().Get("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}

		// Calculate pagination
		start := (page - 1) * pageSize
		if start >= len(filteredEmployees) {
			start = len(filteredEmployees) - pageSize
		}
		if start < 0 {
			start = 0
		}

		end := start + pageSize
		if end > len(filteredEmployees) {
			end = len(filteredEmployees)
		}

		paginatedEmployees := filteredEmployees[start:end]

		employeeData := struct {
			Employees   []models.Employee
			SearchTerm  string
			CurrentPage int
			TotalPages  int
		}{
			Employees:   paginatedEmployees,
			SearchTerm:  searchTerm,
			CurrentPage: page,
			TotalPages:  (len(filteredEmployees) + pageSize - 1) / pageSize,
		}

		renderTemplate(w, "employees", employeeData)
	}
}
