package main

import (
	"log"
	"net/http"
)

// interfaceを使うためにやってるお作法
type LogTransport struct {
    Transport http.RoundTripper
}

func (lt *LogTransport) transport() http.RoundTripper {
    if lt.Transport == nil {
        log.Print("its default")
        return http.DefaultTransport
    }
    log.Print("its not default")
    log.Print(lt.Transport)
    return lt.Transport
}

func (lt *LogTransport) RoundTrip(req *http.Request) (*http.Response, error) {
    res, err := lt.transport().RoundTrip(req)

    log.Printf("%d\t%s\t%s\n", res.StatusCode, req.Method, req.URL.String())

    return res, err
}

func main(){
	client := &http.Client{
        Transport: &LogTransport{},
    }

	res, err := client.Get("https://google.com")
	if err != nil {
		// handle error 今回は実装しない
	}
	defer res.Body.Close()
}