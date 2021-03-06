package main

import "fmt"

func usage() {
	fmt.Print(`Usage: bronzedagger [options] http://example1.com http://example2.com
                    [options] --config example.com.json

Request weight:
---------------
  -d int         Duration to run test (default is indefinitely)
  -r int         Number of requests per routines - executed in sequence (default is 1)
  -c int         Concurrency; number of routines running in parallel (default is 1)
  --concurrent-urls If multiple URLs supplied then test then concurrently rather than sequentially

Request packet:
---------------
  --method         Method (default is GET)
  --ref str        Referrer
  --cookie str     Requester cookie (flag can be issued multiple times)
  --cookies str    Requester cookies (semi-colon delimited, flag only used one)
  --user-agent str Change user agent
  -H str           Header
  --body           Request body
  --stdin          Set HTTP request body to be populated from stdin (overrides --body)

Runtime behavior:
-----------------
  --timeout int  Connection timeout in seconds (default is 4)
  --insecure     Disable TLS certificate checks
  --follow-redirects  Follow redirects

Results:
--------
  -m int         Max ms to display in summary
  --round int    In summary group results by duration, rounded down to the nearest int
  --req          Display request headers and body
  --resp-head    Display response headers
  --resp-body    Display response body

  --log str      Write Apache combined log file for parsing with Firesword / Plasmasword

  --no-summary   Don't show summary (more useful for shell scripting)
  --no-200       Hide HTTP status 200 from summary
  --no-utf8      Disable UTF-8 characters]
  --no-colour    Disables terminal colour escape sequences (--no-color also works)

Help:
-----
  -h | -?        Prints this usage guide
  -v | --version Prints version number
`)
}
