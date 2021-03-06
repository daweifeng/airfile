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

func TestFetchTimestamp(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"testing_2.pdf_1592176571433757000", "1592176571433757000"},
		{"testing.pdf_1592176571433757000", "1592176571433757000"},
		{"", ""},
	}

	for _, c := range cases {
		got := FetchTimestamp(c.in)
		if got != c.want {
			t.Errorf("RemoveTimestamp(%q) == %q, want %q", c.in, got, c.want)
		}
	}

}

func TestGenerateTime(t *testing.T) {
	cases := []struct {
		in   string
		want int64
	}{
		{"1592176571433757000", 1592176571433757000},
		{"1592176571433757000", 1592176571433757000},
	}

	for _, c := range cases {
		got := GenerateTime(c.in)
		if got.UnixNano() != c.want {
			t.Errorf("GenerateTime(%q) == %q, want %q", c.in, got.UnixNano(), c.want)
		}
	}

}
