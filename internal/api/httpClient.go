package api

import (
	"encoding/json"
	apiModels "go-employee-web-server/internal/api/models"
	"go-employee-web-server/internal/models"
	"io"
	"log"
	"net/http"
)

// HTTPClient is an implementation of the APIClient interface
type HTTPClient struct {
	url string
}

func NewHTTPClient(url string) *HTTPClient {
	return &HTTPClient{url: url}
}

// FetchEmployees fetches employees from the remote API
func (h *HTTPClient) FetchEmployees() ([]models.Employee, error) {
	resp, err := http.Get(h.url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Failed to close response body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var employeesResponse apiModels.EmployeesResponse
	err = json.Unmarshal(body, &employeesResponse)
	if err != nil {
		return nil, err
	}

	var internalEmployees []models.Employee
	for _, e := range employeesResponse.Data {
		internalEmployees = append(
			internalEmployees, models.Employee{
				ID:           e.ID,
				Name:         e.EmployeeName,
				Salary:       e.EmployeeSalary,
				Age:          e.EmployeeAge,
				ProfileImage: e.ProfileImage,
			},
		)
	}

	return internalEmployees, nil
}
