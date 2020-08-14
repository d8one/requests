package requests

import "net/http"

const (
	DEFAULT_TIMEOUT     = 5
	DEFAULT_CONTENTTYPE = "appliaction/json"
)

type Options struct {
	Username    string
	Password    string
	Data        []byte //Post data
	ContentType string // default appliaction/json
	Timeout     int    // default 5 sec
	Headers     map[string]string
	Client      *http.Client
}

func (opt *Options) SetDefault() {
	if opt.Timeout == 0 {
		opt.Timeout = DEFAULT_TIMEOUT
	}
	if opt.ContentType == "" {
		opt.ContentType = DEFAULT_CONTENTTYPE
	}

}

func NewOptions() Options {
	opt := Options{Headers: make(map[string]string)}
	opt.SetDefault()
	return opt

}

// SetHeaders sets headers on a http.Request
func SetHeaders(req *http.Request, opt Options) {
	for key, value := range opt.Headers {
		req.Header.Set(key, value)
	}
}
