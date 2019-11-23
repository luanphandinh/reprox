package main

import (
	"log"
	"net/http"
)

func getFunction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-dbType", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server 1"))
}

func main() {
	http.HandleFunc("/", getFunction)
	log.Println("Server 1 start in port 3000")

	http.ListenAndServe(":3000", nil)
}
