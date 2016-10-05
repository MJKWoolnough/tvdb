package tvdb_test

import "testing"

func TestSearch(t *testing.T) {
	t.Parallel()
	tests := []struct {
		Name string
		IDs  []uint64
	}{
		{
			"Red Dwarf",
			[]uint64{71326, 107321},
		},
		{
			"Coupling",
			[]uint64{78131, 70905},
		},
	}

	for n, test := range tests {
		results, err := conn.Search(test.Name)
		if err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
			continue
		}
	Loop:
		for _, id := range test.IDs {
			for _, result := range results {
				if id == result.ID {
					continue Loop
				}
			}
			t.Errorf("test %d: didn't find ID %d", n+1, id)
		}
	}
}

func TestIMDB(t *testing.T) {
	t.Parallel()
	tests := []struct {
		IMDBID string
		ID     uint64
	}{
		{
			"tt0094535",
			71326,
		},
		{
			"tt0237123",
			78131,
		},
	}

	for n, test := range tests {
		result, err := conn.SearchIMDB(test.IMDBID)
		if err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
			continue
		}
		if result.ID != test.ID {
			t.Errorf("test %d: expecting ID %d, got %d", n+1, test.ID, result.ID)
		}
	}
}
