package main

import (
	"log"
	"testing"

	"github.com/briandowns/stock-exchange/config"
	"github.com/briandowns/stock-exchange/database"

	"github.com/boltdb/bolt"
)

var testingDB = "data/testing_db"
var bConf *config.Config
var testDB *bolt.DB

// init
func init() { bSetup() }

// bSetup
func bSetup() {
	var err error
	bConf, err = config.Load(configFile)
	if err != nil {
		log.Fatal(err)
	}
	testDB, err = database.NewDB(testingDB)
	if err != nil {
		log.Fatal(err)
	}
}

// TestNewBoltCache
func TestNewBoltCache(t *testing.T) {
	NewBoltCache(testDB, bConf)
}

// TestBuild
func TestBuildBolt(t *testing.T) {
	bc := NewBoltCache(testDB, bConf)
	if err := bc.Build(); err != nil {
		t.Error(err)
	}
}

// TestGet
func TestGetBolt(t *testing.T) {
}

// TestAdd adds a value to the cache
func TestAddBolt(t *testing.T) {
}

// TestEntries
func TestEntriesBolt(t *testing.T) {
}
