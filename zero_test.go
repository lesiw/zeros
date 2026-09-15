package zeros

import "testing"

func TestIs(t *testing.T) {
	cases := []struct {
		got  int
		want bool
	}{
		{0, true},
		{1, false},
		{-1, false},
	}
	for _, tc := range cases {
		if got := Is(tc.got); got != tc.want {
			t.Errorf("Is(%v) = %v, want %v", tc.got, got, tc.want)
		}
	}
}
