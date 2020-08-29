package files

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

var tr = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}

func URLReader(url string) (r *bytes.Reader, err error) {
	buff, err := URL(url)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(*buff), nil
}

func URL(url string) (*[]byte, error) {
	var err error
	if url[0:4] == "file" {
		buff, err := ioutil.ReadFile(url[6:])
		if err != nil {
			return nil, err
		}
		return &buff, nil
	} else {
		var resp *http.Response
		if url[0:5] == "https" {
			c := &http.Client{
				Transport: tr,
			}
			resp, err = c.Get(url)
		} else {
			resp, err = http.Get(url)
		}
		defer resp.Body.Close()
		if err != nil {
			return nil, err
		}
		buff, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return &buff, nil
	}
}
