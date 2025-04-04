package rope

import "testing"

func TestGetString(t *testing.T) {
	ropeTests := []struct {
		name string
		rope Rope
		want string
	}{
		{name: "First Rope", rope: NewRope(nil, nil, "This is a test"), want: "This is a test"},
	}

	for _, tt := range ropeTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.rope.GetString()
			if got != tt.want {
				t.Errorf("%#v expected %s but got %s", tt.rope, tt.want, got)
			}
		})
	}
}

func TestConcatRopes(t *testing.T) {
	concatTests := []struct {
		name   string
		first  Rope
		second Rope
		want   string
	}{
		{name: "Concat Test 1", first: NewRope(nil, nil, "We like to "), second: NewRope(nil, nil, "move it, move it!"), want: "We like to move it, move it!"},
	}

	for _, tt := range concatTests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConcatRopes(&tt.first, &tt.second)
			if got.GetString() != tt.want {
				t.Errorf("%s and %s not concatenated properly, got: %s", tt.first.GetString(), tt.second.GetString(), got.GetString())
			}
		})
	}
}
