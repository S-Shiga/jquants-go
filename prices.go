package jquants

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"time"
)

type StockPrice struct {
	Date             time.Time    `json:"Date"`
	Code             string       `json:"Code"`
	Open             *json.Number `json:"Open"`
	High             *json.Number `json:"High"`
	Low              *json.Number `json:"Low"`
	Close            *json.Number `json:"Close"`
	UpperLimit       bool         `json:"UpperLimit"`
	LowerLimit       bool         `json:"LowerLimit"`
	Volume           *int64       `json:"Volume"`
	TurnoverValue    *int64       `json:"TurnoverValue"`
	AdjustmentFactor json.Number  `json:"AdjustmentFactor"`
}

func (sp *StockPrice) UnmarshalJSON(b []byte) error {
	var raw struct {
		Date             string       `json:"Date"`
		Code             string       `json:"Code"`
		Open             *json.Number `json:"Open"`
		High             *json.Number `json:"High"`
		Low              *json.Number `json:"Low"`
		Close            *json.Number `json:"Close"`
		UpperLimit       string       `json:"UpperLimit"`
		LowerLimit       string       `json:"LowerLimit"`
		Volume           *float64     `json:"Volume"`
		TurnoverValue    *float64     `json:"TurnoverValue"`
		AdjustmentFactor json.Number  `json:"AdjustmentFactor"`
	}
	var volume, turnoverValue *int64
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	t, err := time.Parse(time.DateOnly, raw.Date)
	if err != nil {
		return err
	}
	upperLimit, err := unmarshalLimit(raw.UpperLimit)
	if err != nil {
		return err
	}
	lowerLimit, err := unmarshalLimit(raw.LowerLimit)
	if err != nil {
		return err
	}
	if raw.Volume != nil {
		v := int64(*raw.Volume)
		volume = &v
	}
	if raw.TurnoverValue != nil {
		v := int64(*raw.TurnoverValue)
		turnoverValue = &v
	}
	sp.Date = t
	sp.Code = raw.Code
	sp.Open = raw.Open
	sp.High = raw.High
	sp.Low = raw.Low
	sp.Close = raw.Close
	sp.UpperLimit = upperLimit
	sp.LowerLimit = lowerLimit
	sp.Volume = volume
	sp.TurnoverValue = turnoverValue
	sp.AdjustmentFactor = raw.AdjustmentFactor
	return nil
}

func unmarshalLimit(s string) (bool, error) {
	switch s {
	case "0":
		return false, nil
	case "1":
		return true, nil
	default:
		return false, fmt.Errorf("unknown value: %s", s)
	}
}

type StockPriceRequest struct {
	Code          *string
	Date          *time.Time
	From          *time.Time
	To            *time.Time
	PaginationKey *string
}

func (r *StockPriceRequest) values() (url.Values, error) {
	v := url.Values{}
	if r.Date != nil {
		v.Add("date", r.Date.Format(time.DateOnly))
	} else {
		if r.Code == nil {
			return nil, fmt.Errorf("code or date is required")
		}
		v.Add("code", *r.Code)
		if r.From != nil {
			v.Add("from", r.From.Format(time.DateOnly))
		}
		if r.To != nil {
			v.Add("to", r.To.Format(time.DateOnly))
		}
	}
	if r.PaginationKey != nil {
		v.Add("pagination_key", *r.PaginationKey)
	}
	return v, nil
}

type stockPriceResponse struct {
	Data          []StockPrice `json:"daily_quotes"`
	PaginationKey *string      `json:"pagination_key"`
}

func (c *Client) sendStockPriceRequest(ctx context.Context, req StockPriceRequest) (*stockPriceResponse, error) {
	var r stockPriceResponse

	u, err := url.Parse(c.baseURL + "/prices/daily_quotes")
	if err != nil {
		panic(err)
	}
	v, err := req.values()
	if err != nil {
		panic(err)
	}
	u.RawQuery = v.Encode()
	resp, err := c.sendGetRequest(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("failed to send GET request: %w", err)
	}

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}
	if err = decodeResponse(resp, &r); err != nil {
		return nil, fmt.Errorf("failed to decode HTTP response: %w", err)
	}

	return &r, nil
}

func (c *Client) StockPrice(ctx context.Context, req StockPriceRequest) ([]StockPrice, error) {
	var sps = make([]StockPrice, 0)
	var paginationKey *string

	ctx, cancel := context.WithTimeout(ctx, c.loopTimeout)
	defer cancel()

	for {
		subReq := StockPriceRequest{req.Code, req.Date, req.From, req.To, paginationKey}
		resp, err := c.sendStockPriceRequest(ctx, subReq)
		if err != nil {
			if errors.As(err, &InternalServerError{}) {
				slog.Warn("Retrying HTTP request", "error", err.Error())
				time.Sleep(c.retryInterval)
				continue
			} else {
				return nil, fmt.Errorf("failed to send stock price request: %w", err)
			}
		}
		sps = append(sps, resp.Data...)
		paginationKey = resp.PaginationKey
		if resp.PaginationKey == nil {
			break
		}
	}
	return sps, nil
}

func (c *Client) StockPriceWithChannel(ctx context.Context, req StockPriceRequest, ch chan<- StockPrice) error {
	var paginationKey *string

	ctx, cancel := context.WithTimeout(ctx, c.loopTimeout)
	defer cancel()

	for {
		subReq := StockPriceRequest{req.Code, req.Date, req.From, req.To, paginationKey}
		resp, err := c.sendStockPriceRequest(ctx, subReq)
		if err != nil {
			if errors.As(err, &InternalServerError{}) {
				slog.Warn("Retrying HTTP request", "error", err.Error())
				time.Sleep(c.retryInterval)
				continue
			} else {
				return fmt.Errorf("failed to send stock price request: %w", err)
			}
		}
		for _, d := range resp.Data {
			ch <- d
		}
		paginationKey = resp.PaginationKey
		if resp.PaginationKey == nil {
			break
		}
	}
	close(ch)
	return nil
}

// Morning Session Stock Prices not implemented
