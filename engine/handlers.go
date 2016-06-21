package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/briandowns/stock-exchange/models"
	"github.com/gorilla/mux"
	"github.com/pborman/uuid"
	"github.com/thoas/stats"
	"github.com/unrolled/render"
)

const (
	ByID = "/{id}"

	// APIBase is the base path for API access
	APIBase = "/api/v1/"

	HealthCheckPath = APIBase + "healthcheck"
	StatsPath       = APIBase + "stats"

	BookPath          = APIBase + "book"
	BookEntryByIDPath = APIBase + "book" + ByID

	SymbolByIDPath = APIBase + "symbol" + ByID
	SymbolsPath    = APIBase + "symbols"

	OrderPath = APIBase + "order"
)

// HealthCheckHandler
func HealthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, http.StatusOK)
	}
}

// StatsHandler
func StatsHandler(ren *render.Render, statsMW *stats.Stats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statsData := statsMW.Data()
		ren.JSON(w, http.StatusOK, statsData)
	}
}

// SymbolsHandler
func SymbolsHandler(ren *render.Render, cacher Cacher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cd, err := cacher.Entries()
		if err != nil {
			log.Fatalln(err)
		}
		ren.JSON(w, http.StatusOK, map[string]interface{}{"symbols": cd})
	}
}

// SymbolByIDHandler
func SymbolByIDHandler(ren *render.Render, cacher Cacher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		symbolID := vars["id"]
		data, err := cacher.Get([]byte(symbolID))
		if err != nil {
			ren.JSON(w, http.StatusOK, map[string]interface{}{"error": "symbol not found"})
			return
		}
		ren.JSON(w, http.StatusOK, map[string]interface{}{"symbol": data})
	}
}

// BookHandler
func BookHandler(ren *render.Render, ob *models.OrderBook) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ren.JSON(w, http.StatusOK, ob)
	}
}

// BookEntryByIDHandler
func BookEntryByIDHandler(ren *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookID := vars["id"]
		ren.JSON(w, http.StatusOK, bookID)
	}
}

// AddOrderHandler
func AddOrderHandler(ren *render.Render, ob *models.OrderBook) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order models.Order
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&order); err != nil {
			ren.JSON(w, http.StatusInternalServerError, map[string]interface{}{"error": err})
			return
		}
		order.ID = uuid.NewUUID().String()
		if err := ob.Add(order); err != nil {
			ren.JSON(w, http.StatusInternalServerError, map[string]interface{}{"error": err})
			return
		}
		ren.JSON(w, http.StatusCreated, order)
	}
}

// CancelTradeHandler
func CancelTradeHandler(ren *render.Render, ob *models.OrderBook) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		orderID := vars["id"]
		ren.JSON(w, http.StatusOK, map[string]interface{}{"id": orderID})
	}
}
