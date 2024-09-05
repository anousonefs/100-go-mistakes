package main

import (
	"net"
	"net/http"
	"time"
)

/*
By default, an http.Client in Go has no timeout, which means that it can potentially wait indefinitely for a response. This could lead to hanging requests if the server doesn't respond or takes too long. To avoid this, it's a good idea to set a timeout for your http.Client operations.

The four main timeouts are the following:
	1. net.Dialer.Timeout—Specifies the maximum amount of time a dial will wait for a connection to complete.
	2. http.Transport.TLSHandshakeTimeout—Specifies the maximum amount of time to wait for the TLS handshake.
	3. http.Transport.ResponseHeaderTimeout—Specifies the amount of time to wait for a server’s response headers.
	4. http.Client.Timeout—Specifies the time limit for a request. It includes all the steps, from step 1 (dial) to step 5 (read the response body).
*/

// default timeout
func getDetail() {
	client := &http.Client{}
	resp, err := client.Get("https://golang.org/")
	if err != nil {
		panic(err)
	}
	/* resp, err := http.Get("https://golang.org/") */
	_ = resp
}

/*
	- http.Transport.MaxIdleConns This value is set to 100 by default
	- http.Transport.MaxIdleConnsPerHost limit per host default is set to 2
	if we trigger 100 requests again, we will have to reopen at least 98 connections
*/
// set timeout
func getDetail2() {
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   time.Second,
			ResponseHeaderTimeout: time.Second,
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   4,
		}}

	for i := 0; i <= 200; i++ {
		resp, err := client.Get("https://golang.org/") // only 4 idle connection
		if err != nil {
			panic(err)
		}

		resp, err = client.Get("https://google.com/") // only 4 idle connection
		if err != nil {
			panic(err)
		}
		_ = resp
	}
}
