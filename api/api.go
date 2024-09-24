package api

import (
	"car_market/cars"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateModelCarHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var carModel cars.Car
		if err := json.NewDecoder(r.Body).Decode(&carModel); err != nil {
			http.Error(w, "Неверный формат данных", http.StatusBadRequest)
			return
		}
		err := cars.CreateModelCar(db, carModel.Brand, carModel.EngineVolume, carModel.Color, carModel.WheelPosition)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Модель автомобиля успешно создана")
	}
}

func BuyCarHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Извлекаем user_id из URL
		vars := mux.Vars(r)
		userID, err := strconv.Atoi(vars["userID"])
		if err != nil {
			http.Error(w, "Неверный ID пользователя", http.StatusBadRequest)
			return
		}

		// Извлекаем модель машины из тела запроса
		var carRequest struct {
			ModelID int `json:"model_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&carRequest); err != nil {
			http.Error(w, "Неверный формат данных", http.StatusBadRequest)
			return
		}

		// Вызов функции покупки автомобиля
		err = cars.BuyCar(db, userID, carRequest.ModelID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Автомобиль успешно куплен")
	}
}

func SellCarHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Извлекаем user_id из URL
		vars := mux.Vars(r)
		userID, err := strconv.Atoi(vars["userID"])
		if err != nil {
			http.Error(w, "Неверный ID пользователя", http.StatusBadRequest)
			return
		}

		// Извлекаем ID машины из тела запроса
		var carRequest struct {
			CarID int `json:"car_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&carRequest); err != nil {
			http.Error(w, "Неверный формат данных", http.StatusBadRequest)
			return
		}

		// Вызов функции продажи автомобиля
		err = cars.SellCar(db, userID, carRequest.CarID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Автомобиль успешно продан")
	}
}

// UpdateCarPriceHandler – обработчик для изменения цены автомобиля
func UpdateCarPriceHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Извлекаем model_id из URL
		vars := mux.Vars(r)
		modelID, err := strconv.Atoi(vars["modelID"])
		if err != nil {
			http.Error(w, "Неверный ID модели", http.StatusBadRequest)
			return
		}

		// Извлекаем новую цену из тела запроса
		var priceRequest struct {
			Price float64 `json:"price"`
		}
		if err := json.NewDecoder(r.Body).Decode(&priceRequest); err != nil {
			http.Error(w, "Неверный формат данных", http.StatusBadRequest)
			return
		}

		// Вызов функции обновления цены автомобиля
		err = cars.UpdateCarPrice(db, modelID, priceRequest.Price)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Цена автомобиля успешно обновлена")
	}
}

func GetCarsByUserIDHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Извлекаем user_id из URL
		vars := mux.Vars(r)
		userID, err := strconv.Atoi(vars["userID"])
		if err != nil {
			http.Error(w, "Неверный ID пользователя", http.StatusBadRequest)
			return
		}

		// Вызов функции получения машин пользователя
		userCars, err := cars.GetCarsByUserID(db, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Возвращаем список машин пользователя в формате JSON
		json.NewEncoder(w).Encode(userCars)
	}
}

// GetCarsHandler – обработчик для получения всех автомобилей с фильтром по марке и сортировкой по цене
func GetCarsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Извлекаем параметры фильтра и сортировки из URL
		brand := r.URL.Query().Get("brand")
		sort := r.URL.Query().Get("sort")

		// Вызов функции получения списка машин
		carsList, err := cars.GetCars(db, sort, brand)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Возвращаем список машин в формате JSON
		json.NewEncoder(w).Encode(carsList)
	}
}
