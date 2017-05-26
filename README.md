# bronzedagger
Simple load testing and page testing.

This is still very much beta and under active development. But most of the basic features I want are already implemented.

Logging on bronzedagger can also be imported into [Firesword](https://github.com/lmorg/firesword) or [Plasmasword](https://github.com/lmorg/plasmasword) for analytics.

### How to compile:

If you don't already have Go installed, then that can be downloaded from https://golang.org/ or installed via your OS's package manager (if available). Make sure you follow the guide to set your GOPATH environmental variables, eg `export GOPATH=~/go; mkdir -p $GOPATH/src $GOPATH/bin`

Then download this repo and install

    go get github.com/lmorg/bronzedagger
    go install github.com/lmorg/bronzedagger

### Command line flags:

    Usage: bronzedagger [options] http://example1.com http://example2.com

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
      --cookies str    Requester cookies
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
      --no-colour    Disables terminal colour escape sequences

    Help:
    -----
      -h | -?        Prints this usage guide
      -v | --version Prints version number
