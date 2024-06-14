// Package tvdb is a simple interface to the TVDB database of TV shows.
package tvdb // import "vimagination.zapto.org/tvdb"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

// Auth represents the information required to get a validated authentication
// token.
type Auth struct {
	APIKey   string `json:"apikey"`
	Username string `json:"username,omitempty"`
	UserKey  string `json:"userkey,omitempty"`
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

// Conn represents a connection to the TVDB database.
type Conn struct {
	headerMutex sync.RWMutex
	headers     http.Header
}

// Token creates a TVDB database connection using a pre-authorised token.
func Token(t string) *Conn {
	return &Conn{
		headers: http.Header{
			"Authorization": []string{
				"Bearer " + t,
			},
			"Content-Type": contentType,
		},
	}
}

// Login creates a TVDB database connection using login credentials.
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

// Token returns the current authentication token.
func (c *Conn) Token() string {
	a := c.headers.Get("Authorization")
	if len(a) < 7 {
		return ""
	}

	return a[7:]
}

// Refresh retrieves a new authentication token without having to use the login
// credentials. Each token only lasts 24 hours and refresh can only be used in
// that time-frame.
func (c *Conn) Refresh() error {
	var ar authResponse

	if err := c.get(refreshURL, &ar); err != nil {
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

// SetLanguage sets the language header used by some queries to return
// information in the requested language.
func (c *Conn) SetLanguage(code string) {
	c.headers.Set("Accept-Language", code)
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

		r.Body = io.NopCloser(&buf)
		r.ContentLength = int64(buf.Len())
	}

	c.headerMutex.RLock()
	resp, err := http.DefaultClient.Do(&r)
	c.headerMutex.RUnlock()

	if err != nil {
		return fmt.Errorf("error making connection: %w", err)
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
			return fmt.Errorf("error decoding response: %w", err)
		}

		if err = resp.Body.Close(); err != nil {
			return fmt.Errorf("error closing response body: %w", err)
		}
	}

	for k := range headers {
		headers[k] = resp.Header[k]
	}

	return nil
}

// Errors.
var (
	ErrInvalidAuth  = errors.New("invalid credentials")
	ErrUnknownError = errors.New("unknown error")
	ErrNotFound     = errors.New("not found")
)
