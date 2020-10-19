package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var HttpClient = &http.Client{
	Timeout: 30 * time.Second,
}

func JSON(url string, data interface{}, result interface{}, funcHeader func(header http.Header)) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = Http("POST", url, bytes.NewBuffer(b), funcHeader, func(response *http.Response) error {
		return json.NewDecoder(response.Body).Decode(result)
	})
	if err != nil {
		return err
	}
	return nil
}

func PostForm(url string, data url.Values, funcHeader func(header http.Header), funcResponse func(response *http.Response) error) error {
	return Http("POST", url, strings.NewReader(data.Encode()), func(header http.Header) {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		funcHeader(header)
	}, funcResponse)
}

func Upload(url string, params map[string]string, fileName, path string, funcResponse func(response *http.Response) error) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fileName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return err
	}
	err = Http("POST", url, body, func(header http.Header) {
		header.Set("Content-Type", writer.FormDataContentType())
	}, funcResponse)
	if err != nil {
		return err
	}
	return nil
}

func Http(method, url string, body io.Reader, funcHeader func(header http.Header), funcResponse func(response *http.Response) error) error {
	return HttpWithContext(context.Background(), method, url, body, funcHeader, funcResponse)
}
func HttpWithContext(ctx context.Context, method, url string, body io.Reader, funcHeader func(header http.Header), funcResponse func(response *http.Response) error) error {
	return CallWithContext(ctx, method, url, body, func() *http.Client {
		return HttpClient
	}, funcHeader, funcResponse)
}

func Call(method, url string, body io.Reader, funcClient func() *http.Client, funcHeader func(header http.Header), funcResponse func(response *http.Response) error) error {
	return CallWithContext(context.Background(), method, url, body, funcClient, funcHeader, funcResponse)
}

func CallWithContext(ctx context.Context, method, url string, body io.Reader, funcClient func() *http.Client, funcHeader func(header http.Header), funcResponse func(response *http.Response) error) error {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return err
	}
	if funcHeader != nil {
		funcHeader(req.Header)
	}
	resp, err := funcClient().Do(req)
	if err != nil {
		return err
	}
	defer close(resp)
	if resp.StatusCode == 200 {
		return funcResponse(resp)
	}
	return errors.New(resp.Status)
}

func close(resp *http.Response) {
	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}
}
