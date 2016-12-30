package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	fConcurrentUrls   bool
	fReqHeaders       bool
	fRespHeaders      bool
	fRespBody         bool
	fUserAgent        string
	fNo200            bool
	fRounding         int
	fNoSummary        bool
	fMethod           string
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

	flag.IntVar(&fDuration, "d", 0, "")
	flag.IntVar(&fConcurrency, "c", 1, "")
	flag.IntVar(&fNReqs, "r", 1, "")
	flag.BoolVar(&fConcurrentUrls, "concurrent-urls", false, "")
	flag.IntVar(&fRounding, "round", 250, "")

	// TODO: needs to be reimplemented
	flag.StringVar(&fConfig, "config", "", "")

	flag.BoolVar(&fReqHeaders, "req", false, "")
	flag.BoolVar(&fRespHeaders, "resp", false, "")
	flag.BoolVar(&fRespBody, "resp-body", false, "")
	flag.IntVar(&fMaxDisplayedTime, "m", 2000, "")

	// HTTP request
	flag.StringVar(&fMethod, "method", "GET", "")
	flag.StringVar(&fReferrer, "ref", "", "")
	flag.Var(&fCookie, "cookie", "")
	flag.StringVar(&fCookies, "cookies", "", "")
	flag.StringVar(&fUserAgent, "user-agent", fmt.Sprintf("%s/%s", AppName, Version), "")
	flag.Var(&fHeaders, "H", "")

	// behavior
	flag.Int64Var(&fTimeout, "timeout", 4000, "")
	flag.BoolVar(&fInsecure, "insecure", false, "")
	flag.BoolVar(&fRedirects, "follow-redirects", false, "")
	flag.BoolVar(&fDebug, "debug", false, "")

	// logging formatting
	flag.BoolVar(&fNo200, "no-200", false, "")
	flag.BoolVar(&fNoSummary, "no-summary", false, "")
	fNoUtf8 := flag.Bool("no-utf8", false, "")
	fNoColor := flag.Bool("no-color", false, "")
	fNoColour := flag.Bool("no-colour", false, "") // I'm english :P
	flag.StringVar(&fFsLog, "log", "", "")

	// help
	fHelp1 := flag.Bool("h", false, "")
	fHelp2 := flag.Bool("?", false, "")
	fVersion1 := flag.Bool("v", false, "")
	fVersion2 := flag.Bool("version", false, "")

	flag.Parse()

	fUrls = flag.Args()

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
