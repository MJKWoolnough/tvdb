package tvdb

import "net/url"

type Search struct {
	Aliases    []string `json:"aliases"`
	Banner     string   `json:"banner"`
	FirstAired string   `json:"firstAired"`
	ID         uint64   `json:"id"`
	Network    string   `json:"network"`
	Overview   string   `json:"overview"`
	Name       string   `json:seriesName"`
	Status     string   `json:"status"`
}

func (c *Conn) search(key, value string) ([]Search, error) {
	var r request
	if err := c.get(&url.URL{
		Scheme:   baseURL[0:5],
		Host:     baseURL[8:],
		Path:     "/search/series",
		RawQuery: key + "=" + url.QueryEscape(value),
	}, &r); err != nil {
		if err == ErrNotFound {
			return []Search{}, nil
		}
		return nil, err
	}
	var ss []Search
	if err := r.Decode(&ss); err != nil {
		return nil, err
	}
	return ss, nil
}

func (c *Conn) Search(name string) ([]Search, error) {
	return c.search("name", name)
}

func (c *Conn) SearchIMDB(imdb string) (*Search, error) {
	ss, err := c.search("imdbId", imdb)
	if err != nil || len(ss) == 0 {
		if err == ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &ss[0], nil
}

func (c *Conn) SearchZap2It(zapit string) (*Search, error) {
	ss, err := c.search("zap2itId", zapit)
	if err != nil || len(ss) == 0 {
		if err == ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &ss[0], nil
}
