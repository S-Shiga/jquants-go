package jquants

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

func (c *Client) fetchRefreshToken(ctx context.Context) (string, error) {
	u, err := url.Parse(c.baseURL + "/token/auth_user")
	if err != nil {
		panic(err)
	}

	body := map[string]string{"mailaddress": c.mailAddress, "password": c.password}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	resp, err := c.sendPostRequest(ctx, u, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to send POST request: %w", err)
	}

	if resp.StatusCode != 200 {
		return "", handleErrorResponse(resp)
	}
	var response struct {
		RefreshToken string `json:"refreshToken"`
	}
	if err = decodeResponse(resp, &response); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}
	return response.RefreshToken, nil
}

func (c *Client) resetRefreshToken(ctx context.Context) error {
	refreshToken, err := c.fetchRefreshToken(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch refresh token: %w", err)
	}
	c.refreshToken = refreshToken
	return nil
}

func (c *Client) fetchIDToken(ctx context.Context) (string, error) {
	u, err := url.Parse(c.baseURL + "/token/auth_refresh")
	if err != nil {
		panic(err)
	}
	if c.refreshToken == "" {
		return "", errors.New("refresh token is empty")
	}
	v := url.Values{"refreshtoken": {c.refreshToken}}
	u.RawQuery = v.Encode()

	resp, err := c.sendPostRequest(ctx, u, nil)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", handleErrorResponse(resp)
	}
	var response struct {
		IDToken string `json:"idToken"`
	}
	if err = decodeResponse(resp, &response); err != nil {
		return "", fmt.Errorf("failed to decode HTTP response: %w", err)
	}
	return response.IDToken, nil
}

func (c *Client) resetIDToken(ctx context.Context) error {
	idToken, err := c.fetchIDToken(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch ID token: %w", err)
	}
	c.idToken = idToken
	return nil
}
