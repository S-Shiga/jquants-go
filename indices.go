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

type IndexPrice struct {
	Date  string      `json:"Date"`
	Code  string      `json:"Code"`
	Open  json.Number `json:"Open"`
	High  json.Number `json:"High"`
	Low   json.Number `json:"Low"`
	Close json.Number `json:"Close"`
}

func (ip *IndexPrice) UnmarshalJSON(b []byte) error {
	var raw struct {
		Date  string      `json:"Date"`
		Code  string      `json:"Code"`
		Open  json.Number `json:"Open"`
		High  json.Number `json:"High"`
		Low   json.Number `json:"Low"`
		Close json.Number `json:"Close"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return fmt.Errorf("failed to decode index price error response: %w", err)
	}
	ip.Date = raw.Date
	ip.Code = raw.Code
	ip.Open = raw.Open
	ip.High = raw.High
	ip.Low = raw.Low
	ip.Close = raw.Close
	return nil
}

type IndexPriceRequest struct {
	Code *string
	Date *string
	From *string
	To   *string
}

type indexPriceParameter struct {
	Code          *string
	Date          *string
	From          *string
	To            *string
	PaginationKey *string
}

func (p indexPriceParameter) values() (url.Values, error) {
	v := url.Values{}
	if p.Date != nil {
		v.Add("date", *p.Date)
	} else {
		if p.Code == nil {
			return nil, errors.New("code or date is required")
		}
		v.Add("code", *p.Code)
		if p.From != nil {
			v.Add("from", *p.From)
		}
		if p.To != nil {
			v.Add("to", *p.To)
		}
	}
	if p.PaginationKey != nil {
		v.Add("pagination_key", *p.PaginationKey)
	}
	return v, nil
}

type indexPriceResponse struct {
	Data          []IndexPrice `json:"indices"`
	PaginationKey *string      `json:"pagination_key"`
}

func (c *Client) sendIndexPriceRequest(ctx context.Context, param indexPriceParameter) (indexPriceResponse, error) {
	var r indexPriceResponse
	resp, err := c.sendRequest(ctx, "/indices", param)
	if err != nil {
		return r, fmt.Errorf("failed to send GET request: %w", err)
	}
	if resp.StatusCode != 200 {
		return r, handleErrorResponse(resp)
	}
	if err = decodeResponse(resp, &r); err != nil {
		return r, fmt.Errorf("failed to decode HTTP reaponse: %w", err)
	}
	return r, nil
}

func (c *Client) IndexPrice(ctx context.Context, req IndexPriceRequest) ([]IndexPrice, error) {
	var data = make([]IndexPrice, 0)
	var paginationKey *string
	ctx, cancel := context.WithTimeout(ctx, c.loopTimeout)
	defer cancel()
	for {
		param := indexPriceParameter{req.Code, req.Date, req.From, req.To, paginationKey}
		resp, err := c.sendIndexPriceRequest(ctx, param)
		if err != nil {
			if errors.As(err, &InternalServerError{}) {
				slog.Warn("Retrying HTTP request", "error", err.Error())
				time.Sleep(c.retryInterval)
				continue
			} else {
				return nil, err
			}
		}
		data = append(data, resp.Data...)
		paginationKey = resp.PaginationKey
		if paginationKey == nil {
			break
		}
	}
	return data, nil
}

type TopixPrice struct {
	Date  string      `json:"Date"`
	Open  json.Number `json:"Open"`
	High  json.Number `json:"High"`
	Low   json.Number `json:"Low"`
	Close json.Number `json:"Close"`
}

func (p *TopixPrice) UnmarshalJSON(b []byte) error {
	var raw struct {
		Date  string      `json:"Date"`
		Open  json.Number `json:"Open"`
		High  json.Number `json:"High"`
		Low   json.Number `json:"Low"`
		Close json.Number `json:"Close"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return fmt.Errorf("failed to unmarshal topix price: %w", err)
	}
	p.Date = raw.Date
	p.Open = raw.Open
	p.High = raw.High
	p.Low = raw.Low
	p.Close = raw.Close
	return nil
}

type TopixPriceRequest struct {
	From *string
	To   *string
}

type topixPriceParameter struct {
	From          *string
	To            *string
	PaginationKey *string
}

func (p topixPriceParameter) values() (url.Values, error) {
	v := url.Values{}
	if p.From != nil {
		v.Add("from", *p.From)
	}
	if p.To != nil {
		v.Add("to", *p.To)
	}
	if p.PaginationKey != nil {
		v.Add("pagination_key", *p.PaginationKey)
	}
	return v, nil
}

type topixPriceResponse struct {
	Topix         []TopixPrice `json:"topix"`
	PaginationKey *string      `json:"pagination_key"`
}

func (c *Client) sendTopixPriceRequest(ctx context.Context, param topixPriceParameter) (topixPriceResponse, error) {
	var r topixPriceResponse
	resp, err := c.sendRequest(ctx, "/indices/topix", param)
	if err != nil {
		return r, fmt.Errorf("failed to send GET request: %w", err)
	}
	if resp.StatusCode != 200 {
		return r, handleErrorResponse(resp)
	}
	if err = decodeResponse(resp, &r); err != nil {
		return r, fmt.Errorf("failed to decode HTTP reaponse: %w", err)
	}
	return r, nil
}

func (c *Client) TopixPrices(ctx context.Context, req TopixPriceRequest) ([]TopixPrice, error) {
	var data = make([]TopixPrice, 0)
	var paginationKey *string
	ctx, cancel := context.WithTimeout(ctx, c.loopTimeout)
	defer cancel()
	for {
		param := topixPriceParameter{req.From, req.To, paginationKey}
		resp, err := c.sendTopixPriceRequest(ctx, param)
		if err != nil {
			if errors.As(err, &InternalServerError{}) {
				slog.Warn("Retrying HTTP request", "error", err.Error())
				time.Sleep(c.retryInterval)
				continue
			} else {
				return nil, fmt.Errorf("failed to send topix price request: %w", err)
			}
		}
		data = append(data, resp.Topix...)
		paginationKey = resp.PaginationKey
		if paginationKey == nil {
			break
		}
	}
	return data, nil
}
