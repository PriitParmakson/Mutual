// This program merges authors.json files produced by git-of-theseus, so that LOC evolution
// chart can be produced over multiple repos.
// Output is written to file mergedAuthors.json.
//
// Usage:
//   go run . <input file>...
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"sort"
)

type author string
type loc int
type ts string
type repo string

// List of repos
var repos []repo

// repoLoc is the structure of input and output files.
type repoLOC struct {
	Labels []author `json:"labels"`
	Y      [][]loc  `json:"y"`
	Ts     []ts     `json:"ts"`
}

// repoState is an intermediate (helper) structure.
type repoState struct {
	repo  repo
	aLOCs map[author]loc
}

var repoStates map[ts]repoState
var mergedStates map[ts]map[author]loc

// List of authors over all repos
var allAuthors map[author]bool

// Author's code in repos
var authorsCodeInRepos map[repo]loc

// Every author's code in repos
var prevMergedState map[author]map[repo]loc

type orderedTssType []ts 
// orderedTssType implements sort.Interface based on field Ts. 
func (a orderedTssType) Len() int {
	return len(a)
}
func (a orderedTssType) Less(i, j int) bool {
	return a[i] < a[j]
}
func (a orderedTssType) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

var orderedTss orderedTssType

func main() {
	inFNames := os.Args[1:]

	// Initialize maps and slices.
	repos = make([]repo, 0)
	allAuthors = make(map[author]bool)
	repoStates = make(map[ts]repoState)

	// Read in and process input files:
	// Add repo's name to list of repos.
	// Add data from parsed JSON to repoStates.
	// Also, add all authors into allAuthors and all repos into repos.
	for _, inFName := range inFNames {
		// Add repo's name to list of repos.
		repos = append(repos, repo(inFName))
		// Read and parse JSON file.
		fmt.Printf("Processing input file: %v\n", inFName)
		f, err := os.Open(inFName)
		defer f.Close()
		if err != nil {
			fmt.Printf("Failure while reading input file:\n    %s\n", err.Error())
			os.Exit(1)
		}
		// Decode and parse JSON.
		jsonParser := json.NewDecoder(f)
		var nLOC repoLOC
		err = jsonParser.Decode(&nLOC)
		if err != nil {
			fmt.Printf("Failure in parsing JSON:\n    %s\n", err)
			os.Exit(1)
		}

		// Add data from parsed JSON to repoStates.
		nC := 0 // For running over timestamps.
		for {
			if nC == len(nLOC.Ts) {
				break
			}
			t := nLOC.Ts[nC] // Timestamp to be processed below
			// Create repoState
			var rs repoState
			rs.repo = repo(inFName)
			rs.aLOCs = make(map[author]loc)
			for i, a := range nLOC.Labels {
				rs.aLOCs[a] = nLOC.Y[i][nC]
				// Add author into list of authors.
				allAuthors[a] = true
			}
			// Add repoState
			repoStates[t] = rs
			nC++
		}
	}

	fmt.Printf("Repos:\n    %v\n", repos)
	fmt.Printf("Authors:\n")
	for a, _ := range allAuthors {
		fmt.Printf(" %v,", a)
	}
	fmt.Printf("\n")

	// Create ordered slice of timestamps, for processing repoStates
	// in sorted order.
	orderedTss = make([]ts, 0, len(repoStates))
	for k := range repoStates {
		orderedTss = append(orderedTss, k)
	}
	sort.Sort(orderedTssType(orderedTss))
	
	// Initialize prevMergedState
	prevMergedState = make(map[author]map[repo]loc)
	for a := range allAuthors {
		prevMergedState[a] = make(map[repo]loc)
		for _, r := range repos {
			prevMergedState[a][r] = 0
		}
	}
	
	mergedStates = make(map[ts]map[author]loc)
	// Create from repoStates mergedStates
	// For every repoState...
	// Pass through repoStates, using sorting keys
	for k := range orderedTss {
		t := orderedTss[k]
		rs := repoStates[t] 
		// Initialize time moment entry in mergedStates
		mergedStates[t] = make(map[author]loc)
		// For every author mentioned in repoState...
		for a, loc := range rs.aLOCs {
			// Update prevMergedState
			prevMergedState[a][rs.repo] = loc
		}
		// and then, for every author
		for a2 := range allAuthors {
			var locOverAllRepos loc
			// Sum author's code over all repos
			for _, loc2 := range prevMergedState[a2] {
				locOverAllRepos = locOverAllRepos + loc2
			}
			// Write sum into mergedStates
			mergedStates[t][a2] = locOverAllRepos
		}
	}

	// Output to console for debugging
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("Merged data:\n%v\n",
	// 	  mergedStates[orderedTss[i]])
	// } 

	// Moodusta väljund-JSON
	var out repoLOC
	// Create Labels and Y
	out.Labels = make([]author, 0)
	out.Y = make([][]loc, 0)
	var i int // Author number
	for a, _ := range allAuthors {
		out.Labels = append(out.Labels, a)
		// Create Y
		out.Y = append(out.Y, make([]loc, 0))
		// Fill in Y
		// Create LOC-vector; it must be sorted.
		var locV []loc
		locV = make([]loc, 0)
		for _, ts := range orderedTss {
			locV = append(locV, mergedStates[ts][a])
		}
		// and insert into Y
		out.Y[i] = locV
		i++
	}
	// create Ts
	out.Ts = make([]ts, 0)
	for _, ts := range orderedTss {
		out.Ts = append(out.Ts, ts)
	}

	// Write output structire (out) info JSON file
	file, err := json.MarshalIndent(out, "", " ")
	if err != nil {
		fmt.Printf("Failure in coding into JSON:\n    %s\n", err)		
	}
		
	err = ioutil.WriteFile("mergedAuthors.json", file, 0644)
	if err != nil {
		fmt.Printf("Failure in writing output file:\n    %s\n", err)		
	}

	fmt.Printf("\nOutput file mergedAuthors.json created\n\n")		

}	

// Notes:
// https://gobyexample.com/command-line-arguments

// https://yourbasic.org/golang/sort-map-keys-values/

// https://yourbasic.org/golang/how-to-sort-in-go/

// https://golang.org/pkg/sort/

// https://www.golangprograms.com/golang-writing-struct-to-json-file.html
