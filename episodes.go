package tvdb

import "net/url"

type Episode struct {
	ID                 uint     `json:"id"`
	AiredSeason        uint     `json:"airedSeason"`
	AiredEpisodeNumber uint     `json:"airedEpisodeNumber"`
	Name               string   `json:"episodeName"`
	FirstAired         string   `json:"firstAired"`
	GuestStars         []string `json:"guestStars"`
	Directors          []string `json:"directors"`
	Writers            []string `json:"writers"`
	Overview           string   `json:"overview"`
	ProductionCode     string   `json:"productionCode"`
	ShowURL            string   `json:"shorUrl"`
	LastUpdated        int64    `json:"lastUpdated"`
	DVDDiscID          string   `json:"dvdDiscid"`
	DVDSeason          uint     `json:"dvdSeason"`
	DVDEpisodeNumber   uint     `json:"dvdEpisodeNumber"`
	DVDChapter         uint     `json:"dvdChapter"`
	AbsoluteNumber     uint     `json:"absoluteNumber"`
	Filename           string   `json:"filename"`
	SeriesID           string   `json:"seriesId"`
	LastUpdatedBy      string   `json:"lasyUpdatedBy"`
	AirsAfterSeason    uint     `json:"airsAfterSeason"`
	AirsBeforeSeason   uint     `json:"airsBeforeSeason"`
	AirsBeforeEpisode  uint     `json:"airsBeforeEpisode"`
	ThumbAuthor        uint64   `json:"thumbAuthor"`
	ThumbAdded         string   `json:"thumbAdded"`
	ThumbWidth         string   `json:"thumbWidth"`
	ThumbHeight        string   `json:"thumbHeight"`
	IMDBID             string   `json:"imdbId"`
	SiteRating         uint     `json:siteRating"`
	SiteRatingCount    uint64   `json:siteRatingCount"`
}

func (c *Conn) Episode(id uint64) (*Episode, error) {
	var r request
	if err := c.get(&url.URL{}, &r); err != nil {
		return nil, err
	}
	var e Episode
	if err := r.Decode(&e); err != nil {
		return nil, err
	}
	return &e, nil
}
