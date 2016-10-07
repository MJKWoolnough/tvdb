package tvdb

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

type Auth struct {
	APIKey   string `json:"apikey"`
	Username string `json:"username"`
	UserKey  string `json:"userkey"`
}

type authResponse struct {
	Token string `json:"token"`
	Error string `json:"Error"`
}

var (
	loginURL   = makeURL("/login", "")
	refreshURL = makeURL("/refresh_token", "")
)

var contentType = []string{
	"application/json",
}

var loginHeaders = http.Header{
	"Content-Type": contentType,
}

type Conn struct {
	headerMutex sync.RWMutex
	headers     http.Header
}

func Login(a Auth) (*Conn, error) {

	c := &Conn{
		headers: loginHeaders,
	}

	var ar authResponse

	if err := c.post(loginURL, a, &ar); err != nil {
		return nil, err
	}

	if ar.Error != "" {
		return nil, errors.New(ar.Error)
	} else if ar.Token == "" {
		return nil, ErrUnknownError
	}

	c.headers = http.Header{
		"Authorization": []string{
			"Bearer " + ar.Token,
		},
		"Content-Type": contentType,
	}
	return c, nil
}

func (c *Conn) Refresh() error {
	var ar authResponse
	err := c.get(refreshURL, &ar)
	if err != nil {
		return err
	}

	if ar.Error != "" {
		return errors.New(ar.Error)
	} else if ar.Token == "" {
		return ErrUnknownError
	}

	c.headerMutex.Lock()
	c.headers["Authorization"][0] = "Bearer " + ar.Token
	c.headerMutex.Unlock()

	return nil
}

func (c *Conn) get(u *url.URL, ret interface{}) error {
	return c.do(http.MethodGet, u, nil, ret, nil)
}

func (c *Conn) post(u *url.URL, data interface{}, ret interface{}) error {
	return c.do(http.MethodPost, u, data, ret, nil)
}

func (c *Conn) put(u *url.URL, ret interface{}) error {
	return c.do(http.MethodPut, u, nil, ret, nil)
}

func (c *Conn) delete(u *url.URL, ret interface{}) error {
	return c.do(http.MethodDelete, u, nil, ret, nil)
}

func (c *Conn) head(u *url.URL, headers http.Header) error {
	return c.do(http.MethodHead, u, nil, nil, headers)
}

func (c *Conn) do(method string, u *url.URL, data interface{}, ret interface{}, headers http.Header) error {
	r := http.Request{
		URL:    u,
		Header: c.headers,
		Method: method,
	}

	if method == http.MethodPost && data != nil {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(data); err != nil {
			return err
		}
		r.Body = ioutil.NopCloser(&buf)
		r.ContentLength = int64(buf.Len())
	}
	c.headerMutex.RLock()
	resp, err := http.DefaultClient.Do(&r)
	c.headerMutex.RUnlock()
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusUnauthorized:
		return ErrInvalidAuth
	default:
		return ErrUnknownError
	}

	if ret != nil {
		if err = json.NewDecoder(resp.Body).Decode(ret); err != nil {
			return err
		}
		resp.Body.Close()
	}

	for k := range headers {
		headers[k] = resp.Header[k]
	}

	return nil
}

var (
	ErrInvalidAuth  = errors.New("Invalid Credentials")
	ErrUnknownError = errors.New("Unknown Error")
	ErrNotFound     = errors.New("Not Found")
)
