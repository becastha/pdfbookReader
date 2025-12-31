package main

import (
	"log"
	"net/http"

	"github.com/becas/pdfbookReader/backend/internal/routes"
)

func main() {

	router := routes.NewRouter()
	log.Println("Server started in the port 8080")
	http.ListenAndServe(":8080", router)
}
