package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

var (
	templates = template.Must(template.ParseFiles("tmpl/employees.html", "tmpl/view.html", "tmpl/edit.html"))
	once      sync.Once
	mu        sync.RWMutex
	employees []Employee
)

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

func fetchEmployees() ([]Employee, error) {
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

	var employeesResponse EmployeesResponse
	err = json.Unmarshal(body, &employeesResponse)
	if err != nil {
		return nil, err
	}

	return employeesResponse.Data, nil
}

func saveEmployeesToFile(employees []Employee) error {
	file, err := os.Create("data/employees.txt")
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

func loadEmployeesFromFile() ([]Employee, error) {
	file, err := os.Open("data/employees.txt")
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

	var employees []Employee
	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		salary, _ := strconv.Atoi(record[2])
		age, _ := strconv.Atoi(record[3])
		employees = append(employees, Employee{
			ID:             id,
			EmployeeName:   record[1],
			EmployeeSalary: salary,
			EmployeeAge:    age,
			ProfileImage:   record[4],
		})
	}
	return employees, nil
}

func initializeEmployees() {
	var err error

	if _, err = os.Stat("data/employees.txt"); os.IsNotExist(err) {
		employees, err = fetchEmployees()
		if err != nil {
			log.Fatalf("Error fetching employees: %v", err)
			return
		}
		err = saveEmployeesToFile(employees)

		if err != nil {
			log.Fatalf("Error saving employees to file: %v", err)
			return
		}
	} else {
		employees, err = loadEmployeesFromFile()
		if err != nil {
			log.Fatalf("Error loading employees from file: %v", err)
			return
		}
	}
}

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	once.Do(func() {
		mu.Lock()
		defer mu.Unlock()
		initializeEmployees()
	})

	mu.RLock()
	defer mu.RUnlock()
	searchTerm := r.URL.Query().Get("search")
	filteredEmployees := filterEmployees(employees, searchTerm)

	data := struct {
		Employees  []Employee
		SearchTerm string
	}{
		Employees:  filteredEmployees,
		SearchTerm: searchTerm,
	}
	renderTemplate(w, "employees", data)
}

func filterEmployees(employees []Employee, searchTerm string) []Employee {
	if searchTerm == "" {
		return employees
	}
	var result []Employee

	for _, e := range employees {
		if containsIgnoreCase(e.EmployeeName, searchTerm) {
			result = append(result, e)
		}
	}
	return result
}

func containsIgnoreCase(str, word string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(word))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	id, err := strconv.Atoi(r.URL.Path[len("/view/"):])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var employee *Employee

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

func editHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(r.URL.Path[len("/edit/"):])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var employee *Employee

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

		_ = saveEmployeesToFile(employees)
		http.Redirect(w, r, fmt.Sprintf("/view/%d", employee.ID), http.StatusFound)
		return
	}
	renderTemplate(w, "edit", employee)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Set up signal handling
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("Shutting down server...")
		cleanup()
		os.Exit(0)
	}()
	http.HandleFunc("/", employeesHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func cleanup() {
	err := os.Remove("data/employees.txt")
	if err != nil {
		log.Printf("Error removing employees.txt: %v", err)
	} else {
		fmt.Println("employees.txt removed successfully")
	}
}
