package main

import (
	"log"
	"testing"

	"github.com/briandowns/stock-exchange/config"
)

var rConf *config.Config

func init() { setup() }

func setup() {
	var err error
	rConf, err = config.Load(configFile)
	if err != nil {
		log.Fatal(err)
	}
}

// TestNewRedisCache
func TestNewRedisCache(t *testing.T) {
	NewRedisCache(rConf)
}

// TestBuild
func TestBuild(t *testing.T) {
	rc := NewRedisCache(rConf)
	if err := rc.Build(); err != nil {
		t.Error(err)
	}
}

// TestGet
func TestGet(t *testing.T) {
}

// TestAdd
func TestAdd(t *testing.T) {
}

// TestEntries
func TestEntries(t *testing.T) {
}

// TestFlushDB
func TestFlushDB(t *testing.T) {
}
