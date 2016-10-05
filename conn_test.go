package tvdb_test

import (
	"testing"

	"github.com/MJKWoolnough/tvdb"
)

func TestLogin(t *testing.T) {
	_, err := tvdb.Login(tvdb.Auth{})
	if err != tvdb.ErrInvalidAuth {
		t.Errorf("expecting invalid auth, got %s", err)
	}
	c, err := tvdb.Login(auth)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	conn = c
}

func TestRefresh(t *testing.T) {
	t.Parallel()
	if err := conn.Refresh(); err != nil {
		t.Errorf("unexpected refresh error: %s", err)
	}
}
