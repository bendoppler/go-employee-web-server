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
				return
			}
			name := r.FormValue("name")
			salary, _ := strconv.Atoi(r.FormValue("salary"))
			age, _ := strconv.Atoi(r.FormValue("age"))
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
