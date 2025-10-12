package hard

import (
	"testing"
)

func TestMaxValue(t *testing.T) {
	tests := []struct {
		name   string
		events [][]int
		k      int
		want   int
	}{
		{
			name:   "example1",
			events: [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}},
			k:      2,
			want:   7,
		},
		{
			name:   "choose best single large value",
			events: [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 10}},
			k:      2,
			want:   10,
		},
		{
			name:   "single event",
			events: [][]int{{1, 5, 10}},
			k:      1,
			want:   10,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := maxValue(tc.events, tc.k)
			if got != tc.want {
				t.Fatalf("maxValue(%v, %d) = %d, want %d", tc.events, tc.k, got, tc.want)
			}
		})
	}
}
