package main

import (
	"log"
	"testing"
)

func TestParseMemeURL(t *testing.T) {
	log.Println("=== Test getMemeImageURL ===")

	cases := []struct {
		in, want string
	}{
		{"test/test2/memegenerator/meme_name", "meme_name"},
		{"test/test2/memegenerator/", ""},
		{"test/test2/meme/<meme_name>", "<meme_name>"},
		{"", ""},
	}

	for _, c := range cases {
		got := parseMemeURL(c.in)
		if got != c.want {
			t.Errorf("parseMemeURL(%q) == %q, want %q", c.in, got, c.want)
			t.Fail()
		}
	}
}

func TestGetMemePageURL(t *testing.T) {
	log.Println("=== Test getMemePageURL ===")
	cases := []struct {
		in, want string
	}{
		{"spongebob", "/memegenerator/Mocking-Spongebob"},
		{"caveman spongebob", "/memegenerator/Spongegar"},
		{"", ""},
	}

	for _, c := range cases {
		got, _ := getMemePageURL(c.in)
		if got != c.want {
			t.Errorf("getMemePageURL(%q) == %q, want %q", c.in, got, c.want)
			t.Fail()
		}
	}
}

func TestGetMemeImageURL(t *testing.T) {
	log.Println("=== Test GetMemeImageURL ===")
	cases := []struct {
		in, want string
	}{
		{"spongebob", "https://imgflip.com/s/meme/Mocking-Spongebob.jpg"},
		{"caveman spongebob", "https://imgflip.com/s/meme/Spongegar.jpg"},
		{"", "https://imgflip.com/s/meme/.jpg"},
	}

	for _, c := range cases {
		got := GetMemeImageURL(c.in)
		if got != c.want {
			t.Errorf("GetMemeImageURL(%q) == %q, want %q", c.in, got, c.want)
			t.Fail()
		}
	}
}
