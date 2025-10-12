package medium

import (
	"testing"
)

func TestMaxEvents(t *testing.T) {
	tests := []struct {
		name   string
		events [][]int
		want   int
	}{
		{
			name:   "example1",
			events: [][]int{{1, 2}, {3, 4}, {2, 3}},
			want:   3,
		},
		{
			name:   "choose best single large value",
			events: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			want:   4,
		},
		{
			name:   "single event",
			events: [][]int{{1, 5}},
			want:   1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := maxEvents(tc.events)
			if got != tc.want {
				t.Fatalf("maxEvents(%v) = %d, want %d", tc.events, got, tc.want)
			}
		})
	}
}
