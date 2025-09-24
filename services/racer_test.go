package services

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare the two servers", func(t *testing.T) {
		first := createServer(20 * time.Millisecond)
		second := createServer(5 * time.Millisecond)

		defer first.Close()
		defer second.Close()

		got, _ := Race(first.URL, second.URL)
		want := second.URL

		if got != want {
			t.Errorf("wanted %q, but ended up getting %q", want, got)
		}
	})

	t.Run("exception occurs if response is longer than 10s", func(t *testing.T) {
		first := createServer(20 * time.Millisecond)
		second := createServer(5 * time.Millisecond)

		defer first.Close()
		defer second.Close()

		_, err := Race(first.URL, second.URL)

		if err == nil {
			t.Error("expected an error but didn't get one.")
		}

	})

}

func createServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
