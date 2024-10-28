package handlers

import (
	"encoding/json"
	"net/http"
)

type Car struct {
	ID    string `json:"id"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	Price int64  `json:"price"`
	Image string `json:"image"`
}

var cars = []Car{
	{ID: "1", Make: "Toyota", Model: "Camry", Year: 2022, Price: 10000, Image: "/images/camry.jpg"},
	{ID: "2", Make: "Honda", Model: "Accord", Year: 2021, Price: 12000, Image: "/images/accord.jpg"},
	{ID: "3", Make: "Honda", Model: "Civic", Year: 2023, Price: 14000, Image: "/images/civic.jpg"},
}

func GetCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}
