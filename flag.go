package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

var (
	f_referrer      string
	f_redirects     bool
	f_duration      int
	f_concurrency   int
	f_insecure      bool
	f_timeout       int64
	f_nreqs         int
	f_firesword_log string
	f_debug         bool
	f_urls          []string
	f_cookie        string
	f_user_agent    string
	f_config        string
)

func Flags() {
	flag.Usage = Usage

	flag.IntVar(&f_duration, "d", 0, "duration")
	flag.IntVar(&f_concurrency, "c", 5, "Concurrency")
	flag.IntVar(&f_nreqs, "r", 5, "requests per thread")
	flag.IntVar(&f_rounding, "round", 250, "rounding")

	flag.StringVar(&f_config, "config", "", "Config file to use")

	flag.BoolVar(&req_headers, "req", false, "Output request headers")
	flag.BoolVar(&resp_headers, "resp", false, "Output response headers")
	flag.BoolVar(&resp_body, "resp-body", false, "Output response body")
	flag.IntVar(&display_max, "m", 2000, "Max ms to display in summary")

	// HTTP request
	flag.StringVar(&f_referrer, "ref", "", "")
	flag.StringVar(&f_cookie, "cookie", "", "")
	flag.StringVar(&f_user_agent, "user-agent", USER_AGENT, "")

	// behavior
	flag.Int64Var(&f_timeout, "timeout", 4000, "connection timeout")
	flag.BoolVar(&f_insecure, "insecure", false, "disable TLS validity check")
	flag.BoolVar(&f_redirects, "follow-redirects", false, "")
	f_gomaxprocs := flag.Bool("no-smp", false, "GOMAXPROCS")
	flag.BoolVar(&f_debug, "debug", false, "debug mode")

	// logging formating
	flag.BoolVar(&f_no_200, "no-200", false, "Hide status 200 responses")
	f_no_utf8 := flag.Bool("no-utf8", false, "Disable UTF8 characters")
	f_no_colour := flag.Bool("no-color", false, "Disable color")
	flag.StringVar(&f_firesword_log, "log", "", "")

	// help
	f_help1 := flag.Bool("h", false, "Prints this message")
	f_help2 := flag.Bool("?", false, "Same as -h")
	f_version1 := flag.Bool("v", false, "Prints version number")
	f_version2 := flag.Bool("version", false, "Prints version number")

	flag.Parse()

	f_urls = flag.Args()

	if *f_help1 || *f_help2 {
		flag.Usage()
		os.Exit(1)
	}

	if *f_version1 || *f_version2 {
		fmt.Println(APP_NAME, VERSION, "\n"+COPYRIGHT)
		os.Exit(1)
	}

	if f_config == "" && len(f_urls) == 0 {
		fmt.Println("Missing parameters:")
		flag.Usage()
		os.Exit(1)
	}

	if f_concurrency == 0 && f_nreqs == 0 {
		fmt.Println("Zero requests to make. Either concurrency and/or requests per thread need to be a non-zero value.")
		os.Exit(1)
	}

	if *f_gomaxprocs {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	var pass_ico, pass_start, pass_end, fail_ico, fail_start, fail_end string

	if !*f_no_colour {
		pass_start = "\x1b[32m"
		pass_end = "\x1b[0m"
		fail_start = "\x1b[31m"
		fail_end = "\x1b[0m"
	}

	if !*f_no_utf8 {
		pass_ico = "✔"
		fail_ico = "✘"
	} else {
		pass_ico = " "
		fail_ico = "X"
	}

	TICK = pass_start + pass_ico + pass_end
	CROSS = fail_start + fail_ico + fail_end
}
