package data

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"go-employee-web-server/internal/models"
)

func FetchEmployees() ([]models.Employee, error) {
	resp, err := http.Get("https://dummy.restapiexample.com/api/v1/employees")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var employeesResponse models.EmployeesResponse
	err = json.Unmarshal(body, &employeesResponse)
	if err != nil {
		return nil, err
	}

	return employeesResponse.Data, nil
}

func SaveEmployeesToFile(employees []models.Employee) error {
	file, err := os.Create("web/data/employees.txt")
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, employee := range employees {
		record := []string{
			fmt.Sprint(employee.ID),
			employee.EmployeeName,
			fmt.Sprint(employee.EmployeeSalary),
			fmt.Sprint(employee.EmployeeAge),
			employee.ProfileImage,
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}

func LoadEmployeesFromFile() ([]models.Employee, error) {
	file, err := os.Open("web/data/employees.txt")
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var employees []models.Employee
	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		salary, _ := strconv.Atoi(record[2])
		age, _ := strconv.Atoi(record[3])
		employees = append(employees, models.Employee{
			ID:             id,
			EmployeeName:   record[1],
			EmployeeSalary: salary,
			EmployeeAge:    age,
			ProfileImage:   record[4],
		})
	}
	return employees, nil
}
