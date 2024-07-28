package data

import (
	"encoding/csv"
	"go-employee-web-server/internal/models"
	"go-employee-web-server/internal/utils"
	"log"
	"os"
	"strconv"
)

type FileStorage struct {
	filePath string
	cache    []models.Employee
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
			strconv.Itoa(employee.ID),
			employee.Name,
			strconv.Itoa(employee.Salary),
			strconv.Itoa(employee.Age),
			employee.ProfileImage,
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	f.cache = employees
	return nil
}

func (f *FileStorage) LoadEmployees(searchTerm string, page int, pageSize int) ([]models.Employee, error) {
	if f.cache == nil {
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
		f.cache = employees
	}

	filteredEmployees := utils.FilterEmployees(f.cache, searchTerm)
	start := (page - 1) * pageSize
	end := start + pageSize

	if start > len(filteredEmployees) {
		return []models.Employee{}, nil
	}

	if end > len(filteredEmployees) {
		end = len(filteredEmployees)
	}

	return filteredEmployees[start:end], nil
}

func (f *FileStorage) CountEmployees(searchTerm string) (int, error) {
	if f.cache == nil {
		if _, err := f.LoadEmployees("", 1, 1); err != nil {
			return 0, err
		}
	}

	filteredEmployees := utils.FilterEmployees(f.cache, searchTerm)
	return len(filteredEmployees), nil
}
