package handlers

import (
	"go-employee-web-server/internal/data"
	"go-employee-web-server/internal/models"
	"net/http"
	"strconv"
)

func AddHandler(storage data.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusBadRequest)
				return
			}

			name := r.FormValue("name")
			salary, err := strconv.Atoi(r.FormValue("salary"))
			if err != nil {
				http.Error(w, "Invalid salary", http.StatusBadRequest)
				return
			}

			age, err := strconv.Atoi(r.FormValue("age"))
			if err != nil {
				http.Error(w, "Invalid age", http.StatusBadRequest)
				return
			}

			profileImage := r.FormValue("profileImage")

			newEmployee := models.Employee{
				ID:           getNextID(),
				Name:         name,
				Salary:       salary,
				Age:          age,
				ProfileImage: profileImage,
			}

			mu.Lock()
			employees = append(employees, newEmployee)
			err = storage.SaveEmployees(employees)
			if err != nil {
				mu.Unlock()
				http.Error(w, "Error saving employee", http.StatusInternalServerError)
				return
			}
			mu.Unlock()

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		renderTemplate(w, "add", nil)
	}
}

func getNextID() int {
	mu.Lock()
	defer mu.Unlock()

	maxID := 0
	for _, employee := range employees {
		if employee.ID > maxID {
			maxID = employee.ID
		}
	}
	return maxID + 1
}
