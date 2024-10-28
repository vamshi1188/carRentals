package main

import (
    "log"
    "net/http"
    "./handlers"  // Use relative import for local package
)

func main() {
    http.HandleFunc("/cars", handlers.GetCars)
    http.HandleFunc("/create-rental", handlers.CreateRental)
    http.HandleFunc("/process-payment", handlers.ProcessPayment)

    log.Println("Starting server on :8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("Server failed to start:", err)
    }
}
