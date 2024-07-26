package main

import (
	"fmt"
	"go-employee-web-server/internal/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Set up signal handling
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("Shutting down server...")
		handlers.Cleanup()
		os.Exit(0)
	}()

	http.HandleFunc("/", handlers.EmployeesHandler)
	http.HandleFunc("/view/", handlers.ViewHandler)
	http.HandleFunc("/edit/", handlers.EditHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
