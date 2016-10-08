package tvdb

import "net/url"

// Search is a representation of the data returned from a tv show search
type Search struct {
	Aliases    []string `json:"aliases"`
	Banner     string   `json:"banner"`
	FirstAired string   `json:"firstAired"`
	ID         uint64   `json:"id"`
	Network    string   `json:"network"`
	Overview   string   `json:"overview"`
	Name       string   `json:"seriesName"`
	Status     string   `json:"status"`
}

func (c *Conn) search(key, value string) ([]Search, error) {
	var r struct {
		Data  []Search      `json:"data"`
		Error requestErrors `json:"error"`
	}
	if err := c.get(makeURL("/search/series", key+"="+url.QueryEscape(value)), &r); err != nil {
		if err == ErrNotFound {
			return []Search{}, nil
		}
		return nil, err
	}
	return r.Data, nil
}

// Search searches the TVDB database for shows with the given name
func (c *Conn) Search(name string) ([]Search, error) {
	return c.search("name", name)
}

// SearchIMDB searches the TVDB database for the show corrensponding to the
// given IMDB ID
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

// SearchZap2It searches the TVDB database for the show corrensponding to the
// given Zap2It ID
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
