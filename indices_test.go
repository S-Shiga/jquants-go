package jquants

import (
	"context"
	"net/http"
	"testing"
)

func TestClient_IndexPrice(t *testing.T) {
	var indexCode = "0000"
	ctx := context.Background()
	httpClient := &http.Client{}
	client, err := NewClient(ctx, httpClient)
	if err != nil {
		t.Fatal(err)
	}
	req := IndexPriceRequest{Code: &indexCode}
	res, err := client.IndexPrice(ctx, req)
	if err != nil {
		t.Errorf("Failed to get index price: %s", err)
	}
	if len(res) == 0 {
		t.Error("Empty index price")
	}
}

func TestClient_TopixPrices(t *testing.T) {
	ctx := context.Background()
	httpClient := &http.Client{}
	client, err := NewClient(ctx, httpClient)
	if err != nil {
		t.Fatal(err)
	}
	res, err := client.TopixPrices(ctx, TopixPriceRequest{})
	if err != nil {
		t.Errorf("Failed to get topix price: %s", err)
	}
	if len(res) == 0 {
		t.Error("Empty topix price")
	}
}
