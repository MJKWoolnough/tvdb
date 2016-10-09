package tvdb

import (
	"fmt"
	"net/url"
	"strconv"
)

type User struct {
	Username             string `json:"userName"`
	Language             string `json:"language"`
	FavoritesDisplaymode string `json:"favoritesDisplaymode"`
}

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

func (c *Conn) Favorites() ([]string, error) {
	var r favourites
	if err := c.get(makeURL("/user/favorites", ""), &r); err != nil {
		return nil, err
	}
	return r.Data.Favorites, nil
}

func (c *Conn) AddFavorite(id uint64) ([]string, error) {
	var r favourites
	if err := c.put(makeURL("/user/favorites/"+strconv.FormatUint(id, 10), ""), &r); err != nil {
		return nil, err
	}
	return r.Data.Favorites, nil
}

func (c *Conn) RemoveFavorite(id uint64) ([]string, error) {
	var r favourites
	if err := c.delete(makeURL("/user/favorites/"+strconv.FormatUint(id, 10), ""), &r); err != nil {
		return nil, err
	}
	return r.Data.Favorites, nil
}

type Rating struct {
	Type   string  `json:"ratingType"`
	ItemID uint64  `json:"ratingItemId"`
	Rating float32 `json:"rating"`
}

type ratings struct {
	Data  []Rating      `json:"data"`
	Error requestErrors `json:"error"`
}

func (c *Conn) Ratings() ([]Rating, error) {
	var r ratings
	if err := c.get(makeURL("/user/ratings", ""), &r); err != nil {
		return nil, err
	}
	return r.Data, nil
}

type RatingItemType string

const (
	RatingSeries  RatingItemType = "series"
	RatingEpisode RatingItemType = "episode"
	RatingBanner  RatingItemType = "banner"
)

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

func (c *Conn) SetRating(rit RatingItemType, id uint64, rating uint32) ([]Rating, error) {
	var r ratings
	if err := c.put(makeURL(fmt.Sprintf("/user/ratings/%s/%d/%d", rit, id, rating), ""), &r); err != nil {
		return nil, err
	}
	return r.Data, nil
}

func (c *Conn) RemoveRating(rit RatingItemType, id uint64) ([]Rating, error) {
	var r ratings
	if err := c.delete(makeURL(fmt.Sprintf("/user/ratings/%s/%d", rit, id), ""), &r); err != nil {
		return nil, err
	}
	return r.Data, nil
}
