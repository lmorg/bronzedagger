package main

import (
	"fmt"
	"sort"
)

var (
	update_results chan Response

	results    map[int]int
	load_times map[int]int
)

type Response struct {
	Status   int
	Duration int
}

func init() {
	update_results = make(chan Response)
}

func UpdateResults() {
	for {
		r := <-update_results
		results[r.Status]++
		load_times[r.Duration]++
	}
}

func summary() {
	fmt.Println("\n        #  Status  Description")
	// sort
	var k_results []int
	for k := range results {
		k_results = append(k_results, k)
	}
	sort.Ints(k_results)
	// output
	for _, k := range k_results {
		fmt.Printf("%9d  %03d     %s\n", results[k], k, status_code[k])
	}

	// response timings

	fmt.Println("\n        #  Milliseconds")
	// sort
	var k_load_times []int
	for k := range load_times {
		k_load_times = append(k_load_times, k)
	}
	sort.Ints(k_load_times)
	// output
	var n_over_max int
	for _, k := range k_load_times {
		if k < display_max {
			fmt.Printf("%9d  %d - %d\n", load_times[k], k, k+f_rounding-1)
		} else {
			n_over_max += load_times[k]
		}
	}
	if n_over_max != 0 {
		fmt.Printf("%9d  > %d\n", n_over_max, display_max)
	}
}
