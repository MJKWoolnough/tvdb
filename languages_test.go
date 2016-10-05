package tvdb_test

import "testing"

func TestLanguage(t *testing.T) {
	t.Parallel()
	tests := []struct {
		ID   uint64
		Abbr string
	}{
		{7, "en"},
		{13, "nl"},
		{27, "zh"},
		{28, "cs"},
	}
	for n, test := range tests {
		l, err := conn.Language(test.ID)
		if err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		}
		if l.Abbreviation != test.Abbr {
			t.Errorf("test %d: expecting abbr. %q, got %q", n+1, test.Abbr, l.Abbreviation)
		}
	}
}
