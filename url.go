package tvdb

import "net/url"

func makeURL(path, query string) *url.URL {
	return &url.URL{
		Scheme:   "https",
		Host:     "api.thetvdb.com",
		Path:     path,
		RawQuery: query,
	}
}
