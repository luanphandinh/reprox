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

func getInfoFunction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-dbType", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server 2 with info"))
}

func main() {
	http.HandleFunc("/", getFunction)
	http.HandleFunc("/info", getInfoFunction)
	log.Println("Server 2 start in port 3001")

	http.ListenAndServe(":3001", nil)
}
