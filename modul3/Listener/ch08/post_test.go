package ch08

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type User struct {
	FirstName string
	LastName string
}

// Dette er serversiden
func handlePostUser(t *testing.T) func (http.ResponseWriter, *http.Request) { // Poste brukerdata over til server
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(r io.ReadCloser) {
			_, _ = io.Copy(ioutil.Discard, r)
			_ = r.Close()
		}(r.Body)

		if r.Method != http.MethodPost {
			http.Error(w, "", http.StatusMethodNotAllowed)
			return
		}

		var u User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			t.Error(err)
			http.Error(w, "Dekoding feilet", http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusAccepted)
	}
}

func TestPostUser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handlePostUser(t)))
	defer ts.Close()							// venter på at alt annet er ferdig (defer) og så avslutter tilkoblinga

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("Forventet status %d; fikk status %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}

	// buffer, mellomlagring
	buf := new(bytes.Buffer)
	u := User{FirstName: "Tom", LastName: "Jensen"}
	err = json.NewEncoder(buf).Encode(&u)							// Legger til en json fil i buffer
	if err != nil {
		t.Fatal(err)
	}

	resp, err = http.Post(ts.URL, "application/json", buf) // Spesifiserer hva slags type data serveren må handle
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusAccepted {
		t.Fatalf("Forventet status %d; fikk status %d", http.StatusAccepted, resp.StatusCode)
	}
	_=resp.Body.Close()
}

// func TestMultipartPost(t *testing.T)