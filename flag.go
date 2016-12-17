package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

var (
	fReqHeaders       bool
	fRespHeaders      bool
	fRespBody         bool
	fUserAgent        string
	fNo200            bool
	fRounding         int
	fNoSummary        bool
	fReferrer         string
	fRedirects        bool
	fDuration         int
	fConcurrency      int
	fInsecure         bool
	fTimeout          int64
	fNReqs            int
	fFsLog            string
	fDebug            bool
	fUrls             []string
	fCookie           FlagStrings
	fCookies          string
	fHeaders          FlagStrings
	fConfig           string
	fMaxDisplayedTime int
)

type FlagStrings []string

func (fs *FlagStrings) String() string         { return fmt.Sprint(*fs) }
func (fs *FlagStrings) Set(value string) error { *fs = append(*fs, value); return nil }

func flags() {
	flag.Usage = usage

	flag.IntVar(&fDuration, "d", 0, "duration")
	flag.IntVar(&fConcurrency, "c", 5, "Concurrency")
	flag.IntVar(&fNReqs, "r", 5, "requests per thread")
	fOneReq := flag.Bool("1", false, "single request (alias for -d 1 -c 1 -r 1)")
	flag.IntVar(&fRounding, "round", 250, "rounding")

	// TODO: needs to be reimplemented
	flag.StringVar(&fConfig, "config", "", "Config file to use")

	flag.BoolVar(&fReqHeaders, "req", false, "Output request headers")
	flag.BoolVar(&fRespHeaders, "resp", false, "Output response headers")
	flag.BoolVar(&fRespBody, "resp-body", false, "Output response body")
	flag.IntVar(&fMaxDisplayedTime, "m", 2000, "Max ms to display in summary")

	// HTTP request
	flag.StringVar(&fReferrer, "ref", "", "")
	flag.Var(&fCookie, "cookie", "")
	flag.StringVar(&fCookies, "cookies", "", "")
	flag.StringVar(&fUserAgent, "user-agent", fmt.Sprintf("%s/%s", AppName, Version), "")
	flag.Var(&fHeaders, "H", "")

	// behavior
	flag.Int64Var(&fTimeout, "timeout", 4000, "connection timeout")
	flag.BoolVar(&fInsecure, "insecure", false, "disable TLS validity check")
	flag.BoolVar(&fRedirects, "follow-redirects", false, "")
	fNoSmp := flag.Bool("no-smp", false, "GOMAXPROCS")
	flag.BoolVar(&fDebug, "debug", false, "debug mode")

	// logging formatting
	flag.BoolVar(&fNo200, "no-200", false, "Hide status 200 responses")
	flag.BoolVar(&fNoSummary, "no-summary", false, "Hide summary")
	fNoUtf8 := flag.Bool("no-utf8", false, "Disable UTF8 characters")
	fNoColor := flag.Bool("no-color", false, "Disable colour")
	fNoColour := flag.Bool("no-colour", false, "Disable colour") // I'm english :P
	flag.StringVar(&fFsLog, "log", "", "")

	// help
	fHelp1 := flag.Bool("h", false, "Prints this message")
	fHelp2 := flag.Bool("?", false, "Same as -h")
	fVersion1 := flag.Bool("v", false, "Prints version number")
	fVersion2 := flag.Bool("version", false, "Prints version number")

	flag.Parse()

	fUrls = flag.Args()

	// set curl-like single request (disables -d / -c / -r)
	if *fOneReq {
		fDuration = 1
		fConcurrency = 1
		fNReqs = 1
	}

	if fConfig != "" {
		fmt.Println("TODO: needs to be reimplemented")
		os.Exit(1)
	}

	if *fHelp1 || *fHelp2 {
		flag.Usage()
		os.Exit(1)
	}

	if *fVersion1 || *fVersion2 {
		fmt.Println(AppName, Version)
		fmt.Println(Copyright)
		os.Exit(1)
	}

	if fConfig == "" && len(fUrls) == 0 {
		fmt.Println("Missing parameters:")
		flag.Usage()
		os.Exit(1)
	}

	if fConcurrency == 0 && fNReqs == 0 {
		fmt.Println("Zero requests to make. Either concurrency and/or requests per thread need to be a non-zero value.")
		os.Exit(1)
	}

	if *fNoSmp {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	var passIcon, passStart, passEnd, failIcon, failStart, failEnd string

	if !*fNoColor || !*fNoColour {
		passStart = "\x1b[32m"
		passEnd = "\x1b[0m"
		failStart = "\x1b[31m"
		failEnd = "\x1b[0m"
	}

	if !*fNoUtf8 {
		passIcon = "✔"
		failIcon = "✘"
	} else {
		passIcon = " "
		failIcon = "X"
	}

	uiPass = passStart + passIcon + passEnd
	uiFail = failStart + failIcon + failEnd
}
