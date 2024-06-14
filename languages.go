package tvdb

import "fmt"

// Language contains information about a supported language.
type Language struct {
	ID           uint64 `json:"id"`
	Abbreviation string `json:"abbreviation"`
	Name         string `json:"name"`
	EnglishName  string `json:"englishName"`
}

var languageURL = makeURL("/languages", "")

// Languages returns a slice of all the languages supported by TVDB.
func (c *Conn) Languages() ([]Language, error) {
	var r struct {
		Data  []Language    `json:"data"`
		Error requestErrors `json:"error"`
	}

	if err := c.get(languageURL, &r); err != nil {
		return nil, err
	}

	return r.Data, nil
}

// Language retrieves information about a specific language, denoted by its
// language id.
func (c *Conn) Language(id uint64) (*Language, error) {
	var r struct {
		Data  *Language     `json:"data"`
		Error requestErrors `json:"error"`
	}

	if err := c.get(makeURL(fmt.Sprintf("/languages/%d", id), ""), &r); err != nil {
		return nil, err
	}

	return r.Data, nil
}
