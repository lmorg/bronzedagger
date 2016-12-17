package main

import (
	"fmt"
	"sort"
)

var (
	chanUpdateResults chan Response

	results   map[int]int
	loadTimes map[int]int
)

type Response struct {
	Status   int
	Duration int
}

func init() {
	chanUpdateResults = make(chan Response)
}

func updateResults() {
	for {
		r := <-chanUpdateResults
		results[r.Status]++
		loadTimes[r.Duration]++
	}
}

func summary() {
	if fNoSummary {
		return
	}

	fmt.Println("\n        #  Status  Description")
	// sort
	var keyResults []int
	for key := range results {
		keyResults = append(keyResults, key)
	}
	sort.Ints(keyResults)
	// output
	for _, key := range keyResults {
		fmt.Printf("%9d  %03d     %s\n", results[key], key, status_code[key])
	}

	// response timings

	fmt.Println("\n        #  Milliseconds")
	// sort
	var keyLoadTimes []int
	for key := range loadTimes {
		keyLoadTimes = append(keyLoadTimes, key)
	}
	sort.Ints(keyLoadTimes)
	// output
	var nOverMax int
	for _, key := range keyLoadTimes {
		if key < fMaxDisplayedTime {
			fmt.Printf("%9d  %d - %d\n", loadTimes[key], key, key+fRounding-1)
		} else {
			nOverMax += loadTimes[key]
		}
	}
	if nOverMax != 0 {
		fmt.Printf("%9d  > %d\n", nOverMax, fMaxDisplayedTime)
	}
}
