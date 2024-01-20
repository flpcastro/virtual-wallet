package main

import (
	"net/http"
	"os"

	"github.com/flpcastro/virtual-wallet/internal/money"
)

func main() {
	http.HandleFunc("/users/", money.UsersHandler)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "3000"
	}

	http.ListenAndServe(":"+port, nil)
}
