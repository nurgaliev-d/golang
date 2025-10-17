package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID      int     `db:"id"`
	Name    string  `db:"name"`
	Email   string  `db:"email"`
	Balance float64 `db:"balance"`
}

func InsertUser(db *sqlx.DB, user User) error {
	_, err := db.NamedExec(`INSERT INTO users (name, email, balance)
                            VALUES (:name, :email, :balance)`, user)
	return err
}

func GetAllUsers(db *sqlx.DB) ([]User, error) {
	var users []User
	err := db.Select(&users, "SELECT * FROM users")
	return users, err
}

func GetUserByID(db *sqlx.DB, id int) (User, error) {
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	return user, err
}

func TransferBalance(db *sqlx.DB, fromID int, toID int, amount float64) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	var fromUser, toUser User

	if err := tx.Get(&fromUser, "SELECT * FROM users WHERE id=$1", fromID); err != nil {
		tx.Rollback()
		return fmt.Errorf("sender not found: %w", err)
	}

	if err := tx.Get(&toUser, "SELECT * FROM users WHERE id=$1", toID); err != nil {
		tx.Rollback()
		return fmt.Errorf("receiver not found: %w", err)
	}

	if fromUser.Balance < amount {
		tx.Rollback()
		return fmt.Errorf("insufficient funds")
	}

	_, err = tx.Exec("UPDATE users SET balance = balance - $1 WHERE id=$2", amount, fromID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update sender: %w", err)
	}

	_, err = tx.Exec("UPDATE users SET balance = balance + $1 WHERE id=$2", amount, toID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update receiver: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit failed: %w", err)
	}

	return nil
}

func main() {
	db, err := sqlx.Open("postgres", "user=user password=password dbname=mydatabase port=5430 sslmode=disable")
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	fmt.Println("✅ Connected to PostgreSQL!")

	// Example: Insert new users (run once)
	/*
		newUser := User{Name: "Alice", Email: "alice@example.com", Balance: 200.0}
		if err := InsertUser(db, newUser); err != nil {
			log.Println("Insert error:", err)
		}
	*/
	newUser := User{Name: "Bob", Email: "bob@example.com", Balance: 150.0}
	if err := InsertUser(db, newUser); err != nil {
		log.Println("Insert error:", err)
	}

	users, err := GetAllUsers(db)
	if err != nil {
		log.Println("Error fetching users:", err)
	} else {
		fmt.Println("Users in database:")
		for _, u := range users {
			fmt.Printf("ID=%d | Name=%s | Email=%s | Balance=%.2f\n", u.ID, u.Name, u.Email, u.Balance)
		}
	}

	fmt.Println("\n--- Simulating Transfer ---")
	err = TransferBalance(db, 1, 2, 50.0)
	if err != nil {
		fmt.Println("❌ Transfer failed:", err)
	} else {
		fmt.Println("💸 Transfer successful!")
	}

	fmt.Println("\nUpdated user balances:")
	updatedUsers, _ := GetAllUsers(db)
	for _, u := range updatedUsers {
		fmt.Printf("ID=%d | Name=%s | Balance=%.2f\n", u.ID, u.Name, u.Balance)
	}
}
