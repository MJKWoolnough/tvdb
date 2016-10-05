package tvdb_test

import (
	"testing"

	"github.com/MJKWoolnough/tvdb"
)

func TestLogin(t *testing.T) {
	t.Parallel()
	_, err := tvdb.Login(auth)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	_, err = tvdb.Login(tvdb.Auth{})
	if err != tvdb.ErrInvalidAuth {
		t.Errorf("expecting invalid auth, got %s", err)
	}
}

func TestRefresh(t *testing.T) {
	t.Parallel()
	c, err := tvdb.Login(auth)
	if err != nil {
		t.Errorf("unexpected login error: %s", err)
	}
	if err := c.Refresh(); err != nil {
		t.Errorf("unexpected refresh error: %s", err)
	}
}
