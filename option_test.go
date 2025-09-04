//go:build fulltest

package jquants

import (
	"context"
	"testing"
)

func TestClient_IndexOptionPrice(t *testing.T) {
	date := "2025-01-06"
	ctx := context.Background()
	if err := setup(ctx); err != nil {
		t.Fatalf("Failed to setup client: %v", err)
	}
	req := IndexOptionPriceRequest{Date: date}
	resp, err := testClient.IndexOptionPrice(ctx, req)
	if err != nil {
		t.Errorf("Failed to get index option price: %v", err)
	}
	if len(resp) == 0 {
		t.Error("Empty response")
	}
}

func TestClient_IndexOptionPriceWithChannel(t *testing.T) {
	date := "2025-01-06"
	ctx := context.Background()
	if err := setup(ctx); err != nil {
		t.Fatalf("Failed to setup client: %v", err)
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	req := IndexOptionPriceRequest{Date: date}
	ch := make(chan IndexOptionPrice)
	go func() {
		if e := testClient.IndexOptionPriceWithChannel(ctx, req, ch); e != nil {
			t.Errorf("Failed to get index option price: %v", e)
		}
	}()
	found := false
	for range ch {
		found = true
	}
	if !found {
		t.Error("Empty response")
	}
}
