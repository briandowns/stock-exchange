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

	OrderPath       = APIBase + "order"
	CancelOrderPath = OrderPath + ByID
)

// HealthCheckHandler handles health checking
func HealthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, http.StatusOK)
	}
}

// StatsHandler handles API stats processing
func StatsHandler(ren *render.Render, statsMW *stats.Stats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statsData := statsMW.Data()
		ren.JSON(w, http.StatusOK, statsData)
	}
}

// SymbolsHandler retrieves all tradable symbols
func SymbolsHandler(ren *render.Render, cacher Cacher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cd, err := cacher.Entries()
		if err != nil {
			log.Println(err)
			ren.JSON(w, http.StatusOK, map[string]interface{}{"error": err})
			return
		}
		ren.JSON(w, http.StatusOK, map[string]interface{}{"symbols": cd})
	}
}

// SymbolByIDHandler retrieves a symbol and data by ID
func SymbolByIDHandler(ren *render.Render, cacher Cacher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		symbolID := vars["id"]
		data, err := cacher.Get([]byte(symbolID))
		if err != nil {
			log.Println(err)
			ren.JSON(w, http.StatusOK, map[string]interface{}{"error": err})
			return
		}
		ren.JSON(w, http.StatusOK, map[string]interface{}{"symbol": data})
	}
}

// BookHandler retrieves the current book
func BookHandler(ren *render.Render, ob *models.OrderBook) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ren.JSON(w, http.StatusOK, ob)
	}
}

// BookEntryByIDHandler retrieves an entry in the book
func BookEntryByIDHandler(ren *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookID := vars["id"]
		ren.JSON(w, http.StatusOK, bookID)
	}
}

// AddOrderHandler adds a new trade
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

// CancelTradeHandler will receive requests to cancel pending trades
func CancelTradeHandler(ren *render.Render, ob *models.OrderBook) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		orderID := vars["id"]
		ob.Cancel(orderID)
		ren.JSON(w, http.StatusOK, map[string]interface{}{"id": orderID})
	}
}
