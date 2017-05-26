package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInit(t *testing.T) {
	expected := "test"
	Init(expected)
	actual := uri

	if actual != expected {
		t.Errorf("uri was %s, not %s", actual, expected)
	}
}

func TestNew(t *testing.T) {
	expectedURI := "test-uri"
	expectedID := "test-user"
	Init(expectedURI)
	u := New(expectedID)

	if u.URI != expectedURI {
		t.Errorf("uri was %s, not %s", u.URI, expectedURI)
	}

	if u.ID != expectedID {
		t.Errorf("id was %s, not %s", u.ID, expectedID)
	}
}

func TestGet(t *testing.T) {
	expectedURI := "uri"
	Init(expectedURI)

	expected := New("id")
	expected.Name = "first-name last-name"
	expected.FirstName = "first-name"
	expected.LastName = "last-name"
	expected.Email = "id@example.com"
	expected.Institution = "institution"
	expected.SourceID = "source-id"

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		actualPath := r.URL.Path
		expectedPath := fmt.Sprintf("/subjects/%s", expected.ID)
		if actualPath != expectedPath {
			t.Errorf("path was %s, not %s", actualPath, expectedPath)
		}
		msg, err := json.Marshal(expected)
		if err != nil {
			t.Error(err)
		}
		w.Write(msg)
	}))
	defer srv.Close()

	actual := New("id")
	actual.URI = srv.URL
	err := actual.Get()
	if err != nil {
		t.Error(err)
	}

	if actual.ID != expected.ID {
		t.Errorf("id was %s, not %s", actual.ID, expected.ID)
	}

	if actual.Name != expected.Name {
		t.Errorf("name was %s, not %s", actual.Name, expected.Name)
	}

	if actual.FirstName != expected.FirstName {
		t.Errorf("first name was %s, not %s", actual.FirstName, expected.FirstName)
	}

	if actual.LastName != expected.LastName {
		t.Errorf("last name was %s, not %s", actual.LastName, expected.LastName)
	}

	if actual.Email != expected.Email {
		t.Errorf("email was %s, not %s", actual.Email, expected.Email)
	}

	if actual.Institution != expected.Institution {
		t.Errorf("institution was %s, not %s", actual.Institution, expected.Institution)
	}

	if actual.SourceID != expected.SourceID {
		t.Errorf("source ID was %s, not %s", actual.SourceID, expected.SourceID)
	}
}
