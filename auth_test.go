package tvdb_test

import (
	"encoding/json"
	"os"

	"github.com/MJKWoolnough/tvdb"
)

var (
	auth tvdb.Auth
	conn *tvdb.Conn
)

func init() {
	f, err := os.Open("apikey") // json encoded data {"apikey":"APIKEY","username":"USERNAME","userkey":"USERKEY"}
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(f).Decode(&auth)
	f.Close()
	if err != nil {
		panic(err)
	}
}
