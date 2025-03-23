// main.go
// curl -X POST http://localhost:4001/form -d "blogName=Big Mike's Blog&blogContent=Big Mike's Adventures"

package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /form", formHandler)

	log.Println("Starting HTTP Server on port 4001...")
	err := http.ListenAndServe(":4001", mux)
	if err != nil {
		log.Fatal("Error occurred while starting the server:", err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parse form data

	blogName := r.FormValue("blogName")
	blogContent := r.FormValue("blogContent")

	log.Printf("Blog Name: %s\n", blogName)
	log.Printf("Blog Content: %s\n", blogContent)
}
