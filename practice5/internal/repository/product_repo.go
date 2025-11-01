package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"practice5/internal/models"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) GetProducts(filters map[string]string, limit, offset int, sort string) ([]models.Product, time.Duration, error) {
	start := time.Now()

	query := `
		SELECT p.id, p.name, c.name AS category, p.price
		FROM products p
		JOIN categories c ON p.category_id = c.id
	`
	var conditions []string
	var args []interface{}
	argID := 1

	if v, ok := filters["category"]; ok && v != "" {
		conditions = append(conditions, fmt.Sprintf("c.name = $%d", argID))
		args = append(args, v)
		argID++
	}
	if v, ok := filters["min_price"]; ok && v != "" {
		conditions = append(conditions, fmt.Sprintf("p.price >= $%d", argID))
		args = append(args, v)
		argID++
	}
	if v, ok := filters["max_price"]; ok && v != "" {
		conditions = append(conditions, fmt.Sprintf("p.price <= $%d", argID))
		args = append(args, v)
		argID++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	switch sort {
	case "price_asc":
		query += " ORDER BY p.price ASC"
	case "price_desc":
		query += " ORDER BY p.price DESC"
	default:
		query += " ORDER BY p.id ASC"
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price); err != nil {
			return nil, 0, err
		}
		products = append(products, p)
	}

	return products, time.Since(start), nil
}
