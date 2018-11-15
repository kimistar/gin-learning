package httplib

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HttpRequest struct {
	header map[string]string
	req    *http.Request
}

// 获取http client
func httpClient() *http.Client {
	trans := &http.Transport{
		// 不验证证书
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: trans,
	}
	return client
}

func Get(url string) (*HttpRequest, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return &HttpRequest{
		req:    req,
		header: map[string]string{},
	}, nil
}

func Post(url string) (*HttpRequest, error) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	return &HttpRequest{
		req:    req,
		header: map[string]string{},
	}, nil
}

// 向请求中添加header
func (r *HttpRequest) Header(key, value string) *HttpRequest {
	r.header[key] = value
	return r
}

// string []byte写入请求body
func (r *HttpRequest) Body(data interface{}) *HttpRequest {
	switch t := data.(type) {
	case string:
		bf := bytes.NewBufferString(t)
		r.req.Body = ioutil.NopCloser(bf)
		r.req.ContentLength = int64(len(t))
	case []byte:
		bf := bytes.NewBuffer(t)
		r.req.Body = ioutil.NopCloser(bf)
		r.req.ContentLength = int64(len(t))
	}
	return r
}

// form写入请求body
func (r *HttpRequest) FormBody(values url.Values) *HttpRequest {
	if r.req.Body == nil && values != nil {
		r.req.Body = ioutil.NopCloser(strings.NewReader(values.Encode()))
		r.req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	return r
}

// json写入请求body
func (r *HttpRequest) JsonBody(v interface{}) (*HttpRequest, error) {
	if r.req.Body == nil && v != nil {
		byts, err := json.Marshal(v)
		if err != nil {
			return r, err
		}
		r.req.Body = ioutil.NopCloser(bytes.NewReader(byts))
		r.req.ContentLength = int64(len(byts))
		r.req.Header.Set("Content-Type", "application/json")
	}
	return r, nil
}

// xml写入请求body
func (r *HttpRequest) XmlBody(v interface{}) (*HttpRequest, error) {
	if r.req.Body == nil && v != nil {
		byts, err := xml.Marshal(v)
		if err != nil {
			return r, err
		}
		r.req.Body = ioutil.NopCloser(bytes.NewReader(byts))
		r.req.ContentLength = int64(len(byts))
		r.req.Header.Set("Content-Type", "application/xml")
	}
	return r, nil
}

// 获取响应对象
func (r *HttpRequest) Response() (*http.Response, error) {
	for k, v := range r.header {
		r.req.Header.Set(k, v)
	}

	client := httpClient()

	resp, err := client.Do(r.req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// 获取响应体（string）
func (r *HttpRequest) String() (string, error) {
	resp, err := r.Response()
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (r *HttpRequest) ParseJson(v interface{}) error {
	body, err := r.String()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(body), v)
}

func (r *HttpRequest) ParseXml(v interface{}) error {
	body, err := r.String()
	if err != nil {
		return err
	}
	return xml.Unmarshal([]byte(body), v)
}
