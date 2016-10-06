package tvdb_test

import "testing"

func TestActors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		ShowID   uint64
		ActorIDs []uint64
	}{
		{
			71326,
			[]uint64{
				8704,
				8705,
				8706,
				8707,
				8708,
				8709,
				8710,
			},
		},
		{
			78131,
			[]uint64{
				61837,
				61838,
				61839,
				61840,
				61841,
				61842,
				61843,
			},
		},
	}

	for n, test := range tests {
		actors, err := conn.Actors(test.ShowID)
		if err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
			continue
		}
	Loop:
		for _, actorID := range test.ActorIDs {
			for _, actor := range actors {
				if actorID == actor.ID {
					continue Loop
				}
			}
			t.Errorf("test %d: failed to find actor ID %d", n+1, actorID)
		}
	}
}

func TestEpisodes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		ShowID     uint64
		EpisodeIDs []uint64
	}{
		{
			80383,
			[]uint64{
				335586,
				335587,
				335588,
				335589,
				335590,
				335591,
			},
		},
		{
			78131,
			[]uint64{
				272769,
				272775,
				272784,
				272791,
				272770,
				272776,
				272785,
				272792,
				272771,
				272777,
				272786,
				272793,
				272772,
				272778,
				272787,
				272794,
				272773,
				272779,
				272788,
				272795,
				272774,
				272780,
				272789,
				272796,
				272781,
				272790,
				272782,
				272783,
			},
		},
	}

	for n, test := range tests {
		episodes, err := conn.Episodes(test.ShowID, 0)
		if err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
			continue
		}
	Loop:
		for _, episodeID := range test.EpisodeIDs {
			for _, episode := range episodes {
				if episodeID == episode.ID {
					continue Loop
				}
			}
			t.Errorf("test %d: failed to find episode ID %d", n+1, episodeID)
		}
	}
}
