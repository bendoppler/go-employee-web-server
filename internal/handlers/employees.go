package handlers

import (
	"go-employee-web-server/internal/api"
	"go-employee-web-server/internal/data"
	"go-employee-web-server/internal/models"
	"go-employee-web-server/internal/utils"
	"log"
	"net/http"
	"os"
	"sync"
)

var (
	once      sync.Once
	mu        sync.RWMutex
	employees []models.Employee
)

func initializeEmployees(storage data.Storage, apiClient api.APIClient) {
	var err error

	if _, err = os.Stat("web/data/employees.txt"); os.IsNotExist(err) {
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
	} else {
		employees, err = storage.LoadEmployees()
		if err != nil {
			log.Fatalf("Error loading employees from file: %v", err)
			return
		}
	}
}

func EmployeesHandler(storage data.Storage, apiClient api.APIClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		once.Do(
			func() {
				mu.Lock()
				defer mu.Unlock()
				initializeEmployees(storage, apiClient)
			},
		)

		mu.RLock()
		defer mu.RUnlock()
		searchTerm := r.URL.Query().Get("search")
		filteredEmployees := utils.FilterEmployees(employees, searchTerm)

		employeeData := struct {
			Employees  []models.Employee
			SearchTerm string
		}{
			Employees:  filteredEmployees,
			SearchTerm: searchTerm,
		}
		renderTemplate(w, "employees", employeeData)
	}
}
