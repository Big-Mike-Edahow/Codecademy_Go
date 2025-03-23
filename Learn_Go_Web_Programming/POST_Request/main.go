// main.go
// curl -X POST http://localhost:4001/data -H "Content-Type: application/json" -d '{"productId": 123456, "quantity": 100}'

package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /data", handlePostRequest)

    log.Println("Starting HTTP Server on port 4001...")
	err := http.ListenAndServe(":4001", mux)
	if err != nil {
		log.Fatal("Error occurred while starting the server: ", err)
	}
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	log.Printf("Received JSON: %s", string(body))
}
