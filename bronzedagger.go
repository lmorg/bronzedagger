// Bronze Dagger
package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	AppName   = "BronzeDagger"
	Version   = "3.04.0620"
	Copyright = "© 2014-2016 Laurence Morgan"
)

var (
	uiPass string
	uiFail string
)

func init() {
	results = make(map[int]int, 10000)
	loadTimes = make(map[int]int, 10000)
}

func main() {
	flags()

	job := NewJob()

	//config := readConfig(conf_filename)

	// sigterm to quit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		summary()
		os.Exit(2)
	}()

	// start results event loop
	go updateResults()
	go StartFSLog()

	// conf file will be supported again at some point in the future
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
	if fConcurrentUrls {
		var wg sync.WaitGroup
		for _, url := range fUrls {
			wg.Add(1)
			fork := job.Fork(url)
			go fork.Start()
		}
		wg.Wait()
	} else {
		for _, url := range fUrls {
			fork := job.Fork(url)
			fork.Start()
		}
	}
	//}

	debugLog("All routines ended. Generating summary")
	summary()
	os.Exit(0)
}
