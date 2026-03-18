package main

import (
	"fmt"
	"net/http"
)

func shipmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logistics Service Running")
}

func main() {
	http.HandleFunc("/shipment", shipmentHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}