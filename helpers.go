package requests

import (
	"io/ioutil"
	"net/http"
)

// ReadBody takes a http.Response and returns a []byte resp
// of the body of the response
func ReadBody(r *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte(""), err
	}
	return body, nil

}

//CloseBody closes the Body of a *http.Response
func CloseBody(r *http.Response) {
	r.Body.Close()

}
