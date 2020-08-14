package requests

import (
	"bytes"
	"net/http"
	"time"
)

func Post(url string, options Options) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(options.Data))
	if err != nil {
		return &http.Response{}, err
	}

	req.Header.Set("Content-Type", options.ContentType)
	if options.Client == nil {
		client := &http.Client{Timeout: time.Second * time.Duration(options.Timeout)}
		resp, err := client.Do(req)
		if err != nil {
			return &http.Response{}, err
		}

		return resp, nil
	}

	resp, err := options.Client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}
