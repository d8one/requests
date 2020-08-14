package requests

import (
	"net/http"
	"time"
)

func Get(url string, options Options) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &http.Response{}, err
	}
	if options.Timeout == 0 {
		// we should always have a timeout >0
		options.Timeout = 15
	}

	client := &http.Client{Timeout: time.Second * time.Duration(options.Timeout)}
	resp, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}
