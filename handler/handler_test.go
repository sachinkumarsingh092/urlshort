package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	fallbackResponse = "fallback"
	path             = "/code_geass"
	dest             = "https://en.wikipedia.org/wiki/Code_Geass"
)

func TestMapHandler(t *testing.T) {
	pathtoUrls := map[string]string{path: dest}

	t.Run("For fallback", func(t *testing.T) {
		result := runMaphandler(pathtoUrls, "/unknown")
		assertbody(t, result, fallbackResponse)
	})

	t.Run("For a known path", func(t *testing.T) {
		result := runMaphandler(pathtoUrls, path)

		assertURL(t, result, dest)
	})
}

func runMaphandler(pathtoUrls map[string]string, path string) *http.Response {
	request, _ := http.NewRequest(http.MethodGet, path, nil)
	response := httptest.NewRecorder()

	mapHandler := MapHandler(pathtoUrls, http.HandlerFunc(fallback))
	mapHandler(response, request)

	return response.Result()
}

func fallback(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fallbackResponse)
}

func assertbody(t *testing.T, resp *http.Response, want string) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal("Could not respond to the body", err)
	}

	got := string(body)
	if want != got {
		t.Errorf("Expected %s, get %s", want, got)
	}
}

func assertURL(t *testing.T, resp *http.Response, want string) {
	url, err := resp.Location()

	if err != nil {
		t.Fatal("Could not read location", err)
	}

	if url.String() != want {
		t.Errorf("Expected %s, got %s", url, want)
	}
}
