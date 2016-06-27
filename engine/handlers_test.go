package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thoas/stats"
	"github.com/unrolled/render"
)

var ren = render.New()

var testCache Cache

// TestHealthCheckHandler
func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", HealthCheckPath, nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// TestStatsHandler
func TestStatsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", StatsPath, nil)
	if err != nil {
		t.Error(err)
	}
	statsMiddleware := stats.New()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StatsHandler(ren, statsMiddleware))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// TestSymbolsHandler
func TestSymbolsHandler(t *testing.T) {
}

// TestSymbolByIDHandler
func TestSymbolByIDHandler(t *testing.T) {
}

// TestBookHandler
func TestBookHandler(t *testing.T) {
}

// TestBookEntryByIDHandler
func TestBookEntryByIDHandler(t *testing.T) {
}

// TestAddOrderHandler
func TestAddOrderHandler(t *testing.T) {
}

// TestCancelTradeHandler
func TestCancelTradeHandler(t *testing.T) {
}
