package tvdb_test

import (
	"encoding/json"
	"os"

	"vimagination.zapto.org/tvdb"
)

var (
	auth = tvdb.Auth{
		APIKey:   os.Getenv("apikey"),
		UserKey:  os.Getenv("userkey"),
		Username: os.Getenv("username"),
	}
	conn *tvdb.Conn
)

func init() {
	if auth.APIKey == "" {
		f, err := os.Open("apikey") // json encoded data {"apikey":"APIKEY"}
		if err != nil {
			panic(err)
		}
		err = json.NewDecoder(f).Decode(&auth)
		f.Close()
		if err != nil {
			panic(err)
		}
	}
}
