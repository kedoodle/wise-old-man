package wiseoldman

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

func (c *Client) UpdatePlayer(ctx context.Context, username string) (*PlayerDetails, error) {
	var pd PlayerDetails

	base, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}
	url := base.JoinPath("players", username)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()	

	if resp.StatusCode != http.StatusOK {
		
		var e = Error{}
		err = json.NewDecoder(resp.Body).Decode(&e)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(e.Message)
	}

	err = json.NewDecoder(resp.Body).Decode(&pd)
	if err != nil {
		return nil, err
	}

	return &pd, nil
}
