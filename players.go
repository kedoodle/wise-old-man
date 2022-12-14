package wiseoldman

import (
	"context"
	"encoding/json"
	"net/http"
)

func (c *Client) UpdatePlayer(ctx context.Context, username string) (*PlayerDetails, error) {
	var pd PlayerDetails

	url := c.baseURL + "/players/" + username
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&pd)
	if err != nil {
		return nil, err
	}

	return &pd, nil
}
