package main

import (
	"github.com/luanphandinh/reprox"
	"log"
	"net/http"
)

func main() {
	gateway := reprox.NewReprox()
	gateway.Register("server1", "http://localhost:3000/")
	gateway.Register("server2", "http://localhost:3001/")

	log.Println("Reverse Proxy server start in port 8080")

	http.HandleFunc("/", gateway.Handle())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
