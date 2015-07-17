package main

import (
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Job struct {
	URL             string
	IPOverride      string //TODO: impliment this
	Cookies         map[string]string
	Headers         map[string]string
	Body            string
	Referrer        string
	Concurrency     int
	Duration        int
	Threads         int
	Timeout         time.Duration
	Insecure        bool
	FollowRedirects bool
	Method          string
	UserAgent       string
}

func NewJob() (job *Job) {
	job = new(Job)
	job.Cookies = make(map[string]string)
	job.Headers = make(map[string]string)

	// some defaults
	job.Concurrency = f_concurrency
	job.Threads = f_nreqs
	job.Duration = f_duration

	job.FollowRedirects = f_redirects
	job.Insecure = f_insecure
	job.Timeout = time.Duration(f_timeout * int64(time.Millisecond))

	job.UserAgent = f_user_agent
	job.Referrer = f_referrer
	job.Method = "GET" // no other method currently supported

	for i, _ := range f_cookies {
		ParseCookie(job, f_cookies[i])
	}

	return
}

func (job *Job) Fork(url string) (fork Job) {
	fork = *job
	fork.URL = url

	return
}

func (job *Job) AddCookies(request *http.Request) {
	debugLog("Adding cookies")

	for name, value := range job.Cookies {
		request.AddCookie(&http.Cookie{
			Name:  name,
			Value: value,
		})
	}
}

func (job *Job) Start(wait *sync.WaitGroup) {
	var (
		wg      sync.WaitGroup
		client  *http.Client
		request *http.Request
	)

	u, err := url.Parse(job.URL)
	isErr(err)
	addr, err := net.LookupHost(u.Host)
	//isErr(err)
	addr = []string{"127.0.0.1"} // TODO: delete this crap

	debugLog("addr: ", addr)

	for i := 0; job.Duration == 0 || i < job.Duration; i++ {
		go func() {
			for j := 0; j < job.Concurrency; j++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for n := 0; n < f_nreqs; n++ {
						client, request = httpClient(job)
						httpRequest(job, client, request)
					}
				}()
			}
		}()

		//if r.Duration == 0 || i < r.Duration {
		time.Sleep(1000 * time.Millisecond)
		//}
	}
	wg.Wait()
	if wait != nil {
		wait.Done()
	}
}
