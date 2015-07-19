package main

import "fmt"

func Usage() {
	fmt.Print(`Usage: bronzedagger [options] http://example1.com http://example2.com
                    [options] --config example.com.json

Request weight:
---------------
  -d int         Duration to run test (default is indefinitely)
  -r int         Number of requests per routines - executed in sequence (default is 5)
  -c int         Concurrency; number of routines running in parallel (default is 5)
  -1             Single request (alias for -d 1 -r 1 -c 1)

Request packet:
---------------
  --ref str        Referrer
  --cookies str    Requester cookies
  --user-agent str Change user agent
  -H str           Header

Runtime behavior:
-----------------
  --timeout int  Connection timeout in seconds (default is 4)
  --insecure     Disable TLS certificate checks
  --no-smp       Disable multi-processor support (SMP enabled by default)
  --follow-redirects  Follow redirects

Results:
--------
  -m int         Max ms to display in summary
  --round int    Group results by duration, rounded down to the nearest int
  --req          Display request headers
  --resp         Display response headers
  --resp-body    Display response body

  --log str      Write Apache combined log file for parsing with Firesword

  --no-200       Hide HTTP status 200 from summary
  --no-utf8      Disable UTF-8 characters]
  --no-colour    Disables terminal colour escape sequences

Help:
-----
  -h | -?        Prints this usage guide
  -v             Prints version number
`)
}