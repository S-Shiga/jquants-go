//go:build fulltest

package jquants

import (
	"context"
	"testing"
)

func TestClient_IndexPrice(t *testing.T) {
	var indexCode = "0000"
	ctx := context.Background()
	if err := setup(ctx); err != nil {
		t.Fatalf("Failed to setup client: %v", err)
	}
	req := IndexPriceRequest{Code: &indexCode}
	res, err := testClient.IndexPrice(ctx, req)
	if err != nil {
		t.Errorf("Failed to get index price: %s", err)
	}
	if len(res) == 0 {
		t.Error("Empty index price")
	}
}

func TestClient_TopixPrices(t *testing.T) {
	ctx := context.Background()
	if err := setup(ctx); err != nil {
		t.Fatalf("Failed to setup client: %v", err)
	}
	res, err := testClient.TopixPrices(ctx, TopixPriceRequest{})
	if err != nil {
		t.Errorf("Failed to get topix price: %s", err)
	}
	if len(res) == 0 {
		t.Error("Empty topix price")
	}
}
