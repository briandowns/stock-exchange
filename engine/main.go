package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/thoas/stats"
	"github.com/unrolled/render"
)

const (
	// APIBase is the base path for API access
	APIBase = "/api/v1/"
)

var signalsChan = make(chan os.Signal, 1)

func main() {
	signal.Notify(signalsChan, os.Interrupt)

	go func() {
		for sig := range signalsChan {
			log.Printf("Exiting... %v\n", sig)
			signalsChan = nil
			os.Exit(1)
		}
	}()

	ren := render.New()

	//ob := models.NewOrderBook()
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)

	statsMiddleware := stats.New()

	// create the router
	router := mux.NewRouter()

	// route handler for a health check
	router.HandleFunc(APIBase+"healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, http.StatusOK)
	}).Methods("HEAD")

	// route handler for statistics
	router.HandleFunc(APIBase+"stats", func(w http.ResponseWriter, r *http.Request) {
		stats := statsMiddleware.Data()
		ren.JSON(w, http.StatusOK, stats)
	}).Methods("GET")

	// route handler for the book
	router.HandleFunc(APIBase+"book", func(w http.ResponseWriter, r *http.Request) {
		ren.JSON(w, http.StatusOK, "book")
	}).Methods("GET")

	// route handler for individual book entries
	router.HandleFunc(APIBase+"book/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookID := vars["id"]
		ren.JSON(w, http.StatusOK, bookID)
	}).Methods("GET")

	n.Use(statsMiddleware)
	n.UseHandler(router)
	n.Run(":7777")
}
