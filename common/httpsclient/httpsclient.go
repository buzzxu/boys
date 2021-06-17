package httpsclient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/buzzxu/boys/common/httpclient"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var HttpsClient = &http.Client{
	Timeout: 3 * time.Second,
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}

func JSON(url string, data interface{}, result interface{}, funcHeader func(header http.Header)) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = Https("POST", url, bytes.NewBuffer(b), funcHeader, func(response *http.Response) error {
		return json.NewDecoder(response.Body).Decode(result)
	})
	if err != nil {
		return err
	}
	return nil
}

func PostForm(url string, data url.Values, funcHeader func(header http.Header), funcResponse func(response *http.Response) error) error {
	return Https("POST", url, strings.NewReader(data.Encode()), func(header http.Header) {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		funcHeader(header)
	}, funcResponse)
}

func Https(method, url string, body io.Reader, funcHeader func(header http.Header), funcResponse func(response *http.Response) error) error {
	return httpclient.Call(method, url, body, func() *http.Client {
		return HttpsClient
	}, funcHeader, funcResponse)
}
