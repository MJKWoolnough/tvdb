package tvdb

type requestErrors struct {
	InvalidFilters     ErrInvalidFilters     `json:"invalidFilters"`
	InvalidLanguage    ErrInvalidLanguage    `json:"invalidLanguage"`
	InvalidQueryParams ErrInvalidQueryParams `json:"invalidQueryParams"`
}

func (r *requestErrors) GetError() error {
	if len(r.InvalidFilters) != 0 {
		return r.InvalidFilters
	} else if len(r.InvalidLanguage) != 0 {
		return r.InvalidLanguage
	} else if len(r.InvalidQueryParams) != 0 {
		return r.InvalidQueryParams
	}
	return nil
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
