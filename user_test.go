package tvdb_test

import "testing"

func TestUser(t *testing.T) {
	t.Parallel()
	u, err := conn.User()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	if u.Username != auth.Username {
		t.Errorf("expecting username %s, got %s", auth.Username, u.Username)
	}
}
