// Bronze Dagger
package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

const (
	APP_NAME  = "BronzeDagger"
	VERSION   = "3.01.0444 ALPHA"
	COPYRIGHT = "© 2014-2015 Laurence Morgan"
)

var (
	f_no_200     bool
	f_rounding   int
	req_headers  bool
	resp_headers bool
	resp_body    bool
	display_max  int
	return_val   int
	TICK         string
	CROSS        string
	USER_AGENT   string
)

func init() {
	results = make(map[int]int, 10000)
	load_times = make(map[int]int, 10000)
	USER_AGENT = strings.Replace(fmt.Sprintf("%s/%s", APP_NAME, VERSION), " ", "", -1)
}

func main() {
	Flags()

	job := NewJob()

	//config := readConfig(conf_filename)

	// sigterm to quit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		summary()
		os.Exit(return_val)
	}()

	// start results event loop
	go UpdateResults()
	go StartFSLog()

	/*if conf_filename != "" {
		for i := 0; i < len(config); i++ {
			switch config[i][0] {
			case "cookie":
				//cookies = append(cookies, addCookie(config[i]))
			case "url":
				loopRequests(config[i][1], cookies, nil)
			case "t":
				seconds, _ = strconv.Atoi(config[i][1])
			case "c":
				concurrency, _ = strconv.Atoi(config[i][1])
			}
		}

	} else {*/
	var wg sync.WaitGroup
	for _, url := range f_urls {
		wg.Add(1)
		fork := job.Fork(url)
		go fork.Start(&wg)
	}
	wg.Wait()
	//}

	debugLog("All routines ended. Generating summary")
	summary()
	os.Exit(return_val)
}
