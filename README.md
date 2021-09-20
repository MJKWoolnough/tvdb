# tvdb
--
    import "vimagination.zapto.org/tvdb"

Package tvdb is a simple interface to the TVDB database of TV shows

## Usage

```go
var (
	ErrInvalidAuth  = errors.New("invalid credentials")
	ErrUnknownError = errors.New("unknown error")
	ErrNotFound     = errors.New("not found")
)
```
Errors

#### type Actor

```go
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
```

Actor represents all of the information about an actor in a show

#### type Auth

```go
type Auth struct {
	APIKey   string `json:"apikey"`
	Username string `json:"username,omitempty"`
	UserKey  string `json:"userkey,omitempty"`
}
```

Auth represents the information required to get a validated authentication
token.

#### type Conn

```go
type Conn struct {
}
```

Conn represents a connection to the TVDB database

#### func  Login

```go
func Login(a Auth) (*Conn, error)
```
Login creates a TVDB database connection using login credentials

#### func  Token

```go
func Token(t string) *Conn
```
Token creates a TVDB database connection using a pre-authorised token

#### func (*Conn) Actors

```go
func (c *Conn) Actors(id uint64) ([]Actor, error)
```
Actors returns information about the actors in a show, denoted by its ID

#### func (*Conn) AddFavorite

```go
func (c *Conn) AddFavorite(id uint64) error
```
AddFavorite adds a show id to the list of user favorites

#### func (*Conn) DVDSeasonEpisode

```go
func (c *Conn) DVDSeasonEpisode(id uint64, season, episode uint64) (*SeriesEpisode, error)
```
DVDSeasonEpisode returns the information about a particular episode in a series
denoted by its DVD season and episode numbers

#### func (*Conn) DVDSeasonEpisodes

```go
func (c *Conn) DVDSeasonEpisodes(id uint64, season uint64, page uint64) ([]SeriesEpisode, error)
```
DVDSeasonEpisodes returns a paginatied view (100 per page) of the episodes in
the DVD season of a show

#### func (*Conn) Episode

```go
func (c *Conn) Episode(id uint64) (*Episode, error)
```
Episode returns the information about a single tv episode denoted by the episode
id

#### func (*Conn) Episodes

```go
func (c *Conn) Episodes(id uint64, page uint64) ([]SeriesEpisode, error)
```
Episodes returns a paginated view (100 per page) of the episodes in a particular
series

#### func (*Conn) Favorites

```go
func (c *Conn) Favorites() ([]uint64, error)
```
Favorites returns a list of show ids that the user has set as favorites

#### func (*Conn) Language

```go
func (c *Conn) Language(id uint64) (*Language, error)
```
Language retrieves information about a specific language, denoted by its
language id

#### func (*Conn) Languages

```go
func (c *Conn) Languages() ([]Language, error)
```
Languages returns a slice of all the languages supported by TVDB

#### func (*Conn) Ratings

```go
func (c *Conn) Ratings() ([]Rating, error)
```
Ratings returns a list of ratings that the user has set

#### func (*Conn) RatingsByType

```go
func (c *Conn) RatingsByType(rit RatingItemType) ([]Rating, error)
```
RatingsByType returns a list of ratings for a specific type

#### func (*Conn) Refresh

```go
func (c *Conn) Refresh() error
```
Refresh retrieves a new authentication token without having to use the login
credentials. Each token only lasts 24 hours and refresh can only be used in that
time-frame

#### func (*Conn) RemoveFavorite

```go
func (c *Conn) RemoveFavorite(id uint64) error
```
RemoveFavorite removes a show id to the list of user favorites

#### func (*Conn) RemoveRating

```go
func (c *Conn) RemoveRating(rit RatingItemType, id uint64) error
```
RemoveRating removes a user rating for a specific series, episode or banner

#### func (*Conn) Search

```go
func (c *Conn) Search(name string) ([]Search, error)
```
Search searches the TVDB database for shows with the given name

#### func (*Conn) SearchIMDB

```go
func (c *Conn) SearchIMDB(imdb string) (*Search, error)
```
SearchIMDB searches the TVDB database for the show corrensponding to the given
IMDB ID

#### func (*Conn) SearchZap2It

```go
func (c *Conn) SearchZap2It(zapit string) (*Search, error)
```
SearchZap2It searches the TVDB database for the show corrensponding to the given
Zap2It ID

#### func (*Conn) SeasonEpisode

```go
func (c *Conn) SeasonEpisode(id uint64, season, episode uint64) (*SeriesEpisode, error)
```
SeasonEpisode returns the information about a particular episode in a series
denoted by its season and episode numbers

#### func (*Conn) SeasonEpisodes

```go
func (c *Conn) SeasonEpisodes(id uint64, season uint64, page uint64) ([]SeriesEpisode, error)
```
SeasonEpisodes returns a paginated view (100 per page) of the episodes in a
season of a show

#### func (*Conn) Series

```go
func (c *Conn) Series(id uint64) (*Series, error)
```
Series retrieves the information about a particular series by its ID

#### func (*Conn) SeriesEpisode

