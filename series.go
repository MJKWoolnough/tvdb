package tvdb

import (
	"fmt"
	"net/url"
	"strconv"
)

// Series represents all of the information about a particular show
type Series struct {
	ID              uint64   `json:"id"`
	Name            string   `json:"seriesName"`
	Aliases         []string `json:"aliases"`
	Banner          string   `json:"banner"`
	SeriesID        uint64   `json:"seriesId"`
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
	SiteRating      float32  `json:"siteRating"`
	SiteRatingCount uint64   `json:"siteRatingCount"`
}

// Series retrieves the information about a particular series by its ID
func (c *Conn) Series(id uint64) (*Series, error) {
	var r struct {
		Data  *Series       `json:"data"`
		Error requestErrors `json:"error"`
	}
	if err := c.get(makeURL(fmt.Sprintf("/series/%d", id), ""), &r); err != nil {
		return nil, err
	}
	return r.Data, nil
}

// Actor represents all of the information about an actor in a show
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

// Actors returns information about the actors in a show, denoted by its ID
func (c *Conn) Actors(id uint64) ([]Actor, error) {
	var r struct {
		Data  []Actor       `json:"data"`
		Error requestErrors `json:"error"`
	}
	if err := c.get(makeURL(fmt.Sprintf("/series/%d/actors", id), ""), &r); err != nil {
		return nil, err
	}
	return r.Data, nil
}

// SeriesEpisode represents all of the information about a particular episode
// of a show
type SeriesEpisode struct {
	AbsoluteNumber     uint   `json:"absoluteNumber"`
	AiredEpisodeNumber uint   `json:"airedEpisodeNumber"`
	AiredSeason        uint   `json:"airedSeason"`
	DVDEpisodeNumber   uint   `json:"dvdEpisodeNumber"`
	DVDSeason          uint   `json:"dvdSeason"`
	Name               string `json:"episodeName"`
	ID                 uint64 `json:"id"`
	Overview           string `json:"overview"`
	FirstAired         string `json:"firstAired"`
	LastUpdated        uint64 `json:"lastUpdated"`
}

// Episodes returns a paginated view (100 per page) of the episodes in a
// particular series
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
	if err := c.get(makeURL(fmt.Sprintf(path, id), v.Encode()), &r); err != nil {
		return nil, err
	}
	return r.Data, nil
}

// SeasonEpisodes returns a paginated view (100 per page) of the episodes in a
// season of a show
func (c *Conn) SeasonEpisodes(id uint64, season uint64, page uint64) ([]SeriesEpisode, error) {
	v := make(url.Values)
	v.Set("airedSeason", strconv.FormatUint(season, 10))
	return c.episodes(id, v, page)
}

// DVDSeasonEpisodes returns a paginatied view (100 per page) of the episodes
// in the DVD season of a show
func (c *Conn) DVDSeasonEpisodes(id uint64, season uint64, page uint64) ([]SeriesEpisode, error) {
	v := make(url.Values)
	v.Set("dvdSeason", strconv.FormatUint(season, 10))
	return c.episodes(id, v, page)
}

// SeriesEpisode returns the information about a particular episode in a series
// denoted by its absolute episode number
func (c *Conn) SeriesEpisode(id uint64, abs uint64) (*SeriesEpisode, error) {
	v := make(url.Values)
	v.Set("absoluteNumber", strconv.FormatUint(abs, 10))
	se, err := c.episodes(id, v, 0)
	if err != nil || len(se) == 0 {
		return nil, err
	}
	return &se[0], nil
}

// SeasonEpisode returns the information about a particular episode in a series
// denoted by its season and episode numbers
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

// DVDSeasonEpisode returns the information about a particular episode in a
// series denoted by its DVD season and episode numbers
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

// Summary represents the information about episodes for a particular show
type Summary struct {
	AiredSeasons  []string `json:"airedSeasons"`
	AiredEpisodes string   `json:"airedEpisodes"`
	DVDSeasons    []string `json:"dvdSeasons"`
	DVDEpisodes   string   `json:"dvdEpisodes"`
}

// SeriesSummary returns the summary information about episodes for a tv show
func (c *Conn) SeriesSummary(id uint) (*Summary, error) {
	var r struct {
		Data  *Summary      `json:"data"`
		Error requestErrors `json:"error"`
	}
	if err := c.get(makeURL(fmt.Sprintf("/series/%d/episodes/summary", id), ""), &r); err != nil {
		return nil, err
	}
	return r.Data, nil
}
