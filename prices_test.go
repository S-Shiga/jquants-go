//go:build fulltest

package jquants

import (
	"context"
	"testing"
)

func TestClient_StockPrice(t *testing.T) {
	var code = "13010"
	ctx := context.Background()
	if err := setup(ctx); err != nil {
		t.Fatalf("Failed to setup client: %v", err)
	}
	req := StockPriceRequest{Code: &code}
	res, err := testClient.StockPrice(ctx, req)
	if err != nil {
		t.Errorf("Failed to get stock price: %s", err)
	}
	if len(res) == 0 {
		t.Error("Empty stock price")
	}
}

func TestClient_StockPriceWithChannel(t *testing.T) {
	var code = "13010"
	ctx := context.Background()
	if err := setup(ctx); err != nil {
		t.Fatalf("Failed to setup client: %v", err)
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	req := StockPriceRequest{Code: &code}
	ch := make(chan StockPrice)
	go func() {
		if e := testClient.StockPriceWithChannel(ctx, req, ch); e != nil {
			t.Errorf("Failed to get stock price: %s", e)
		}
	}()
	found := false
	for range ch {
		found = true
	}
	if !found {
		t.Error("Empty stock price")
	}
}
