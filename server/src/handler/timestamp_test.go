package handler

import (
	"regexp"
	"testing"
)

func TestAddTimestamp(t *testing.T) {
	got := AddTimestamp("testing.pdf")
	matched, err := regexp.Match(`testing.pdf_\d+`, []byte(got))

	if !matched {
		t.Errorf("AddTimestamp(%q) == %q, want %q", "testing.pdf", got, err)
	}
}

func TestRemoveTimestamp(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"testing_2.pdf_1592176571433757000", "testing_2.pdf"},
		{"testing.pdf_1592176571433757000", "testing.pdf"},
		{"", ""},
	}

	for _, c := range cases {
		got := RemoveTimestamp(c.in)
		if got != c.want {
			t.Errorf("RemoveTimestamp(%q) == %q, want %q", c.in, got, c.want)
		}
	}

}
