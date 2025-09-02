package jquants

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	ctx := context.Background()
	httpClient := &http.Client{Timeout: time.Second}
	client, err := NewClient(ctx, httpClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	if client.idToken == "" {
		t.Error("Empty ID Token")
	}
}
