package tvdb_test

import (
	"testing"

	"github.com/MJKWoolnough/tvdb"
)

func TestLogin(t *testing.T) {
	_, err := tvdb.Login(auth)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	_, err = tvdb.Login(tvdb.Auth{})
	if err != tvdb.ErrInvalidAuth {
		t.Errorf("expecting invalid auth, got %s", err)
	}
}
