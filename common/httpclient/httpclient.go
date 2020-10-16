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
	err = Http("POST", url, bytes.NewBuffer(b), funcHeader, func(body *io.ReadCloser) error {
		return json.NewDecoder(*body).Decode(result)
	})
	if err != nil {
		return err
	}
	return nil
}

func PostForm(url string, data url.Values, funcHeader func(header http.Header), funcBody func(body *io.ReadCloser) error) error {
	return Http("POST", url, strings.NewReader(data.Encode()), func(header http.Header) {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		funcHeader(header)
	}, funcBody)
}

func Upload(url string, params map[string]string, fileName, path string, funcBody func(body *io.ReadCloser) error) error {
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
	}, funcBody)
	if err != nil {
		return err
	}
	return nil
}

func Http(method, url string, body io.Reader, funcHeader func(header http.Header), funcBody func(body *io.ReadCloser) error) error {
	return HttpWithContext(context.Background(), method, url, body, funcHeader, funcBody)
}
func HttpWithContext(ctx context.Context, method, url string, body io.Reader, funcHeader func(header http.Header), funcBody func(body *io.ReadCloser) error) error {
	return CallWithContext(ctx, method, url, body, func() *http.Client {
		return HttpClient
	}, funcHeader, funcBody)
}

func Call(method, url string, body io.Reader, funcClient func() *http.Client, funcHeader func(header http.Header), funcBody func(body *io.ReadCloser) error) error {
	return CallWithContext(context.Background(), method, url, body, funcClient, funcHeader, funcBody)
}

func CallWithContext(ctx context.Context, method, url string, body io.Reader, funcClient func() *http.Client, funcHeader func(header http.Header), funcBody func(body *io.ReadCloser) error) error {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return err
	}
	if funcHeader != nil {
		funcHeader(req.Header)
	}
	resp, err := funcClient().Do(req)
	defer close(resp)
	if resp.StatusCode == 200 {
		return funcBody(&resp.Body)
	}
	return errors.New(resp.Status)
}

func close(resp *http.Response) {
	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}
}
