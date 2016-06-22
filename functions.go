package main

import (
	//"io/ioutil"
	"regexp"
	//"strings"
)

var (
	rx_tsv *regexp.Regexp
)

func init() {
	rx_tsv = regexCompile(`\s+`)
}

func regexCompile(s string) *regexp.Regexp {
	r, err := regexp.Compile(s)
	isErr(err)
	return r
}

func Lower(val int) int {
	return int(val/f_rounding) * f_rounding
}

/*
func readConfig(filename string) (config [][]string) {
	if filename == "" {
		return
	}

	b, err := ioutil.ReadFile(filename)
	if !isErr(err) {
		lines := strings.Split(string(b), "\n")
		for _, l := range lines {
			config = append(config, rx_tsv.Split(l, -1))
		}
	}

	return
}
*/
