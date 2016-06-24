package config

import (
	"encoding/json"
	"os"
)

// Engine contains the Application configuration
type Engine struct {
	Port          string `json:"port"`
	CacheLocation string `json:"cache_location"`
	BookLocation  string `json:"book_location"`
	Debug         bool   `json:"debug"`
}

// Reporter contains the configuration for Reporter component
type Reporter struct {
	Port  string `json:"port"`
	Debug bool   `json:"debug"`
}

// Database
type Database struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	User  string `json:"user"`
	Pass  string `json:"pass"`
	Name  string `json:"name"`
	Debug bool   `json:"debug"`
}

// BoltDB holds the BoltDB configuration
type BoltDB struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Name   string `json:"name"`
	Bucket string `json:"bucket"`
}

// Redis holds the redis configuration
type Redis struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

// Cache
type Cache struct {
	SymbolFile string `json:"symbol_file"`
	BoltDB     BoltDB `json:"boltdb"`
	Redis      Redis  `json:"redis"`
	Debug      bool   `json:"debug"`
}

// Config contains the Stock Exchange configuration
type Config struct {
	Exchange string   `json:"exchange"`
	Engine   Engine   `json:"engine"`
	Reporter Reporter `json:"reporter"`
	Database Database `json:"database"`
	Cache    Cache    `json:"cache"`
}

// Load builds a config obj
func Load(cf string) (*Config, error) {
	confFile, err := os.Open(cf)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(confFile)
	var conf Config
	if err = decoder.Decode(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
