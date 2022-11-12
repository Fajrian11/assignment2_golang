package main

import (
	"assignment2_golang/router"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := router.StartAPP()
	port := os.Getenv("SERVICE_PORT")
	r.Run(port)
}
