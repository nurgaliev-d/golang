package main

import (
	"database/sql"
	"log"
	"net/http"

	"practice5/internal/handlers"
	"practice5/internal/repository"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost user=diasnurgaliev password=secret dbname=productsdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	productRepo := repository.NewProductRepository(db)
	productHandler := handlers.NewProductHandler(productRepo)

	http.HandleFunc("/products", productHandler.GetProducts)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
