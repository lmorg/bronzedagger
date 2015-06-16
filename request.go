package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func ParseCookie(job *Job, s string) {
	split := strings.Split(s, "=")
	if len(split) >= 2 {
		job.Cookies[split[0]] = strings.Join(split[1:], "=")
	} else {
		fmt.Println("Invalid cookie format")
		os.Exit(1)
	}
}

func dialTimeout(job *Job) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		// http://stackoverflow.com/questions/16895294/how-to-set-timeout-for-http-get-requests-in-golang#16930649
		conn, err := net.DialTimeout(netw, addr, job.Timeout) //connect
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(job.Timeout)) //reply
		return conn, nil
	}
}

func httpClient(job *Job) (client *http.Client, request *http.Request) {
	client = new(http.Client)

	u, err := url.Parse(job.URL)
	isErr(err)

	tr := http.Transport{
		Dial: dialTimeout(job),
	}

	tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: job.Insecure}

	client.Transport = &tr
	client.Timeout = job.Timeout

	//request, err = http.NewRequest("GET", u.Scheme+"://"+ip+u.RequestURI(), nil)
	request, err = http.NewRequest(job.Method, job.URL, nil) // TODO: this will eventually support IPs with hostnames using the above code
	isErr(err)

	request.Header.Set("User-Agent", job.UserAgent)
	request.Header.Set("Referer", job.Referrer)
	debugLog("Host:", u.Host)
	request.Host = u.Host
	job.AddCookies(request)

	return client, request
}

func httpRequest(job *Job, client *http.Client, request *http.Request) {
	var (
		resp *http.Response
		body []byte
		err  error
	)

	t_start := time.Now()

	if job.FollowRedirects {
		resp, err = client.Do(request)
	} else {
		resp, err = client.Transport.RoundTrip(request)
	}

	if err == nil {
		body, err = ioutil.ReadAll(resp.Body)
	}
	t_duration := time.Now().Sub(t_start)
	duration := t_duration.Nanoseconds() / 1000000

	if err != nil {
		ret := func(status int) {
			update_results <- Response{Status: status}
		}

		if err.Error() == "Fetch "+job.URL+": EOF" {
			log.Println(CROSS, `001 empty response`, job.URL, len(body), duration)
			ret(1)
		} else if strings.Contains(err.Error(), "too many open files") {
			log.Println(CROSS, `002 too many open files`, job.URL, len(body), duration)
			ret(2)
		} else if strings.Contains(err.Error(), "connection reset by peer") {
			log.Println(CROSS, `003 connection reset by peer`, job.URL, len(body), duration)
			ret(3)
		} else if strings.Contains(err.Error(), "connection refused") {
			log.Println(CROSS, `004 connection refused`, job.URL, len(body), duration)
			ret(4)
		} else if strings.Contains(err.Error(), "EOF") {
			log.Println(CROSS, `005 EOF`, job.URL, len(body), duration)
			ret(5)
		} else if strings.Contains(err.Error(), "i/o timeout") {
			log.Println(CROSS, `006 connection timed out`, job.URL, len(body), duration)
			ret(6)
		} else {
			log.Println(CROSS, `000 `+err.Error(), job.URL, len(body), duration)
			ret(0)
		}
		return
	}

	ret := Response{Status: resp.StatusCode}

	if resp.StatusCode == 200 {
		ret.Duration = Lower(int(duration))
		if !f_no_200 {
			log.Println(TICK, resp.StatusCode, job.URL, len(body), duration)
		}
	} else {
		log.Println(CROSS, resp.StatusCode, job.URL, len(body), duration)
	}
	update_results <- ret
	resp.Body.Close()

	if req_headers {
		for key, val := range request.Header {
			fmt.Printf("%20s: %s\n", key, val)
		}
		fmt.Println()
	}

	if resp_headers {
		for key, val := range resp.Header {
			fmt.Printf("%20s: %s\n", key, val)
		}
		fmt.Println()
	}

	if resp_body {
		fmt.Println(string(body))
	}

	if f_firesword_log != "" {
		firesword_log <- fmt.Sprintf(`127.0.0.1 - - [%s] "GET %s HTTP/1.1" %d %d "%s" "%s" %d`,
			t_start.Format("02/Jan/2006:15:04:05 -0700"),
			job.URL, //TODO: no domain name please
			resp.StatusCode,
			len(body),
			"-", //TODO: isBlank(job.Referrer)
			job.UserAgent,
			duration,
		)
	}
}