```go
func (c *Conn) SeriesEpisode(id uint64, abs uint64) (*SeriesEpisode, error)
```
SeriesEpisode returns the information about a particular episode in a series
denoted by its absolute episode number

#### func (*Conn) SeriesSummary

```go
func (c *Conn) SeriesSummary(id uint) (*Summary, error)
```
SeriesSummary returns the summary information about episodes for a tv show

#### func (*Conn) SetLanguage

```go
func (c *Conn) SetLanguage(code string)
```
SetLanguage sets the language header used by some queries to return information
in the requested language

#### func (*Conn) SetRating

```go
func (c *Conn) SetRating(rit RatingItemType, id uint64, rating uint32) error
```
SetRating sets a user rating for a specific series, episode or banner

#### func (*Conn) Token

```go
func (c *Conn) Token() string
```
Token returns the current authentication token

#### func (*Conn) User

```go
func (c *Conn) User() (*User, error)
```
User returns the logged in user details

#### type Episode

```go
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
	SiteRating        float32 `json:"siteRating"`
	SiteRatingCount   uint64  `json:"siteRatingCount"`
}
```

Episode represents the data for a single episode of a programme

#### type ErrInvalidFilters

```go
type ErrInvalidFilters []string
```

ErrInvalidFilters is returned from query that use an unknown or invalid filter

#### func (ErrInvalidFilters) Error

```go
func (ErrInvalidFilters) Error() string
```
Error satisfies the error interface

#### func (ErrInvalidFilters) UnmarshalJSON

```go
func (ErrInvalidFilters) UnmarshalJSON(b []byte) error
```
UnmarshalJSON uses the json decoding of the error to generate an error instead
of decoding

#### type ErrInvalidLanguage

```go
type ErrInvalidLanguage string
```

ErrInvalidLanguage is returned when a query requests an unknown or invalid
language

#### func (ErrInvalidLanguage) Error

```go
func (ErrInvalidLanguage) Error() string
```
Error satisfies the error interface

#### func (ErrInvalidLanguage) UnmarshalJSON

```go
func (ErrInvalidLanguage) UnmarshalJSON(b []byte) error
```
UnmarshalJSON uses the json decoding of the error to generate an error instead
of decoding

#### type ErrInvalidQueryParams

```go
type ErrInvalidQueryParams []string
```

ErrInvalidQueryParams is returned when a query uses unknown or invalid params

#### func (ErrInvalidQueryParams) Error

```go
func (ErrInvalidQueryParams) Error() string
```
Error satisfies the error interface

#### func (ErrInvalidQueryParams) UnmarshalJSON

```go
func (ErrInvalidQueryParams) UnmarshalJSON(b []byte) error
```
UnmarshalJSON uses the json decoding of the error to generate an error instead
of decoding

#### type Language

```go
type Language struct {
	ID           uint64 `json:"id"`
	Abbreviation string `json:"abbreviation"`
	Name         string `json:"name"`
	EnglishName  string `json:"englishName"`
}
```

Language contains information about a supported language

#### type Rating

```go
type Rating struct {
	Type   string  `json:"ratingType"`
	ItemID uint64  `json:"ratingItemId"`
	Rating float32 `json:"rating"`
}
```

Rating represents a single rating for an item

#### type RatingItemType

```go
type RatingItemType string
```

RatingItemType represents the type of rating, currently one of series, episode
and banner

```go
const (
	RatingSeries  RatingItemType = "series"
	RatingEpisode RatingItemType = "episode"
	RatingBanner  RatingItemType = "banner"
)
```
The currently available Item Types for Rating

#### type Search

```go
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
```

Search is a representation of the data returned from a tv show search

#### type Series

```go
type Series struct {
	ID              uint64   `json:"id"`
	Name            string   `json:"seriesName"`
	Aliases         []string `json:"aliases"`
	Banner          string   `json:"banner"`
	SeriesID        string   `json:"seriesId"`
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
```

Series represents all of the information about a particular show

#### type SeriesEpisode

```go
type SeriesEpisode struct {
	AbsoluteNumber     uint    `json:"absoluteNumber"`
	AiredEpisodeNumber uint    `json:"airedEpisodeNumber"`
	AiredSeason        uint    `json:"airedSeason"`
	DVDEpisodeNumber   float32 `json:"dvdEpisodeNumber"`
	DVDSeason          uint    `json:"dvdSeason"`
	Name               string  `json:"episodeName"`
	ID                 uint64  `json:"id"`
	Overview           string  `json:"overview"`
	FirstAired         string  `json:"firstAired"`
	LastUpdated        uint64  `json:"lastUpdated"`
}
```

SeriesEpisode represents all of the information about a particular episode of a
show

#### type Summary

```go
type Summary struct {
	AiredSeasons  []string `json:"airedSeasons"`
	AiredEpisodes string   `json:"airedEpisodes"`
	DVDSeasons    []string `json:"dvdSeasons"`
	DVDEpisodes   string   `json:"dvdEpisodes"`
}
```

Summary represents the information about episodes for a particular show

#### type User

```go
type User struct {
	Username             string `json:"userName"`
	Language             string `json:"language"`
	FavoritesDisplaymode string `json:"favoritesDisplaymode"`
}
```

User represents the user details given by the user endpoint
