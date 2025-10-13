package hard

import "testing"

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
		{"", "a*", true},
		{"", "", true},
		{"abc", "abc", true},
		{"abcd", "d*", false},
		{"abcdefg", ".*", true},
	}

	for _, tc := range tests {
		got := isMatch(tc.s, tc.p)
		if got != tc.want {
			t.Errorf("isMatch(%q, %q) = %v, want %v", tc.s, tc.p, got, tc.want)
		}
	}
}
