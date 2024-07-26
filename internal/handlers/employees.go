package handlers

import (
	"log"
	"net/http"
	"os"
	"sync"

	"go-employee-web-server/internal/data"
	"go-employee-web-server/internal/models"
	"go-employee-web-server/internal/utils"
)

var (
	once      sync.Once
	mu        sync.RWMutex
	employees []models.Employee
)

func initializeEmployees() {
	var err error

	if _, err = os.Stat("web/data/employees.txt"); os.IsNotExist(err) {
		employees, err = data.FetchEmployees()
		if err != nil {
			log.Fatalf("Error fetching employees: %v", err)
			return
		}
		err = data.SaveEmployeesToFile(employees)
		if err != nil {
			log.Fatalf("Error saving employees to file: %v", err)
			return
		}
	} else {
		employees, err = data.LoadEmployeesFromFile()
		if err != nil {
			log.Fatalf("Error loading employees from file: %v", err)
			return
		}
	}
}

func EmployeesHandler(w http.ResponseWriter, r *http.Request) {
	once.Do(
		func() {
			mu.Lock()
			defer mu.Unlock()
			initializeEmployees()
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
