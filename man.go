package main

import "fmt"

func Usage() {
	fmt.Print(`Usage: bronzedagger [options] http://example1.com http://example2.com
                    [options] --config example.com.json

Request weight:
---------------
  -d int         Duration to run test (default is indefinitely)
  -r int         Number of requests per routines (default is 5)
  -c int         Concurrency; number of routines running in parallel (default is 5)

Request packet:
---------------
  --ref          Referrer
  --cookie       Requester cookies
  --user-agent   Change user agent

Runtime behavior:
-----------------
  --timeout      Connection timeout in seconds (default is 4)
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
  --no-color     Disables terminal color escape sequences

Help:
-----
  -h | -?        Prints this usage guide
  -v             Prints version number
`)
}
