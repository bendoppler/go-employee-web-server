package handlers

import (
	"net/http"
	"strconv"

	"go-employee-web-server/internal/data"
	"go-employee-web-server/internal/models"
)

func ViewHandler(storage data.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mu.RLock()
		defer mu.RUnlock()

		id, err := strconv.Atoi(r.URL.Path[len("/view/"):])
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

		renderTemplate(w, "view", employee)
	}
}
