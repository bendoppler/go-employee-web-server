package main

import (
	"go-employee-web-server/internal/api"
	"go-employee-web-server/internal/data"
	"go-employee-web-server/internal/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	baseURL := "https://dummy.restapiexample.com/api/v1/employees"
	storage := data.NewFileStorage("web/data/employees.txt")
	apiClient := api.NewHTTPClient(baseURL)

	handlerFactory := handlers.NewHandlerFactory(storage, apiClient)

	http.HandleFunc("/", handlerFactory.MakeEmployeesHandler())
	http.HandleFunc("/add", handlerFactory.MakeAddHandler())
	http.HandleFunc("/view/", handlerFactory.MakeViewHandler())
	http.HandleFunc("/edit/", handlerFactory.MakeEditHandler())
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		log.Println("Shutting down server...")
		os.Exit(0)
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
