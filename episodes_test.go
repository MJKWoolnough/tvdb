package tvdb_test

import (
	"testing"

	"github.com/MJKWoolnough/tvdb"
)

func TestEpisode(t *testing.T) {
	t.Parallel()
	tests := []tvdb.Episode{
		{
			ID:                 272788,
			AiredSeason:        3,
			AiredSeasonID:      14342,
			AiredEpisodeNumber: 5,
			DVDSeason:          3,
			DVDEpisodeNumber:   5,
			AbsoluteNumber:     20,
			SeriesID:           78131,
		},
		{
			ID:                 174095,
			AiredSeason:        1,
			AiredSeasonID:      9090,
			AiredEpisodeNumber: 1,
			DVDSeason:          1,
			DVDEpisodeNumber:   1,
			SeriesID:           75628,
		},
		{
			ID:                 42378,
			AiredSeason:        4,
			AiredSeasonID:      2079,
			AiredEpisodeNumber: 4,
			DVDSeason:          4,
			DVDEpisodeNumber:   4,
			SeriesID:           71326,
		},
		{
			ID:                 4562135,
			AiredSeason:        0,
			AiredSeasonID:      26260,
			AiredEpisodeNumber: 84,
			SeriesID:           78804,
		},
	}

	for n, episode := range tests {
		e, err := conn.Episode(episode.ID)
		if err != nil {
			t.Errorf("unexpected error: %#v", err)
		}
		if e.ID != episode.ID {
			t.Errorf("test %d: inconsistant ID, expecting %d, got %d", n+1, episode.ID, e.ID)
		}
		if e.AiredSeason != episode.AiredSeason {
			t.Errorf("test %d: expecting Aired Season %d, got %d", n+1, episode.AiredSeason, e.AiredSeason)
		}
		if e.AiredSeasonID != episode.AiredSeasonID {
			t.Errorf("test %d: expecting Aired Season ID %d, got %d", n+1, episode.AiredSeasonID, e.AiredSeasonID)
		}
		if e.AiredEpisodeNumber != episode.AiredEpisodeNumber {
			t.Errorf("test %d: expecting Aired Episode Number %d, got %d", n+1, episode.AiredEpisodeNumber, e.AiredEpisodeNumber)
		}
		if e.DVDSeason != episode.DVDSeason {
			t.Errorf("test %d: expecting DVD Season %d, got %d", n+1, episode.DVDSeason, e.DVDSeason)
		}
		if e.DVDEpisodeNumber != episode.DVDEpisodeNumber {
			t.Errorf("test %d: expecting DVD Episode Number %d, got %d", n+1, episode.DVDEpisodeNumber, e.DVDEpisodeNumber)
		}
		if e.AbsoluteNumber != episode.AbsoluteNumber {
			t.Errorf("test %d: expecting Absolute Number %d, got %d", n+1, episode.AbsoluteNumber, e.AbsoluteNumber)
		}
		if e.SeriesID != episode.SeriesID {
			t.Errorf("test %d: expecting Series ID %d, got %d", n+1, episode.SeriesID, e.SeriesID)
		}
	}
}
