package cars

import (
	"database/sql"
	"fmt"
)

type Car struct {
	ID            int
	EngineVolume  float64
	Color         string
	Brand         string
	WheelPosition string
	Price         float64
}

func CreatModelCar(db *sql.DB, brand string, engineVolume float64, color string, wheelPosition string) error {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM car_models WHERE brand brand = $1 AND engine_volume = $2 AND color = $3 AND wheel_position a= $4)",
		brand, engineVolume, color, wheelPosition).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("модель существует")
	}

	query := `INSERT INTO car_models (brand, engine_volume, color, wheel_position) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(query, brand, engineVolume, color, wheelPosition)
	return err
}

func BuyCar(db *sql.DB, userID int, carID int) error {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM car_models WHERE id = $1)", carID).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("модель автомобиля не существует")
	}
	// начало транзакции
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var price float64
	err = tx.QueryRow("SELECT price FROM cars WHERE id = $1 AND owner_id IS NULL", carID).Scan(&price)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE users SET balance = balance - $1 WHERE id = $2", price, userID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE cars SET owner_id = $1 WHERE id = $2", userID, carID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func SellCar(db *sql.DB, userID int, carID int) error {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM user_cars WHERE id = $1)", carID).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("модель автомобиля не существует")
	}
	// начало транзакции
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var price float64
	err = tx.QueryRow("SELECT price FROM cars WHERE id = $1", carID).Scan(&price)
	if err != nil {
		return err
	}

	// вернем 50% стоимости пользователю
	refund := price * 0.5
	_, err = tx.Exec("UPDATE users SET balance = balance + $1 WHERE id = $2", refund, userID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE cars SET owner_id = NULL WHERE id = $1", carID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func UpdateCarPrice(db *sql.DB, carID int, newPrice float64) error {
	_, err := db.Exec("UPDATE cars SET price = $1 WHERE id = $2", newPrice, carID)
	if err != nil {
		return err
	}
	return err
}
