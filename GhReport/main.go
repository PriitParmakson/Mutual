/*
GhReport jalutab läbi kausta (ja kõik selle alamkaustad), parsib .feature failid
ja koostab nimekirja Gherkinis (kergkeeles) spetsifitseeritud omadustest (feature)
ja stsenaariumitest (scenario). Iga stsenaariumi kohta väljastab sammude arvu.
*/
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	messages "github.com/cucumber/cucumber-messages-go/v8"
	gherkin "github.com/cucumber/gherkin-go/v9"
	"github.com/fatih/color"
)

type a gherkin.GherkinDialect

func main() {
	rootFolder := flag.String("dir", ".", ".feature failide kaust")
	flag.Parse()

	color.Cyan("ARUANNE")
	color.Cyan(*rootFolder)

	// Töödeldud faile
	nof := 0

	// Jaluta läbi kaustapuu
	err := filepath.Walk(
		*rootFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// Töötle fail
			// Ei ole kaust
			if !info.IsDir() &&
				// on .feature fail
				filepath.Ext(path) == ".feature" {
				parsiKergkeeleTekst(path)
				nof++
			}
			return nil
		},
	)

	if err != nil {
		fmt.Printf("Viga kausta läbijalutamisel: %v\n", err)
		os.Exit(1)
	}

	// Väljasta statistika
	fmt.Printf("\nFaile: %v\n", nof)

}

func parsiKergkeeleTekst(path string) {
	// Ava fail
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var r io.Reader
	r = f

	// Parsi kergkeeledokument
	gherkinDocument, err := gherkin.ParseGherkinDocument(r, (&messages.Incrementing{}).NewId)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
		return
	}

	// Kuva omadus
	feature := gherkinDocument.Feature
	fmt.Fprintf(os.Stdout, "Omadus: %+v\n", feature.Name)

	// Kuva kõigi stsenaariumite nimed
	for _, c := range feature.Children {
		scenario := c.GetScenario()
		if scenario != nil {
			fmt.Fprintf(os.Stdout, "    Testistsenaarium: %+v", scenario.Name)
			fmt.Fprintf(os.Stdout, " (%+v sammu)\n", len(scenario.Steps))
		}
	}
}

// Märkmed:
// Kaustapuu läbijalutamine: https://yourbasic.org/golang/list-files-in-directory/
