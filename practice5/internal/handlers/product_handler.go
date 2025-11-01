package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"practice5/internal/repository"
)

type ProductHandler struct {
	repo *repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	filters := map[string]string{
		"category":  q.Get("category"),
		"min_price": q.Get("min_price"),
		"max_price": q.Get("max_price"),
	}
	sort := q.Get("sort")

	limit := 10
	offset := 0
	if v := q.Get("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			limit = n
		}
	}
	if v := q.Get("offset"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			offset = n
		}
	}

	products, queryTime, err := h.repo.GetProducts(filters, limit, offset, sort)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Query-Time", queryTime.String())
	json.NewEncoder(w).Encode(products)
}
