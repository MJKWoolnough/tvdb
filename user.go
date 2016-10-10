package tvdb

import (
	"fmt"
	"net/url"
	"strconv"
)

// User represents the user details given by the user endpoint
type User struct {
	Username             string `json:"userName"`
	Language             string `json:"language"`
	FavoritesDisplaymode string `json:"favoritesDisplaymode"`
}

// User returns the logged in user details
func (c *Conn) User() (*User, error) {
	var r struct {
		Data  *User         `json:"data"`
		Error requestErrors `json:"error"`
	}
	if err := c.get(makeURL("/user", ""), &r); err != nil {
		return nil, err
	}
	return r.Data, nil
}

type favourites struct {
	Data struct {
		Favorites []string `json:"favorites"`
	} `json:"data"`
	Error requestErrors `json:"error"`
}

// Favorites returns a list of show ids that the user has set as favorites
func (c *Conn) Favorites() ([]uint64, error) {
	var r favourites
	if err := c.get(makeURL("/user/favorites", ""), &r); err != nil {
		return nil, err
	}
	ids := make([]uint64, 0, len(r.Data.Favorites))
	for _, id := range r.Data.Favorites {
		if id == "" {
			continue
		}
		i, _ := strconv.ParseUint(id, 10, 64)
		ids = append(ids, i)
	}
	return ids, nil
}

// AddFavorite adds a show id to the list of user favorites
func (c *Conn) AddFavorite(id uint64) error {
	return c.put(makeURL("/user/favorites/"+strconv.FormatUint(id, 10), ""), new(favourites))
}

// RemoveFavorite removes a show id to the list of user favorites
func (c *Conn) RemoveFavorite(id uint64) error {
	return c.delete(makeURL("/user/favorites/"+strconv.FormatUint(id, 10), ""), new(favourites))
}

// Rating represents a single rating for an item
type Rating struct {
	Type   string  `json:"ratingType"`
	ItemID uint64  `json:"ratingItemId"`
	Rating float32 `json:"rating"`
}

type ratings struct {
	Data  []Rating      `json:"data"`
	Error requestErrors `json:"error"`
}

// Ratings returns a list of ratings that the user has set
func (c *Conn) Ratings() ([]Rating, error) {
	var r ratings
	if err := c.get(makeURL("/user/ratings", ""), &r); err != nil {
		return nil, err
	}
	return r.Data, nil
}

// RatingItemType represents the type of rating, currently one of series,
// episode and banner
type RatingItemType string

// The currently available Item Types for Rating
const (
	RatingSeries  RatingItemType = "series"
	RatingEpisode RatingItemType = "episode"
	RatingBanner  RatingItemType = "banner"
)

// RatingsByType returns a list of ratings for a specific type
func (c *Conn) RatingsByType(rit RatingItemType) ([]Rating, error) {
	var (
		r ratings
		v url.Values
	)
	v.Set("itemType", string(rit))
	if err := c.get(makeURL("/user/ratings/query", v.Encode()), &r); err != nil {
		return nil, err
	}
	return r.Data, nil
}

// SetRating sets a user rating for a specific series, episode or banner
func (c *Conn) SetRating(rit RatingItemType, id uint64, rating uint32) error {
	return c.put(makeURL(fmt.Sprintf("/user/ratings/%s/%d/%d", rit, id, rating), ""), new(ratings))
}

// RemoveRating removes a user rating for a specific series, episode or banner
func (c *Conn) RemoveRating(rit RatingItemType, id uint64) error {
	return c.delete(makeURL(fmt.Sprintf("/user/ratings/%s/%d", rit, id), ""), new(ratings))
}
