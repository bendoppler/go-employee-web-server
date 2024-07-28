package data

import (
	"database/sql"
	"fmt"
	"go-employee-web-server/internal/models"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(db *sql.DB) *MySQLStorage {
	return &MySQLStorage{db: db}
}

func (ms *MySQLStorage) SaveEmployees(employees []models.Employee) error {
	// Clear the table before saving
	_, err := ms.db.Exec("DELETE FROM Employee")
	if err != nil {
		return err
	}

	// Insert new records
	for _, employee := range employees {
		_, err := ms.db.Exec(
			"INSERT INTO Employee (id, employee_name, employee_salary, employee_age, profile_image) VALUES (?, ?, ?, ?, ?)",
			employee.ID, employee.Name, employee.Salary, employee.Age, employee.ProfileImage,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ms *MySQLStorage) LoadEmployees(searchTerm string, page int, pageSize int) ([]models.Employee, error) {
	offset := (page - 1) * pageSize
	query := "SELECT id, employee_name, employee_salary, employee_age, profile_image FROM Employee WHERE employee_name LIKE ? LIMIT ? OFFSET ?"
	rows, err := ms.db.Query(query, "%"+searchTerm+"%", pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			fmt.Printf("Error closing rows: %v\n", err)
		}
	}()

	var employees []models.Employee
	for rows.Next() {
		var employee models.Employee
		err := rows.Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age, &employee.ProfileImage)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}
