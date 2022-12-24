package wiseoldman

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdatePlayer(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/players/zezima", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		w.WriteHeader(http.StatusOK)
		body, err := os.ReadFile("testdata/player_details.json")
		if err != nil {
			t.Fatal(err)
		}
		w.Write(body)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := &Client{
		httpClient: http.DefaultClient,
		baseURL:    server.URL,
	}
	ctx := context.Background()

	pd, err := client.UpdatePlayer(ctx, "zezima")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1135, pd.ID)
	assert.Equal(t, "zezima", pd.Username)
	assert.Equal(t, 93, pd.CombatLevel)
	assert.NotNil(t, pd.LatestSnapshot)
}

func TestUpdatePlayerUsernameTooLong(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/players/long%20username", r.URL.RequestURI())
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"Validation error: Username must be between 1 and 12 characters long."}`))
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := &Client{
		httpClient: http.DefaultClient,
		baseURL:    server.URL,
	}
	ctx := context.Background()

	pd, err := client.UpdatePlayer(ctx, "long username")
	assert.Nil(t, pd)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Username must be between 1 and 12 characters long")
}
