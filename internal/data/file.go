package data

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"go-employee-web-server/internal/models"
)

type FileStorage struct {
	filePath string
}

func NewFileStorage(filePath string) *FileStorage {
	return &FileStorage{filePath: filePath}
}

func (f *FileStorage) SaveEmployees(employees []models.Employee) error {
	file, err := os.Create(f.filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, employee := range employees {
		record := []string{
			fmt.Sprint(employee.ID),
			employee.Name,
			fmt.Sprint(employee.Salary),
			fmt.Sprint(employee.Age),
			employee.ProfileImage,
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}

func (f *FileStorage) LoadEmployees() ([]models.Employee, error) {
	file, err := os.Open(f.filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
		}
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
		employees = append(
			employees, models.Employee{
				ID:           id,
				Name:         record[1],
				Salary:       salary,
				Age:          age,
				ProfileImage: record[4],
			},
		)
	}
	return employees, nil
}
