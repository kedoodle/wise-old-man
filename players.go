package wiseoldman

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func trimUsername(username string) string {
	return strings.Trim(username, " _-")
}

func validateUsername(username string) error {
	username = trimUsername(username)

	if username == "" {
		return errors.New("validation error: username must be defined")
	}

	if len(username) > 12 {
		return errors.New("validation error: username must be between 1 and 12 characters long")
	}

	matched, err := regexp.MatchString(`^[[:alnum:] _-]+$`, username)
	if err != nil {
		return err
	}
	hasSpecialCharacter := !matched
	if hasSpecialCharacter {
		return errors.New("validation error: username cannot contain any special characters")
	}

	return nil
}

func (c *Client) UpdatePlayer(ctx context.Context, username string) (*PlayerDetails, error) {
	err := validateUsername(username)
	if err != nil {
		return nil, err
	}

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
