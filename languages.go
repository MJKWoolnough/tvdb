package tvdb

import (
	"fmt"
	"net/url"
)

type Language struct {
	ID           uint64 `json:"id"`
	Abbreviation string `json:"abbreviation"`
	Name         string `json:"name"`
	EnglishName  string `json:"englishName"`
}

var languageURL = &url.URL{
	Scheme: baseURL[0:5],
	Host:   baseURL[8:],
	Path:   "/languages",
}

func (c *Conn) Languages() ([]Language, error) {
	var r request
	if err := c.get(languageURL, &r); err != nil {
		return nil, err
	}
	var l []Language
	if err := r.Decode(&l); err != nil {
		return nil, err
	}
	return l, nil
}

func (c *Conn) Language(id uint64) (*Language, error) {
	var r request
	if err := c.get(&url.URL{
		Scheme: baseURL[0:5],
		Host:   baseURL[8:],
		Path:   fmt.Sprintf("/languages/%d", id),
	}, &r); err != nil {
		return nil, err
	}
	var l Language
	if err := r.Decode(&l); err != nil {
		return nil, err
	}
	return &l, nil
}
