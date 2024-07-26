package handlers

import (
	"fmt"
	"go-employee-web-server/internal/data"
	"go-employee-web-server/internal/models"
	"net/http"
	"strconv"
)

func EditHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(r.URL.Path[len("/edit/"):])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var employee *models.Employee

	for _, e := range employees {
		if e.ID == id {
			employee = &e
			break
		}
	}

	if employee == nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	if r.Method == "POST" {
		employee.EmployeeName = r.FormValue("name")
		employee.EmployeeSalary, _ = strconv.Atoi(r.FormValue("salary"))
		employee.EmployeeAge, _ = strconv.Atoi(r.FormValue("age"))
		employee.ProfileImage = r.FormValue("image")

		_ = data.SaveEmployeesToFile(employees)
		http.Redirect(w, r, fmt.Sprintf("/view/%d", employee.ID), http.StatusFound)
		return
	}
	renderTemplate(w, "edit", employee)
}
