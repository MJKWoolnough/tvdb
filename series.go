package tvdb

import (
	"fmt"
	"net/url"
	"strconv"
)

type Series struct {
	ID              uint64   `json:"id"`
	Name            string   `json:"seriesName"`
	Aliases         []string `json:"aliases"`
	Banner          string   `json:"banner"`
	SeriesID        uint64   `json:"seriesID"`
	Status          string   `json:"status"`
	FirstAired      string   `json:"firstAired"`
	Network         string   `json:"network"`
	NetworkID       string   `json:"networkId"`
	Runtime         string   `json:"runtime"`
	Genre           []string `json:"genre"`
	Overview        string   `json:"overview"`
	LastUpdated     uint64   `json:"lastUpdated"`
	AirsDayOfWeek   string   `json:"airsDayOfWeek"`
	AirsTime        string   `json:"airsTime"`
	Rating          string   `json:"rating"`
	IMDBID          string   `json:"imdbId"`
	Zap2ItID        string   `json:"zap2ItId"`
	Added           string   `json:"added"`
	siteRating      float32  `json:"siteRating"`
	siteRatingCount uint64   `json:"siteRatingCount"`
}

func (c *Conn) Series(id uint64) (*Series, error) {
	var r struct {
		Data  *Series       `json:"data"`
		Error requestErrors `json:"error"`
	}
	if err := c.get(&url.URL{
		Scheme: baseURL[0:5],
		Host:   baseURL[8:],
		Path:   fmt.Sprintf("/series/%d", id),
	}, &r); err != nil {
		return nil, err
	}
	if err := r.Error.GetError(); err != nil {
		return nil, err
	}
	return r.Data, nil
}

type Actor struct {
	ID          uint64 `json:"id"`
	SeriesID    uint64 `json:"seriesId"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	SortOrder   uint   `json:"sortOrder"`
	Image       string `json:"image"`
	ImageAuthor uint64 `json:"imageAuthor"`
	ImageAdded  string `json:"imageAdded"`
	LastUpdated string `json:"lastUpdated"`
}

func (c *Conn) Actors(id uint64) ([]Actor, error) {
	var r struct {
		Data  []Actor       `json:"data"`
		Error requestErrors `json:"error"`
	}
	if err := c.get(&url.URL{
		Scheme: baseURL[0:5],
		Host:   baseURL[8:],
		Path:   fmt.Sprintf("/series/%d/actors", id),
	}, &r); err != nil {
		return nil, err
	}
	if err := r.Error.GetError(); err != nil {
		return nil, err
	}
	return r.Data, nil
}

type SeriesEpisode struct {
	AbsoluteNumber     uint   `json:"absoluteNumber"`
	AiredEpisodeNumber uint   `json:"airedEpisodeNumber"`
	AiredSeason        uint   `json:"airedSeason"`
	DVDEpisodeNumber   uint   `json:"dvdEpisodeNumber"`
	DVDSeason          uint   `json:"dvdSeason"`
	Name               string `json:"episodeName"`
	ID                 uint64 `json:"id"`
	Overview           string `json:"overview"`
}

func (c *Conn) Episodes(id uint64, page uint64) ([]SeriesEpisode, error) {
	return c.episodes(id, make(url.Values), page)
}

func (c *Conn) episodes(id uint64, v url.Values, page uint64) ([]SeriesEpisode, error) {
	path := "/series/%d/episodes/query"
	if len(v) == 0 {
		path = "/series/%d/episodes"
	}
	if page > 0 {
		v.Set("page", strconv.FormatUint(page, 10))
	}
	var r struct {
		Data  []SeriesEpisode `json:"data"`
		Error requestErrors   `json:"error"`
	}
	if err := c.get(&url.URL{
		Scheme:   baseURL[0:5],
		Host:     baseURL[8:],
		Path:     fmt.Sprintf(path, id),
		RawQuery: v.Encode(),
	}, &r); err != nil {
		return nil, err
	}
	if err := r.Error.GetError(); err != nil {
		return nil, err
	}
	return r.Data, nil
}

func (c *Conn) SeasonEpisodes(id uint64, season uint64, page uint64) ([]SeriesEpisode, error) {
	v := make(url.Values)
	v.Set("airedSeason", strconv.FormatUint(season, 10))
	return c.episodes(id, v, page)
}

func (c *Conn) DVDSeasonEpisodes(id uint64, season uint64, page uint64) ([]SeriesEpisode, error) {
	v := make(url.Values)
	v.Set("dvdSeason", strconv.FormatUint(season, 10))
	return c.episodes(id, v, page)
}

func (c *Conn) SeriesEpisode(id uint64, abs uint64) (*SeriesEpisode, error) {
	v := make(url.Values)
	v.Set("absoluteNumber", strconv.FormatUint(abs, 10))
	se, err := c.episodes(id, v, 0)
	if err != nil || len(se) == 0 {
		return nil, err
	}
	return &se[0], nil
}

func (c *Conn) SeasonEpisode(id uint64, season, episode uint64) (*SeriesEpisode, error) {
	v := make(url.Values)
	v.Set("airedSeason", strconv.FormatUint(season, 10))
	v.Set("airedEpisode", strconv.FormatUint(episode, 10))
	se, err := c.episodes(id, v, 0)
	if err != nil || len(se) == 0 {
		return nil, err
	}
	return &se[0], nil
}

func (c *Conn) DVDSeasonEpisode(id uint64, season, episode uint64) (*SeriesEpisode, error) {
	v := make(url.Values)
	v.Set("dvdSeason", strconv.FormatUint(season, 10))
	v.Set("dvdEpisode", strconv.FormatUint(episode, 10))
	se, err := c.episodes(id, v, 0)
	if err != nil || len(se) == 0 {
		return nil, err
	}
	return &se[0], nil
}

type Summary struct {
	AiredSeasons  []string `json:"airedSeasons"`
	AiredEpisodes string   `json:"airedEpisodes"`
	DVDSeasons    []string `json:"dvdSeasons"`
	DVDEpisodes   string   `json:"dvdEpisodes"`
}

func (c *Conn) SeriesSummary(id uint) (*Summary, error) {
	var r struct {
		Data  *Summary      `json:"data"`
		Error requestErrors `json:"error"`
	}
	if err := c.get(&url.URL{
		Scheme: baseURL[0:5],
		Host:   baseURL[8:],
		Path:   fmt.Sprintf("/series/%d/episodes/summary", id),
	}, &r); err != nil {
		return nil, err
	}
	if err := r.Error.GetError(); err != nil {
		return nil, err
	}
	return r.Data, nil
}
