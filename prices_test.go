//go:build fulltest

package jquants

import (
	"context"
	"net/http"
	"testing"
)

var code = "13010"

func TestClient_StockPrice(t *testing.T) {
	ctx := context.Background()
	httpClient := &http.Client{}
	client, err := NewClient(ctx, httpClient)
	if err != nil {
		t.Fatal(err)
	}

	req := StockPriceRequest{Code: &code}
	res, err := client.StockPrice(ctx, req)
	if err != nil {
		t.Errorf("Failed to get stock price: %s", err)
	}
	if len(res) == 0 {
		t.Error("Empty stock price")
	}
}

func TestClient_StockPriceWithChannel(t *testing.T) {
	ctx := context.Background()
	httpClient := &http.Client{}
	client, err := NewClient(ctx, httpClient)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	req := StockPriceRequest{Code: &code}
	ch := make(chan StockPrice)
	go func() {
		if e := client.StockPriceWithChannel(ctx, req, ch); e != nil {
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
