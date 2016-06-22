package config

import (
	"encoding/json"
	"os"
)

//Engine contains the Application configuration
type Engine struct {
	Port          string `json:"port"`
	Debug         bool   `json:"debug"`
	CacheLocation string `json:"cache_location"`
	BookLocation  string `json:"book_location"`
}

//Book contains the configuration for Book component
type Book struct {
	Port  string `json:"port"`
	Debug bool   `json:"debug"`
}

//Reporter contains the configuration for Reporter component
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

//Cache
type Cache struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	User  string `json:"user"`
	Pass  string `json:"pass"`
	Debug bool   `json:"debug"`
}

// Config contains the Stock Exchange configuration
type Config struct {
	Engine   Engine   `json:"engine"`
	Book     Book     `json:"book"`
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
