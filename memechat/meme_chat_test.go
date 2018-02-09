package main

import (
	"log"
	"testing"
)

func TestGetMemeImageURL(t *testing.T) {
	log.Println("=== Test getMemeImageURL ===")

	cases := []struct {
		in, want string
	}{
		{"test/test2/memegenerator/meme_name", "meme_name"},
		{"test/test2/memegenerator/", ""},
		{"test/test2/meme/<meme_name>", "<meme_name>"},
	}

	for _, c := range cases {
		got := getMemeImageURL(c.in)
		if got != c.want {
			t.Errorf("getMemeImageURL(%q) == %q, want %q", c.in, got, c.want)
			t.Fail()
		}
	}
}
