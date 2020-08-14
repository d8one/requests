package requests

import (
	"net/http"
	"testing"
)

func TestSetHeadersWithoutHeaders(t *testing.T) {
	options := NewOptions()
	req, _ := http.NewRequest("GET", "", nil)
	SetHeaders(req, options)

}

func TestSetHeaders(t *testing.T) {
	options := NewOptions()
	options.Headers["test"] = "test"
	req, _ := http.NewRequest("GET", "", nil)
	SetHeaders(req, options)

}
