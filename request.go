package tvdb

import "encoding/json"

type requestErrors struct {
	InvalidFilters     ErrInvalidFilters     `json:"invalidFilters"`
	InvalidLanguage    ErrInvalidLanguage    `json:"invalidLanguage"`
	InvalidQueryParams ErrInvalidQueryParams `json:"invalidQueryParams"`
}

type ErrInvalidFilters []string

func (ErrInvalidFilters) Error() string {
	return "invalid filters"
}

func (ErrInvalidFilters) UnmarshalJSON(b []byte) error {
	var e []string
	if err := json.Unmarshal(b, &e); err != nil {
		return err
	} else if len(e) > 0 {
		return ErrInvalidFilters(e)
	}
	return nil
}

type ErrInvalidLanguage string

func (ErrInvalidLanguage) Error() string {
	return "invalid language"
}

func (ErrInvalidLanguage) UnmarshalJSON(b []byte) error {
	var e string
	if err := json.Unmarshal(b, &e); err != nil {
		return err
	} else if len(e) > 0 {
		return ErrInvalidLanguage(e)
	}
	return nil
}

type ErrInvalidQueryParams []string

func (ErrInvalidQueryParams) Error() string {
	return "invalid query params"
}

func (ErrInvalidQueryParams) UnmarshalJSON(b []byte) error {
	var e []string
	if err := json.Unmarshal(b, &e); err != nil {
		return err
	} else if len(e) > 0 {
		return ErrInvalidQueryParams(e)
	}
	return nil
}
