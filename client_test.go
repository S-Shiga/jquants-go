package jquants

import (
	"context"
	"net/http"
	"testing"
	"time"
)

var testClient *Client

func setup(ctx context.Context) error {
	if testClient != nil {
		return nil
	}
	var err error
	httpClient := &http.Client{Timeout: time.Second * 5}
	testClient, err = NewClient(ctx, httpClient)
	return err
}

func TestNewClient(t *testing.T) {
	ctx := context.Background()
	if err := setup(ctx); err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	if testClient.IDToken == "" {
		t.Error("Empty ID Token")
	}
}
