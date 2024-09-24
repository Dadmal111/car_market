package main

import (
	"car_market/api"
	"car_market/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()

	db := config.ConnectDB()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		log.Fatalf("Ошибка при подключении к БД: %v", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/cars", api.CreateModelCarHandler(db)).Methods("POST")
	r.HandleFunc("/users/{userID}/buy", api.BuyCarHandler(db)).Methods("POST")
	r.HandleFunc("/users/{userID}/sell", api.SellCarHandler(db)).Methods("POST")
	r.HandleFunc("/cars{carID}/price", api.UpdateCarPriceHandler(db)).Methods("PUT")
	r.HandleFunc("/users/{userID}/cars", api.GetCarsByUserIDHandler(db)).Methods("GET")
	r.HandleFunc("/cars", api.GetCarsHandler(db)).Methods("GET")

	//Запуск HTTP-сервера
	log.Println("Сервер запущен на порту 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
