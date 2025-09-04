package jquants

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type IssueInformation struct {
	Date               string `json:"Date"`
	Code               string `json:"Code"`
	CompanyName        string `json:"CompanyName"`
	CompanyNameEnglish string `json:"CompanyNameEnglish"`
	Sector17Code       int8   `json:"Sector17Code"`
	Sector33Code       string `json:"Sector33Code"`
	ScaleCategory      string `json:"ScaleCategory"`
	MarketCode         string `json:"MarketCode"`
	MarginCode         *int8  `json:"MarginCode"`
}

func (ii *IssueInformation) UnmarshalJSON(b []byte) error {
	var raw struct {
		Date               string  `json:"Date"`
		Code               string  `json:"Code"`
		CompanyName        string  `json:"CompanyName"`
		CompanyNameEnglish string  `json:"CompanyNameEnglish"`
		Sector17Code       string  `json:"Sector17Code"`
		Sector17CodeName   string  `json:"Sector17CodeName"`
		Sector33Code       string  `json:"Sector33Code"`
		Sector33CodeName   string  `json:"Sector33CodeName"`
		ScaleCategory      string  `json:"ScaleCategory"`
		MarketCode         string  `json:"MarketCode"`
		MarketCodeName     string  `json:"MarketCodeName"`
		MarginCode         *string `json:"MarginCode"`
		MarginCodeName     *string `json:"MarginCodeName"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	ii.Date = raw.Date
	ii.Code = raw.Code
	ii.CompanyName = raw.CompanyName
	ii.CompanyNameEnglish = raw.CompanyNameEnglish
	sector17Code, err := strconv.ParseInt(raw.Sector17Code, 10, 8)
	if err != nil {
		return err
	}
	ii.Sector17Code = int8(sector17Code)
	ii.Sector33Code = raw.Sector33Code
	ii.ScaleCategory = raw.ScaleCategory
	ii.MarketCode = raw.MarketCode
	if raw.MarginCode != nil {
		marginCode, err := strconv.ParseInt(*raw.MarginCode, 10, 8)
		if err != nil {
			return err
		}
		v := int8(marginCode)
		ii.MarginCode = &v
	}
	return nil
}

type IssueInformationRequest struct {
	Code *string
	Date *string
}

type issueInformationParameters struct {
	IssueInformationRequest
}

func (p issueInformationParameters) values() (url.Values, error) {
	v := url.Values{}
	if p.Code != nil {
		v.Add("code", *p.Code)
	}
	if p.Date != nil {
		v.Add("date", *p.Date)
	}
	return v, nil
}

type issueInformationResponse struct {
	Information []IssueInformation `json:"info"`
}

func (c *Client) IssueInformation(ctx context.Context, req IssueInformationRequest) ([]IssueInformation, error) {
	var r issueInformationResponse
	params := issueInformationParameters{req}
	resp, err := c.sendRequest(ctx, "/listed/info", params)
	if err != nil {
		return nil, fmt.Errorf("failed to send GET request: %w", err)
	}

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}
	if err = decodeResponse(resp, &r); err != nil {
		return nil, fmt.Errorf("failed to decode HTTP response: %w", err)
	}
	return r.Information, nil
}
