package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var HttpClient = &http.Client{
	Timeout: 3 * time.Second,
}

func JSON(url string, data interface{}, result interface{}, funcHeader func(header http.Header)) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body, err := Http("POST", url, bytes.NewBuffer(b), funcHeader)
	if err != nil {
		return err
	}
	json.Unmarshal(*body, result)
	return nil
}

func PostForm(url string, data url.Values, funcHeader func(header http.Header)) (*[]byte, error) {
	return Http("POST", url, strings.NewReader(data.Encode()), func(header http.Header) {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		funcHeader(header)
	})
}

func Upload(url string, params map[string]string, fileName, path string) (*[]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fileName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return Http("POST", url, body, func(header http.Header) {
		header.Set("Content-Type", writer.FormDataContentType())
	})
}

func Http(method, url string, body io.Reader, funcHeader func(header http.Header)) (*[]byte, error) {
	return HttpWithContext(context.Background(), method, url, body, funcHeader)
}
func HttpWithContext(ctx context.Context, method, url string, body io.Reader, funcHeader func(header http.Header)) (*[]byte, error) {
	return CallWithContext(ctx, method, url, body, func() *http.Client {
		return HttpClient
	}, funcHeader)
}

func Call(method, url string, body io.Reader, funcClient func() *http.Client, funcHeader func(header http.Header)) (*[]byte, error) {
	return CallWithContext(context.Background(), method, url, body, funcClient, funcHeader)
}

func CallWithContext(ctx context.Context, method, url string, body io.Reader, funcClient func() *http.Client, funcHeader func(header http.Header)) (*[]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	if funcHeader != nil {
		funcHeader(req.Header)
	}
	resp, err := funcClient().Do(req)
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return &data, nil
	}
	return nil, errors.New(resp.Status)
}
