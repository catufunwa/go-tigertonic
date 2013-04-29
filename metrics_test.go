package tigertonic

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"
)

func TestCounter(t *testing.T) {
	w := &testResponseWriter{}
	r, _ := http.NewRequest("POST", "http://example.com/foo", bytes.NewBufferString(`{"foo":"bar"}`))
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Content-Type", "application/json")
	counter := Counted(Marshaled(func(u *url.URL, h http.Header, rq *testRequest) (int, http.Header, *testResponse, error) {
		return http.StatusOK, nil, &testResponse{"bar"}, nil
	}), "counted", nil)
	counter.ServeHTTP(w, r)
	if 1 != counter.Count() {
		t.Fatal(counter.Count())
	}
}

func TestTimer(t *testing.T) {
	w := &testResponseWriter{}
	r, _ := http.NewRequest("POST", "http://example.com/foo", bytes.NewBufferString(`{"foo":"bar"}`))
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Content-Type", "application/json")
	timer := Timed(Marshaled(func(u *url.URL, h http.Header, rq *testRequest) (int, http.Header, *testResponse, error) {
		return http.StatusOK, nil, &testResponse{"bar"}, nil
	}), "timed", nil)
	timer.ServeHTTP(w, r)
	if 1 != timer.Count() {
		t.Fatal(timer.Count())
	}
}
