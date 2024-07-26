package handlers

import (
	"fmt"
	"log"
	"os"
)

func Cleanup() {
	err := os.Remove("web/data/employees.txt")
	if err != nil {
		log.Printf("Error removing employees.txt: %v", err)
	} else {
		fmt.Println("employees.txt removed successfully")
	}
}
