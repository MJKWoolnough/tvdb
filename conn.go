package tvdb

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"sync"

	"github.com/MJKWoolnough/errors"
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

const baseURL = "https://api.thetvdb.com"

var (
	loginURL = &url.URL{
		Scheme: baseURL[0:5],
		Host:   baseURL[8:],
		Path:   "/login",
	}
	refreshURL = &url.URL{
		Scheme: baseURL[0:5],
		Host:   baseURL[8:],
		Path:   "/refresh_token",
	}
)

var contentType = []string{
	"application/json",
}

var loginHeaders = map[string][]string{
	"Content-Type": contentType,
}

type Conn struct {
	headerMutex sync.RWMutex
	headers     http.Header
}

func Login(a Auth) (*Conn, error) {
	pr, pw := io.Pipe()
	go func() {
		json.NewEncoder(pw).Encode(a)
		pw.Close()
	}()

	resp, err := http.DefaultClient.Do(&http.Request{
		URL:    loginURL,
		Header: loginHeaders,
		Method: http.MethodPost,
		Body:   pr,
	})

	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusUnauthorized:
		return nil, ErrInvalidAuth
	default:
		return nil, ErrUnknownError
	}

	var ar authResponse

	err = json.NewDecoder(resp.Body).Decode(&ar)

	resp.Body.Close()

	if err != nil {
		return nil, err
	} else if ar.Error != "" {
		return nil, errors.Error(ar.Error)
	} else if ar.Token == "" {
		return nil, ErrUnknownError
	}

	return &Conn{
		headers: map[string][]string{
			"Authorization": []string{
				"Bearer " + ar.Token,
			},
			"Content-Type": contentType,
		},
	}, nil
}

func (c *Conn) do(u *url.URL, body io.ReadCloser) (*http.Response, error) {
	c.headerMutex.RLock()
	resp, err := http.DefaultClient.Do(&http.Request{
		URL:    u,
		Header: c.headers,
		Method: http.MethodPost,
		Body:   body,
	})
	c.headerMutex.RUnlock()
	return resp, err
}

func (c *Conn) Refresh() error {
	resp, err := c.do(refreshURL, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return ErrUnknownError
	}

	var ar authResponse

	err = json.NewDecoder(resp.Body).Decode(&ar)

	resp.Body.Close()

	if err != nil {
		return err
	} else if ar.Error != "" {
		return errors.Error(ar.Error)
	} else if ar.Token == "" {
		return ErrUnknownError
	}

	c.headerMutex.Lock()
	c.headers["Authorization"][0] = "Bearer " + ar.Token
	c.headerMutex.Unlock()

	return nil
}

const (
	ErrInvalidAuth  errors.Error = "Invalid Credentials"
	ErrUnknownError errors.Error = "Unknown Error"
)
