package tvdb

import (
	"encoding/json"
	"fmt"
)

type requestErrors struct {
	InvalidFilters     ErrInvalidFilters     `json:"invalidFilters"`
	InvalidLanguage    ErrInvalidLanguage    `json:"invalidLanguage"`
	InvalidQueryParams ErrInvalidQueryParams `json:"invalidQueryParams"`
}

// ErrInvalidFilters is returned from query that use an unknown or invalid
// filter
type ErrInvalidFilters []string

// Error satisfies the error interface
func (ErrInvalidFilters) Error() string {
	return "invalid filters"
}

// UnmarshalJSON uses the json decoding of the error to generate an error
// instead of decoding
func (ErrInvalidFilters) UnmarshalJSON(b []byte) error {
	var e []string
	if err := json.Unmarshal(b, &e); err != nil {
		return fmt.Errorf("error decoding Invalid Filters error: %w", err)
	} else if len(e) > 0 {
		return ErrInvalidFilters(e)
	}
	return nil
}

// ErrInvalidLanguage is returned when a query requests an unknown or invalid
// language
type ErrInvalidLanguage string

// Error satisfies the error interface
func (ErrInvalidLanguage) Error() string {
	return "invalid language"
}

// UnmarshalJSON uses the json decoding of the error to generate an error
// instead of decoding
func (ErrInvalidLanguage) UnmarshalJSON(b []byte) error {
	var e string
	if err := json.Unmarshal(b, &e); err != nil {
		return fmt.Errorf("error decoding Invalid Language error: %w", err)
	} else if len(e) > 0 {
		return ErrInvalidLanguage(e)
	}
	return nil
}

// ErrInvalidQueryParams is returned when a query uses unknown or invalid
// params
type ErrInvalidQueryParams []string

// Error satisfies the error interface
func (ErrInvalidQueryParams) Error() string {
	return "invalid query params"
}

// UnmarshalJSON uses the json decoding of the error to generate an error
// instead of decoding
func (ErrInvalidQueryParams) UnmarshalJSON(b []byte) error {
	var e []string
	if err := json.Unmarshal(b, &e); err != nil {
		return fmt.Errorf("error decoding Invalid Query Params error: %w", err)
	} else if len(e) > 0 {
		return ErrInvalidQueryParams(e)
	}
	return nil
}
