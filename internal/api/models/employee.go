package models

type Employee struct {
	ID             int    `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary int    `json:"employee_salary"`
	EmployeeAge    int    `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

type EmployeesResponse struct {
	Status string     `json:"status"`
	Data   []Employee `json:"data"`
}
