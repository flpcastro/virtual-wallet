package main

import (
	"net/http"
	"os"

	"github.com/flpcastro/virtual-wallet/internal/money"
)

func main() {
	http.HandleFunc("/users/", money.UsersHandler)
	http.HandleFunc("/transfer/", money.TranferHandler)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "3000"
	}

	http.ListenAndServe(":"+port, nil)
}
