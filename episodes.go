package tvdb

import (
	"fmt"
	"net/url"
)

type Episode struct {
	ID                 uint64   `json:"id"`
	AiredSeason        uint     `json:"airedSeason"`
	AiredSeasonID      uint64   `json:"airedSeasonId"`
	AiredEpisodeNumber uint     `json:"airedEpisodeNumber"`
	Name               string   `json:"episodeName"`
	FirstAired         string   `json:"firstAired"`
	GuestStars         []string `json:"guestStars"`
	Directors          []string `json:"directors"`
	Writers            []string `json:"writers"`
	Overview           string   `json:"overview"`
	Language           struct {
		EpisodeName string `json:"episodeName"`
		Overview    string `json:"overview"`
	} `json:"language"`
	ProductionCode    string  `json:"productionCode"`
	ShowURL           string  `json:"shorUrl"`
	LastUpdated       int64   `json:"lastUpdated"`
	DVDDiscID         string  `json:"dvdDiscid"`
	DVDSeason         uint    `json:"dvdSeason"`
	DVDEpisodeNumber  uint    `json:"dvdEpisodeNumber"`
	DVDChapter        uint    `json:"dvdChapter"`
	AbsoluteNumber    uint    `json:"absoluteNumber"`
	Filename          string  `json:"filename"`
	SeriesID          uint64  `json:"seriesId"`
	LastUpdatedBy     uint64  `json:"lastUpdatedBy"`
	AirsAfterSeason   uint    `json:"airsAfterSeason"`
	AirsBeforeSeason  uint    `json:"airsBeforeSeason"`
	AirsBeforeEpisode uint    `json:"airsBeforeEpisode"`
	ThumbAuthor       uint64  `json:"thumbAuthor"`
	ThumbAdded        string  `json:"thumbAdded"`
	ThumbWidth        string  `json:"thumbWidth"`
	ThumbHeight       string  `json:"thumbHeight"`
	IMDBID            string  `json:"imdbId"`
	SiteRating        float32 `json:siteRating"`
	SiteRatingCount   uint64  `json:siteRatingCount"`
}

func (c *Conn) Episode(id uint64) (*Episode, error) {
	var r struct {
		Data  *Episode `json:"data"`
		Error requestErrors
	}
	if err := c.get(&url.URL{
		Scheme: baseURL[0:5],
		Host:   baseURL[8:],
		Path:   fmt.Sprintf("/episodes/%d", id),
	}, &r); err != nil {
		return nil, err
	}
	if err := r.Error.GetError(); err != nil {
		return nil, err
	}
	return r.Data, nil
}
