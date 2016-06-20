package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/briandowns/stock-exchange/models"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/thoas/stats"
	"github.com/unrolled/render"
)

var signalsChan = make(chan os.Signal, 1)

func main() {
	signal.Notify(signalsChan, os.Interrupt)

	go func() {
		for sig := range signalsChan {
			fmt.Printf("\nEngine shutting down... %v\n", sig)
			signalsChan = nil
			os.Exit(1)
		}
	}()

	sc, err := NewSymbolCache()
	if err != nil {
		log.Fatal(err)
	}
	cacher := Cacher(sc)
	cacher.Build()

	ren := render.New()

	ob := models.NewOrderBook()

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)

	statsMiddleware := stats.New()

	// create the router
	router := mux.NewRouter()

	// route handler for a health check
	router.HandleFunc(HealthCheckPath, HealthCheckHandler()).Methods("HEAD")

	// route handler for statistics
	router.HandleFunc(StatsPath, StatsHandler(ren, statsMiddleware)).Methods("GET")

	// route handler for the book
	router.HandleFunc(BookPath, BookHandler(ren, ob)).Methods("GET")

	// route handler for individual book entries
	router.HandleFunc(BookEntryByIDPath, BookEntryByIDHandler(ren)).Methods("GET")

	// route handler for viewing symbol data
	router.HandleFunc(SymbolsPath, SymbolsHandler(ren, cacher)).Methods("GET")

	// route handler for viewing symbol data by ID
	router.HandleFunc(SymbolByIDPath, SymbolByIDHandler(ren, cacher)).Methods("GET")

	n.Use(statsMiddleware)
	n.UseHandler(router)
	n.Run(":7777")
}
