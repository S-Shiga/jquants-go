package jquants

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const baseURL = "https://api.jquants.com/v1"

type Client struct {
	HttpClient    *http.Client
	BaseURL       string
	MailAddress   string
	Password      string
	refreshToken  string
	idToken       string
	retryInterval time.Duration
	loopTimeout   time.Duration
}

func NewClient(ctx context.Context, httpClient *http.Client) (*Client, error) {
	var err error
	email, ok := os.LookupEnv("J_QUANTS_EMAIL_ADDRESS")
	if !ok {
		return nil, errors.New("J_QUANTS_EMAIL_ADDRESS not set")
	}
	password, ok := os.LookupEnv("J_QUANTS_PASSWORD")
	if !ok {
		return nil, errors.New("J_QUANTS_PASSWORD not set")
	}
	client := &Client{
		HttpClient:    httpClient,
		BaseURL:       baseURL,
		MailAddress:   email,
		Password:      password,
		retryInterval: 5 * time.Second,
		loopTimeout:   20 * time.Second,
	}
	refreshToken := os.Getenv("J_QUANTS_REFRESH_TOKEN")
	if refreshToken == "" {
		if err = client.resetRefreshToken(ctx); err != nil {
			return nil, err
		}
	}
	if err = client.resetIDToken(ctx); err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) sendPostRequest(ctx context.Context, u *url.URL, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) sendGetRequest(ctx context.Context, u *url.URL) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	if c.idToken == "" {
		panic("idToken is empty")
	}
	req.Header.Set("Authorization", "Bearer "+c.idToken)
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type parameters interface {
	values() (url.Values, error)
}

func (c *Client) sendRequest(ctx context.Context, urlPath string, param parameters) (*http.Response, error) {
	u, err := url.Parse(c.BaseURL + urlPath)
	if err != nil {
		panic(err)
	}
	v, err := param.values()
	if err != nil {
		panic(err)
	}
	u.RawQuery = v.Encode()
	return c.sendGetRequest(ctx, u)
}

type BadRequest struct {
	err error
}

func (e BadRequest) Error() string {
	return fmt.Sprintf("400 bad request: %v", e.err)
}

func (e BadRequest) Unwrap() error {
	return e.err
}

type Unauthorized struct {
	err error
}

func (e Unauthorized) Error() string {
	return fmt.Sprintf("401 unauthorized: %v", e.err)
}

func (e Unauthorized) Unwrap() error {
	return e.err
}

type Forbidden struct {
	err error
}

func (e Forbidden) Error() string {
	return fmt.Sprintf("403 forbidden: %v", e.err)
}

func (e Forbidden) Unwrap() error {
	return e.err
}

type PayloadTooLarge struct {
	err error
}

func (e PayloadTooLarge) Error() string {
	return fmt.Sprintf("413 payload too large: %v", e.err)
}

func (e PayloadTooLarge) Unwrap() error {
	return e.err
}

type InternalServerError struct {
	err error
}

func (e InternalServerError) Error() string {
	return fmt.Sprintf("500 internal server error: %v", e.err)
}

func (e InternalServerError) Unwrap() error {
	return e.err
}

func decodeResponse(resp *http.Response, body any) error {
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			log.Printf("failed to close response body: %s", closeErr.Error())
		}
	}()
	if err := json.NewDecoder(resp.Body).Decode(body); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	return nil
}

type ErrResponse struct {
	Message string `json:"message"`
}

func handleErrorResponse(resp *http.Response) error {
	if resp.StatusCode == 400 {
		return BadRequest{decodeErrorResponse(resp)}
	} else if resp.StatusCode == 401 {
		return Unauthorized{decodeErrorResponse(resp)}
	} else if resp.StatusCode == 403 {
		return Forbidden{decodeErrorResponse(resp)}
	} else if resp.StatusCode == 413 {
		return PayloadTooLarge{decodeErrorResponse(resp)}
	} else if resp.StatusCode == 500 {
		return InternalServerError{decodeErrorResponse(resp)}
	} else {
		return decodeErrorResponse(resp)
	}
}

func decodeErrorResponse(resp *http.Response) error {
	var errResp ErrResponse
	if err := decodeResponse(resp, &errResp); err != nil {
		return fmt.Errorf("failed to decode error response: %w", err)
	}
	return errors.New(errResp.Message)
}
