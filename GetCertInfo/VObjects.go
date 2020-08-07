package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// VObject on kontrolliobjekt.
type VObject struct {
	Name string `json:"Name"`
	URL  string `json:"URL"`
}

// VObjects on kontrolliobjektid (nende URL-id).
type VObjects struct {
	VObjects []VObject `json:"VObjects"`
}

// loadVObjects loeb JSON-failist f sisse kontrolli-URL-id.
func loadVObjects(f string) VObjects {
	var d VObjects // Dekodeeritud kontrolliobjektid

	fmt.Printf("loadVObjects:\n--- Loen kontrolli-URL-id\n")

	// Ava fail
	fh, err := os.Open(f) // File handle
	if err != nil {
		fmt.Printf("    Kontrolli-URL-de lugemine ebaõnnestus. %s\n", err.Error())
		os.Exit(1)
	}

	defer fh.Close()
	// Dekodeeri JSON-struktuur
	jsonParser := json.NewDecoder(fh)
	err = jsonParser.Decode(&d)
	if err != nil {
		fmt.Println("    Kontrolliobjektide dekodeerimine ebaõnnestus.")
		os.Exit(1)
	}
	fmt.Println("    Loetud kontrolliobjektid:")
	for _, obj := range d.VObjects {
		fmt.Printf("    %s: %s\n", obj.Name, obj.URL)
	}
	return d
}
