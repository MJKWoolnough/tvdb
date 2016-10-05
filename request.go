package tvdb

import "encoding/json"

type request struct {
	Data   json.RawMessage `json:"data"`
	Errors struct {
		InvalidFilters     ErrInvalidFilters     `json:"invalidFilters"`
		InvalidLanguage    ErrInvalidLanguage    `json:"invalidLanguage"`
		InvalidQueryParams ErrInvalidQueryParams `json:"invalidQueryParams"`
	} `json:"errors"`
}

func (r *request) Decode(v interface{}) error {
	if len(r.Errors.InvalidFilters) == 0 {
		return r.Errors.InvalidFilters
	} else if len(r.Errors.InvalidLanguage) == 0 {
		return r.Errors.InvalidLanguage
	} else if len(r.Errors.InvalidQueryParams) == 0 {
		return r.Errors.InvalidQueryParams
	}
	return json.Unmarshal(r.Data, v)
}

type ErrInvalidFilters []string

func (ErrInvalidFilters) Error() string {
	return "invalid filters"
}

type ErrInvalidLanguage string

func (ErrInvalidLanguage) Error() string {
	return "invalid language"
}

type ErrInvalidQueryParams []string

func (ErrInvalidQueryParams) Error() string {
	return "invalid query params"
}
