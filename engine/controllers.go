package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/briandowns/stock-exchange/models"
	"github.com/gorilla/mux"
	"github.com/thoas/stats"
	"github.com/unrolled/render"
)

const (
	ByID = "/{id}"

	// APIBase is the base path for API access
	APIBase = "/api/v1/"

	BookPath   = APIBase + "book"
	SymbolPath = APIBase + "symbols"

	HealthCheckPath = APIBase + "healthcheck"
	StatsPath       = APIBase + "stats"
	BookByIDPath    = "book" + ByID
	SymbolsByIDPath = "symbol" + ByID
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
		cd, err := cacher.All()
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
