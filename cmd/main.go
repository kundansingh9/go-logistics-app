package main

import (
	"fmt"
	"net/http"
)

// CORS Middleware
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next(w, r)
	}
}

func shipmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logistics Service Running")
}

func main() {
	http.HandleFunc("/shipment", enableCORS(shipmentHandler))

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func shipmentHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Logistics Service Running")
// }

// func main() {
// 	http.HandleFunc("/shipment", shipmentHandler)
// 	fmt.Println("Server started at :8080")
// 	http.ListenAndServe(":8080", nil)
// }
