//go:build fulltest

package jquants

import (
	"context"
	"net/http"
	"testing"
)

func TestClient_StockTradingValue(t *testing.T) {
	var section = "TSEPrime"
	ctx := context.Background()
	httpClient := &http.Client{}
	client, err := NewClient(ctx, httpClient)
	if err != nil {
		t.Fatal(err)
	}
	req := StockTradingValueRequest{Section: &section}
	res, err := client.StockTradingValue(ctx, req)
	if err != nil {
		t.Errorf("Failed to get stock trading value: %s", err)
	}
	if len(res) == 0 {
		t.Errorf("Empty stock trading value")
	}
}

func TestClient_MarginTradingVolume(t *testing.T) {
	var code = "13010"
	ctx := context.Background()
	httpClient := &http.Client{}
	client, err := NewClient(ctx, httpClient)
	if err != nil {
		t.Fatal(err)
	}
	req := MarginTradingVolumeRequest{Code: &code}
	res, err := client.MarginTradingVolume(ctx, req)
	if err != nil {
		t.Errorf("Failed to get margin trading volume: %s", err)
	}
	if len(res) == 0 {
		t.Errorf("Empty margin trading volume")
	}
}

func TestClient_ShortSellingValue(t *testing.T) {
	var sector33Code = "0050"
	ctx := context.Background()
	httpClient := &http.Client{}
	client, err := NewClient(ctx, httpClient)
	if err != nil {
		t.Fatal(err)
	}
	req := ShortSellingValueRequest{Sector33Code: &sector33Code}
	res, err := client.ShortSellingValue(ctx, req)
	if err != nil {
		t.Errorf("Failed to get short selling value: %s", err)
	}
	if len(res) == 0 {
		t.Errorf("Empty short selling value")
	}
}
