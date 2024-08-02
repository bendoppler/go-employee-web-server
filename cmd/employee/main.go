package main

import (
	"database/sql"
	"go-employee-web-server/internal/api"
	"go-employee-web-server/internal/configs"
	"go-employee-web-server/internal/data"
	"go-employee-web-server/internal/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Initialize database connection
	db, err := configs.GetDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Cannot init db connection")
		}
	}(db)
	baseURL := "https://dummy.restapiexample.com/api/v1/employees"
	redisClient := configs.GetRedisClient()
	storage := data.NewMySQLStorage(db)
	apiClient := api.NewHTTPClient(baseURL)
	handlerFactory := handlers.NewHandlerFactory(storage, apiClient, redisClient)

	http.HandleFunc("/", handlerFactory.MakeEmployeesHandler())
	http.HandleFunc("/add", handlerFactory.MakeAddHandler())
	http.HandleFunc("/view/", handlerFactory.MakeViewHandler())
	http.HandleFunc("/edit/", handlerFactory.MakeEditHandler())
	http.HandleFunc("/login/", handlerFactory.MakeLoginHandler())
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
