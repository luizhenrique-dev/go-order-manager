package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/luizhenriquees/go-order-manager/internal/entity"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/order", CreateOrderAndReturn)
	http.ListenAndServe(":8888", r)
}

func CreateOrderAndReturn(w http.ResponseWriter, r *http.Request) {
	order, err := entity.NewOrder("1", 100, 10)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	order.CalculateFinalPrice()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
