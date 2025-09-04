//go:build fulltest

package jquants

import (
	"context"
	"testing"

	"github.com/S-Shiga/jquants-go/codes"
)

func TestClient_StockTradingValue(t *testing.T) {
	var code = codes.SectionPrime
	ctx := context.Background()
	if err := setup(ctx); err != nil {
		t.Fatalf("Failed to setup client: %v", err)
	}
	req := StockTradingValueRequest{Section: &code}
	res, err := testClient.StockTradingValue(ctx, req)
	if err != nil {
		t.Errorf("Failed to get stock trading value: %s", err)
	}
	if len(res) == 0 {
		t.Error("Empty stock trading value")
	}
}

func TestClient_MarginTradingVolume(t *testing.T) {
	var code = "13010"
	ctx := context.Background()
	if err := setup(ctx); err != nil {
		t.Fatalf("Failed to setup client: %v", err)
	}
	req := MarginTradingVolumeRequest{Code: &code}
	res, err := testClient.MarginTradingVolume(ctx, req)
	if err != nil {
		t.Errorf("Failed to get margin trading volume: %s", err)
	}
	if len(res) == 0 {
		t.Error("Empty margin trading volume")
	}
}

func TestClient_ShortSellingValue(t *testing.T) {
	var sector33Code = codes.Sector33FisheryAgricultureAndForestry
	ctx := context.Background()
	if err := setup(ctx); err != nil {
		t.Fatalf("Failed to setup client: %v", err)
	}
	req := ShortSellingValueRequest{Sector33Code: &sector33Code}
	res, err := testClient.ShortSellingValue(ctx, req)
	if err != nil {
		t.Errorf("Failed to get short selling value: %s", err)
	}
	if len(res) == 0 {
		t.Errorf("Empty short selling value")
	}
}

func TestClient_TradingCalendar(t *testing.T) {
	ctx := context.Background()
	if err := setup(ctx); err != nil {
		t.Fatalf("Failed to setup client: %v", err)
	}
	res, err := testClient.TradingCalendar(ctx, TradingCalendarRequest{})
	if err != nil {
		t.Errorf("Failed to get trading calendar: %s", err)
	}
	if len(res) == 0 {
		t.Errorf("Empty trading calendar")
	}
}
