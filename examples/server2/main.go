package main

import (
	"log"
	"net/http"
)

func getFunction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-dbType", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server 2"))
}

func main() {
	http.HandleFunc("/", getFunction)
	log.Println("Server 2 start in port 3001")

	http.ListenAndServe(":3001", nil)
}
