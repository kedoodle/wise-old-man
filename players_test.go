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

func TestUpdatePlayerInvalidUsername(t *testing.T) {
	testCases := []struct {
		name                 string
		username             string
		expectedErrorMessage string
	}{
		{
			name:                 "long username",
			username:             "long username",
			expectedErrorMessage: "validation error: username must be between 1 and 12 characters long",
		},
		{
			name:                 "whitespace only username",
			username:             " ",
			expectedErrorMessage: "validation error: username must be defined",
		},
		{
			name:                 "underscore only username",
			username:             "_",
			expectedErrorMessage: "validation error: username must be defined",
		},
		{
			name:                 "dash only username",
			username:             "-",
			expectedErrorMessage: "validation error: username must be defined",
		},
		{
			name:                 "empty username",
			username:             "_ -",
			expectedErrorMessage: "validation error: username must be defined",
		},
		{
			name:                 "special character username",
			username:             "$pecial",
			expectedErrorMessage: "validation error: username cannot contain any special characters",
		},
	}


	client := New()
	ctx := context.Background()
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pd, err := client.UpdatePlayer(ctx, tc.username)
			assert.Nil(t, pd)
			assert.EqualError(t, err, tc.expectedErrorMessage)
		})
	}
}
