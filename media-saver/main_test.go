package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestClassify(t *testing.T) {
	cases := map[string]struct{ platform, url string }{
		"https://www.instagram.com/reel/ABC123/?igsh=xyz": {"instagram", "https://www.instagram.com/reel/ABC123/"},
		"https://vm.tiktok.com/ZMabc/":                    {"tiktok", "https://vm.tiktok.com/ZMabc/"},
		"https://x.com/user/status/123?s=20":              {"x", "https://x.com/user/status/123"},
		"https://fxtwitter.com/user/status/123":           {"x", "https://x.com/user/status/123"},
	}
	for in, want := range cases {
		p, u, err := classify(in)
		if err != nil || p.name != want.platform || u != want.url {
			t.Errorf("classify(%q) = %q, %q, %v; want %q, %q", in, p.name, u, err, want.platform, want.url)
		}
	}
	if _, _, err := classify("https://youtube.com/watch?v=1"); err == nil {
		t.Error("expected unsupported site error")
	}
}

func TestReadURLFromSharedText(t *testing.T) {
	r := httptest.NewRequest("POST", "/save", strings.NewReader("Check out this video! https://vt.tiktok.com/ZSabc/ via TikTok"))
	got, err := readURL(r)
	if err != nil || got != "https://vt.tiktok.com/ZSabc/" {
		t.Fatalf("got %q, %v", got, err)
	}
}
