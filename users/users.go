package users

import "database/sql"

type User struct {
	ID      int
	Name    string
	Balance float64
}

func CreateUser(db *sql.DB, user User) error {
	query := `INSERT INTO users (name, balance) VALUES ($1, $2)`
	_, err := db.Exec(query, user.Name, user.Balance)
	return err
}
