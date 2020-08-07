package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config on seadistuse tüüp.
type Config struct {
	ClientCert string `json:"clientCert"`
	ClientKey  string `json:"clientKey"`
	RootCA     string `json:"rootCA"`
	// Kontrolliobjektide fail
	VObjectsFile string `json:"vObjectsFile"`
}

// Sisseloetud seadistus.
var conf Config

// loadConf loeb JSON-failist f sisse seadistuse.
func loadConf(f string) Config {
	fmt.Printf("loadConf:\n--- Loen seadistuse failist: %v\n", f)
	var config Config
	configFile, err := os.Open(f)
	defer configFile.Close()
	if err != nil {
		fmt.Printf("    Seadistuse lugemine ebaõnnestus:\n    %s\n", err.Error())
		os.Exit(1)
	}
	// Dekodeeri JSON-struktuur
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		fmt.Printf("    Seadistuse dekodeerimine ebaõnnestus:\n    %s\n", err)
		os.Exit(1)
	}
	// Kuva konf-n
	fmt.Printf("    Seadistus laetud:\n    %#v\n", config)
	return config
}
