package jquants

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"strconv"
	"time"
)

type StockTradingValue struct {
	PublishedDate                       time.Time `json:"PublishedDate"`
	StartDate                           time.Time `json:"StartDate"`
	EndDate                             time.Time `json:"EndDate"`
	Section                             string    `json:"Section"`
	ProprietarySales                    int64     `json:"ProprietarySales"`
	ProprietaryPurchases                int64     `json:"ProprietaryPurchases"`
	ProprietaryTotal                    int64     `json:"ProprietaryTotal"`
	ProprietaryBalance                  int64     `json:"ProprietaryBalance"`
	BrokerageSales                      int64     `json:"BrokerageSales"`
	BrokeragePurchases                  int64     `json:"BrokeragePurchases"`
	BrokerageTotal                      int64     `json:"BrokerageTotal"`
	BrokerageBalance                    int64     `json:"BrokerageBalance"`
	NetSales                            int64     `json:"TotalSales"`
	NetPurchases                        int64     `json:"TotalPurchases"`
	NetTotal                            int64     `json:"TotalTotal"`
	NetBalance                          int64     `json:"TotalBalance"`
	IndividualsSales                    int64     `json:"IndividualsSales"`
	IndividualsPurchases                int64     `json:"IndividualsPurchases"`
	IndividualsTotal                    int64     `json:"IndividualsTotal"`
	IndividualsBalance                  int64     `json:"IndividualsBalance"`
	ForeignersSales                     int64     `json:"ForeignersSales"`
	ForeignersPurchases                 int64     `json:"ForeignersPurchases"`
	ForeignersTotal                     int64     `json:"ForeignersTotal"`
	ForeignersBalance                   int64     `json:"ForeignersBalance"`
	SecuritiesCosSales                  int64     `json:"SecuritiesCosSales"`
	SecuritiesCosPurchases              int64     `json:"SecuritiesCosPurchases"`
	SecuritiesCosTotal                  int64     `json:"SecuritiesCosTotal"`
	SecuritiesCosBalance                int64     `json:"SecuritiesCosBalance"`
	InvestmentTrustsSales               int64     `json:"InvestmentTrustsSales"`
	InvestmentTrustsPurchases           int64     `json:"InvestmentTrustsPurchases"`
	InvestmentTrustsTotal               int64     `json:"InvestmentTrustsTotal"`
	InvestmentTrustsBalance             int64     `json:"InvestmentTrustsBalance"`
	BusinessCosSales                    int64     `json:"BusinessCosSales"`
	BusinessCosPurchases                int64     `json:"BusinessCosPurchases"`
	BusinessCosTotal                    int64     `json:"BusinessCosTotal"`
	BusinessCosBalance                  int64     `json:"BusinessCosBalance"`
	OtherCosSales                       int64     `json:"OtherCosSales"`
	OtherCosPurchases                   int64     `json:"OtherCosPurchases"`
	OtherCosTotal                       int64     `json:"OtherCosTotal"`
	OtherCosBalance                     int64     `json:"OtherCosBalance"`
	InsuranceCosSales                   int64     `json:"InsuranceCosSales"`
	InsuranceCosPurchases               int64     `json:"InsuranceCosPurchases"`
	InsuranceCosTotal                   int64     `json:"InsuranceCosTotal"`
	InsuranceCosBalance                 int64     `json:"InsuranceCosBalance"`
	BanksSales                          int64     `json:"CityBKsRegionalBKsEtcSales"`
	BanksPurchases                      int64     `json:"CityBKsRegionalBKsEtcPurchases"`
	BanksTotal                          int64     `json:"CityBKsRegionalBKsEtcTotal"`
	BanksBalance                        int64     `json:"CityBKsRegionalBKsEtcBalance"`
	TrustBanksSales                     int64     `json:"TrustBanksSales"`
	TrustBanksPurchases                 int64     `json:"TrustBanksPurchases"`
	TrustBanksTotal                     int64     `json:"TrustBanksTotal"`
	TrustBanksBalance                   int64     `json:"TrustBanksBalance"`
	OtherFinancialInstitutionsSales     int64     `json:"OtherFinancialInstitutionsSales"`
	OtherFinancialInstitutionsPurchases int64     `json:"OtherFinancialInstitutionsPurchases"`
	OtherFinancialInstitutionsTotal     int64     `json:"OtherFinancialInstitutionsTotal"`
	OtherFinancialInstitutionsBalance   int64     `json:"OtherFinancialInstitutionsBalance"`
}

func (stv *StockTradingValue) UnmarshalJSON(b []byte) error {
	var raw struct {
		PublishedDate                       string  `json:"PublishedDate"`
		StartDate                           string  `json:"StartDate"`
		EndDate                             string  `json:"EndDate"`
		Section                             string  `json:"Section"`
		ProprietarySales                    float64 `json:"ProprietarySales"`
		ProprietaryPurchases                float64 `json:"ProprietaryPurchases"`
		ProprietaryTotal                    float64 `json:"ProprietaryTotal"`
		ProprietaryBalance                  float64 `json:"ProprietaryBalance"`
		BrokerageSales                      float64 `json:"BrokerageSales"`
		BrokeragePurchases                  float64 `json:"BrokeragePurchases"`
		BrokerageTotal                      float64 `json:"BrokerageTotal"`
		BrokerageBalance                    float64 `json:"BrokerageBalance"`
		TotalSales                          float64 `json:"TotalSales"`
		TotalPurchases                      float64 `json:"TotalPurchases"`
		TotalTotal                          float64 `json:"TotalTotal"`
		TotalBalance                        float64 `json:"TotalBalance"`
		IndividualsSales                    float64 `json:"IndividualsSales"`
		IndividualsPurchases                float64 `json:"IndividualsPurchases"`
		IndividualsTotal                    float64 `json:"IndividualsTotal"`
		IndividualsBalance                  float64 `json:"IndividualsBalance"`
		ForeignersSales                     float64 `json:"ForeignersSales"`
		ForeignersPurchases                 float64 `json:"ForeignersPurchases"`
		ForeignersTotal                     float64 `json:"ForeignersTotal"`
		ForeignersBalance                   float64 `json:"ForeignersBalance"`
		SecuritiesCosSales                  float64 `json:"SecuritiesCosSales"`
		SecuritiesCosPurchases              float64 `json:"SecuritiesCosPurchases"`
		SecuritiesCosTotal                  float64 `json:"SecuritiesCosTotal"`
		SecuritiesCosBalance                float64 `json:"SecuritiesCosBalance"`
		InvestmentTrustsSales               float64 `json:"InvestmentTrustsSales"`
		InvestmentTrustsPurchases           float64 `json:"InvestmentTrustsPurchases"`
		InvestmentTrustsTotal               float64 `json:"InvestmentTrustsTotal"`
		InvestmentTrustsBalance             float64 `json:"InvestmentTrustsBalance"`
		BusinessCosSales                    float64 `json:"BusinessCosSales"`
		BusinessCosPurchases                float64 `json:"BusinessCosPurchases"`
		BusinessCosTotal                    float64 `json:"BusinessCosTotal"`
		BusinessCosBalance                  float64 `json:"BusinessCosBalance"`
		OtherCosSales                       float64 `json:"OtherCosSales"`
		OtherCosPurchases                   float64 `json:"OtherCosPurchases"`
		OtherCosTotal                       float64 `json:"OtherCosTotal"`
		OtherCosBalance                     float64 `json:"OtherCosBalance"`
		InsuranceCosSales                   float64 `json:"InsuranceCosSales"`
		InsuranceCosPurchases               float64 `json:"InsuranceCosPurchases"`
		InsuranceCosTotal                   float64 `json:"InsuranceCosTotal"`
		InsuranceCosBalance                 float64 `json:"InsuranceCosBalance"`
		CityBKsRegionalBKsEtcSales          float64 `json:"CityBKsRegionalBKsEtcSales"`
		CityBKsRegionalBKsEtcPurchases      float64 `json:"CityBKsRegionalBKsEtcPurchases"`
		CityBKsRegionalBKsEtcTotal          float64 `json:"CityBKsRegionalBKsEtcTotal"`
		CityBKsRegionalBKsEtcBalance        float64 `json:"CityBKsRegionalBKsEtcBalance"`
		TrustBanksSales                     float64 `json:"TrustBanksSales"`
		TrustBanksPurchases                 float64 `json:"TrustBanksPurchases"`
		TrustBanksTotal                     float64 `json:"TrustBanksTotal"`
		TrustBanksBalance                   float64 `json:"TrustBanksBalance"`
		OtherFinancialInstitutionsSales     float64 `json:"OtherFinancialInstitutionsSales"`
		OtherFinancialInstitutionsPurchases float64 `json:"OtherFinancialInstitutionsPurchases"`
		OtherFinancialInstitutionsTotal     float64 `json:"OtherFinancialInstitutionsTotal"`
		OtherFinancialInstitutionsBalance   float64 `json:"OtherFinancialInstitutionsBalance"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	var err error
	stv.PublishedDate, err = time.Parse(time.DateOnly, raw.PublishedDate)
	if err != nil {
		return err
	}
	stv.StartDate, err = time.Parse(time.DateOnly, raw.StartDate)
	if err != nil {
		return err
	}
	stv.EndDate, err = time.Parse(time.DateOnly, raw.EndDate)
	if err != nil {
		return err
	}
	stv.Section = raw.Section
	stv.ProprietarySales = int64(raw.ProprietarySales)
	stv.ProprietaryPurchases = int64(raw.ProprietaryPurchases)
	stv.ProprietaryTotal = int64(raw.ProprietaryTotal)
	stv.ProprietaryBalance = int64(raw.ProprietaryBalance)
	stv.BrokerageSales = int64(raw.BrokerageSales)
	stv.BrokeragePurchases = int64(raw.BrokeragePurchases)
	stv.BrokerageTotal = int64(raw.BrokerageTotal)
	stv.BrokerageBalance = int64(raw.BrokerageBalance)
	stv.NetSales = int64(raw.TotalSales)
	stv.NetPurchases = int64(raw.TotalPurchases)
	stv.NetTotal = int64(raw.TotalTotal)
	stv.NetBalance = int64(raw.TotalBalance)
	stv.IndividualsSales = int64(raw.IndividualsSales)
	stv.IndividualsPurchases = int64(raw.IndividualsPurchases)
	stv.IndividualsTotal = int64(raw.IndividualsTotal)
	stv.IndividualsBalance = int64(raw.IndividualsBalance)
	stv.ForeignersSales = int64(raw.ForeignersSales)
	stv.ForeignersPurchases = int64(raw.ForeignersPurchases)
	stv.ForeignersTotal = int64(raw.ForeignersTotal)
	stv.ForeignersBalance = int64(raw.ForeignersBalance)
	stv.SecuritiesCosSales = int64(raw.SecuritiesCosSales)
	stv.SecuritiesCosPurchases = int64(raw.SecuritiesCosPurchases)
	stv.SecuritiesCosTotal = int64(raw.SecuritiesCosTotal)
	stv.SecuritiesCosBalance = int64(raw.SecuritiesCosBalance)
	stv.InvestmentTrustsSales = int64(raw.InvestmentTrustsSales)
	stv.InvestmentTrustsPurchases = int64(raw.InvestmentTrustsPurchases)
	stv.InvestmentTrustsTotal = int64(raw.InvestmentTrustsTotal)
	stv.InvestmentTrustsBalance = int64(raw.InvestmentTrustsBalance)
	stv.BusinessCosSales = int64(raw.BusinessCosSales)
	stv.BusinessCosPurchases = int64(raw.BusinessCosPurchases)
	stv.BusinessCosTotal = int64(raw.BusinessCosTotal)
	stv.BusinessCosBalance = int64(raw.BusinessCosBalance)
	stv.OtherCosSales = int64(raw.OtherCosSales)
	stv.OtherCosPurchases = int64(raw.OtherCosPurchases)
	stv.OtherCosTotal = int64(raw.OtherCosTotal)
	stv.OtherCosBalance = int64(raw.OtherCosBalance)
	stv.InsuranceCosSales = int64(raw.InsuranceCosSales)
	stv.InsuranceCosPurchases = int64(raw.InsuranceCosPurchases)
	stv.InsuranceCosTotal = int64(raw.InsuranceCosTotal)
	stv.InsuranceCosBalance = int64(raw.InsuranceCosBalance)
	stv.BanksSales = int64(raw.CityBKsRegionalBKsEtcSales)
	stv.BanksPurchases = int64(raw.CityBKsRegionalBKsEtcPurchases)
	stv.BanksTotal = int64(raw.CityBKsRegionalBKsEtcTotal)
	stv.BanksBalance = int64(raw.CityBKsRegionalBKsEtcBalance)
	stv.TrustBanksSales = int64(raw.TrustBanksSales)
	stv.TrustBanksPurchases = int64(raw.TrustBanksPurchases)
	stv.TrustBanksTotal = int64(raw.TrustBanksTotal)
	stv.TrustBanksBalance = int64(raw.TrustBanksBalance)
	stv.OtherFinancialInstitutionsSales = int64(raw.OtherFinancialInstitutionsSales)
	stv.OtherFinancialInstitutionsPurchases = int64(raw.OtherFinancialInstitutionsPurchases)
	stv.OtherFinancialInstitutionsTotal = int64(raw.OtherFinancialInstitutionsTotal)
	stv.OtherFinancialInstitutionsBalance = int64(raw.OtherFinancialInstitutionsBalance)
	return nil
}

type StockTradingValueRequest struct {
	Section       *string
	From          *time.Time
	To            *time.Time
	PaginationKey *string
}

type stockTradingValueParameter struct {
	Section       *string
	From          *time.Time
	To            *time.Time
	PaginationKey *string
}

func (p stockTradingValueParameter) values() (url.Values, error) {
	v := url.Values{}
	if p.Section != nil {
		v.Add("section", *p.Section)
	}
	if p.From != nil {
		v.Add("from", p.From.Format(time.DateOnly))
	}
	if p.To != nil {
		v.Add("to", p.To.Format(time.DateOnly))
	}
	if p.PaginationKey != nil {
		v.Add("pagination_key", *p.PaginationKey)
	}
	return v, nil
}

type stockTradingValueResponse struct {
	Data          []StockTradingValue `json:"trades_spec"`
	PaginationKey *string             `json:"pagination_key"`
}

func (c *Client) sendStockTradingValueRequest(ctx context.Context, param stockTradingValueParameter) (stockTradingValueResponse, error) {
	var r stockTradingValueResponse
	resp, err := c.sendRequest(ctx, "/markets/trades_spec", param)
	if err != nil {
		return r, fmt.Errorf("failed to send GET request: %w", err)
	}
	if resp.StatusCode != 200 {
		return r, handleErrorResponse(resp)
	}
	if err = decodeResponse(resp, &r); err != nil {
		return r, fmt.Errorf("failed to decode HTTP response: %w", err)
	}
	return r, nil
}

// StockTradingValue provides trading by type of investors.
// https://jpx.gitbook.io/j-quants-en/api-reference/trades_spec
func (c *Client) StockTradingValue(ctx context.Context, req StockTradingValueRequest) ([]StockTradingValue, error) {
	var data []StockTradingValue
	var paginationKey *string
	ctx, cancel := context.WithTimeout(ctx, c.loopTimeout)
	defer cancel()
	for {
		param := stockTradingValueParameter{req.Section, req.From, req.To, paginationKey}
		resp, err := c.sendStockTradingValueRequest(ctx, param)
		if err != nil {
			if errors.As(err, &InternalServerError{}) {
				slog.Warn("Retrying HTTP request", "error", err.Error())
				time.Sleep(c.retryInterval)
				continue
			} else {
				return nil, fmt.Errorf("failed to send stock trading value request: %w", err)
			}
		}
		data = append(data, resp.Data...)
		paginationKey = resp.PaginationKey
		if resp.PaginationKey == nil {
			break
		}
	}
	return data, nil
}

type MarginTradingVolume struct {
	Date                               time.Time `json:"Date"`
	Code                               string    `json:"Code"`
	ShortMarginTradeVolume             int64     `json:"ShortMarginTradeVolume"`
	LongMarginTradeVolume              int64     `json:"LongMarginTradeVolume"`
	ShortNegotiableMarginTradeVolume   int64     `json:"ShortNegotiableMarginTradeVolume"`
	LongNegotiableMarginTradeVolume    int64     `json:"LongNegotiableMarginTradeVolume"`
	ShortStandardizedMarginTradeVolume int64     `json:"ShortStandardizedMarginTradeVolume"`
	LongStandardizedMarginTradeVolume  int64     `json:"LongStandardizedMarginTradeVolume"`
	IssueType                          int8      `json:"IssueType"`
}

func (mtv *MarginTradingVolume) UnmarshalJSON(b []byte) error {
	var raw struct {
		Date                               string  `json:"Date"`
		Code                               string  `json:"Code"`
		ShortMarginTradeVolume             float64 `json:"ShortMarginTradeVolume"`
		LongMarginTradeVolume              float64 `json:"LongMarginTradeVolume"`
		ShortNegotiableMarginTradeVolume   float64 `json:"ShortNegotiableMarginTradeVolume"`
		LongNegotiableMarginTradeVolume    float64 `json:"LongNegotiableMarginTradeVolume"`
		ShortStandardizedMarginTradeVolume float64 `json:"ShortStandardizedMarginTradeVolume"`
		LongStandardizedMarginTradeVolume  float64 `json:"LongStandardizedMarginTradeVolume"`
		IssueType                          string  `json:"IssueType"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return fmt.Errorf("failed to decode margin trade volume error response: %w", err)
	}
	var err error
	mtv.Date, err = time.Parse(time.DateOnly, raw.Date)
	if err != nil {
		return fmt.Errorf("failed to decode margin trade volume error response: %w", err)
	}
	issueType, err := strconv.ParseInt(raw.IssueType, 10, 8)
	if err != nil {
		return fmt.Errorf("failed to decode margin trade volume error response: %w", err)
	}
	mtv.Code = raw.Code
	mtv.ShortMarginTradeVolume = int64(raw.ShortMarginTradeVolume)
	mtv.LongMarginTradeVolume = int64(raw.LongMarginTradeVolume)
	mtv.ShortNegotiableMarginTradeVolume = int64(raw.ShortNegotiableMarginTradeVolume)
	mtv.LongNegotiableMarginTradeVolume = int64(raw.LongNegotiableMarginTradeVolume)
	mtv.ShortStandardizedMarginTradeVolume = int64(raw.ShortStandardizedMarginTradeVolume)
	mtv.LongStandardizedMarginTradeVolume = int64(raw.LongStandardizedMarginTradeVolume)
	mtv.IssueType = int8(issueType)
	return nil
}

type MarginTradingVolumeRequest struct {
	Code *string
	Date *time.Time
	From *time.Time
	To   *time.Time
}

type marginTradingVolumeParameter struct {
	Code          *string
	Date          *time.Time
	From          *time.Time
	To            *time.Time
	PaginationKey *string
}

func (p marginTradingVolumeParameter) values() (url.Values, error) {
	v := url.Values{}
	if p.Date != nil {
		v.Add("date", p.Date.Format(time.DateOnly))
	} else {
		if p.Code == nil {
			return nil, errors.New("code or date is required")
		}
		v.Add("code", *p.Code)
		if p.From != nil {
			v.Add("from", p.From.Format(time.DateOnly))
		}
		if p.To != nil {
			v.Add("to", p.To.Format(time.DateOnly))
		}
	}
	if p.PaginationKey != nil {
		v.Add("pagination_key", *p.PaginationKey)
	}
	return v, nil
}

type marginTradingVolumeResponse struct {
	Data          []MarginTradingVolume `json:"weekly_margin_interest"`
	PaginationKey *string               `json:"pagination_key"`
}

func (c *Client) sendMarginTradingVolumeRequest(ctx context.Context, param marginTradingVolumeParameter) (marginTradingVolumeResponse, error) {
	var r marginTradingVolumeResponse
	resp, err := c.sendRequest(ctx, "/markets/weekly_margin_interest", param)
	if err != nil {
		return r, fmt.Errorf("failed to send GET request: %w", err)
	}
	if resp.StatusCode != 200 {
		return r, handleErrorResponse(resp)
	}
	if err = decodeResponse(resp, &r); err != nil {
		return r, fmt.Errorf("failed to decode HTTP response: %w", err)
	}
	return r, nil
}

// MarginTradingVolume provides margin trading outstandings.
// https://jpx.gitbook.io/j-quants-en/api-reference/weekly_margin_interest
func (c *Client) MarginTradingVolume(ctx context.Context, req MarginTradingVolumeRequest) ([]MarginTradingVolume, error) {
	var data []MarginTradingVolume
	var paginationKey *string
	ctx, cancel := context.WithTimeout(ctx, c.loopTimeout)
	defer cancel()
	for {
		param := marginTradingVolumeParameter{req.Code, req.Date, req.From, req.To, paginationKey}
		resp, err := c.sendMarginTradingVolumeRequest(ctx, param)
		if err != nil {
			if errors.As(err, &InternalServerError{}) {
				slog.Warn("Retrying HTTP request", "error", err.Error())
				time.Sleep(c.retryInterval)
				continue
			} else {
				return nil, fmt.Errorf("failed to send margin trading volume request: %w", err)
			}
		}
		data = append(data, resp.Data...)
		paginationKey = resp.PaginationKey
		if resp.PaginationKey == nil {
			break
		}
	}
	return data, nil
}

// Outstanding Short Selling Positions Reported not implemented

type ShortSellingValue struct {
	Date                                         time.Time `json:"Date"`
	Sector33Code                                 string    `json:"Sector33Code"`
	SellingExcludingShortSellingTurnoverValue    int64     `json:"SellingExcludingShortSellingTurnoverValue"`
	ShortSellingWithRestrictionsTurnoverValue    int64     `json:"ShortSellingWithRestrictionsTurnoverValue"`
	ShortSellingWithoutRestrictionsTurnoverValue int64     `json:"ShortSellingWithoutRestrictionsTurnoverValue"`
}

func (ssv *ShortSellingValue) UnmarshalJSON(b []byte) error {
	var raw struct {
		Date                                         string  `json:"Date"`
		Sector33Code                                 string  `json:"Sector33Code"`
		SellingExcludingShortSellingTurnoverValue    float64 `json:"SellingExcludingShortSellingTurnoverValue"`
		ShortSellingWithRestrictionsTurnoverValue    float64 `json:"ShortSellingWithRestrictionsTurnoverValue"`
		ShortSellingWithoutRestrictionsTurnoverValue float64 `json:"ShortSellingWithoutRestrictionsTurnoverValue"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return fmt.Errorf("failed to decode short selling value error response: %w", err)
	}
	t, err := time.Parse(time.DateOnly, raw.Date)
	if err != nil {
		return fmt.Errorf("failed to decode short selling value error response: %w", err)
	}
	ssv.Date = t
	ssv.Sector33Code = raw.Sector33Code
	ssv.SellingExcludingShortSellingTurnoverValue = int64(raw.SellingExcludingShortSellingTurnoverValue)
	ssv.ShortSellingWithRestrictionsTurnoverValue = int64(raw.ShortSellingWithRestrictionsTurnoverValue)
	ssv.ShortSellingWithoutRestrictionsTurnoverValue = int64(raw.ShortSellingWithoutRestrictionsTurnoverValue)
	return nil
}

type ShortSellingValueRequest struct {
	Sector33Code *string
	Date         *time.Time
	From         *time.Time
	To           *time.Time
}

type shortSellingValueParameter struct {
	Sector33Code  *string
	Date          *time.Time
	From          *time.Time
	To            *time.Time
	PaginationKey *string
}

func (p shortSellingValueParameter) values() (url.Values, error) {
	v := url.Values{}
	if p.Sector33Code != nil {
		v.Add("sector33code", *p.Sector33Code)
		if p.Date != nil {
			v.Add("date", p.Date.Format(time.DateOnly))
		} else {
			if p.From != nil {
				v.Add("from", p.From.Format(time.DateOnly))
			}
			if p.To != nil {
				v.Add("to", p.To.Format(time.DateOnly))
			}
		}
	} else {
		if p.Date == nil {
			return nil, errors.New("sector33code or date is required")
		}
		v.Add("date", p.Date.Format(time.DateOnly))
	}
	if p.PaginationKey != nil {
		v.Add("pagination_key", *p.PaginationKey)
	}
	return v, nil
}

type shortSellingValueResponse struct {
	Data          []ShortSellingValue `json:"short_selling"`
	PaginationKey *string             `json:"pagination_key"`
}

func (c *Client) sendShortSellingValueRequest(ctx context.Context, req shortSellingValueParameter) (shortSellingValueResponse, error) {
	var r shortSellingValueResponse
	resp, err := c.sendRequest(ctx, "/markets/short_selling", req)
	if err != nil {
		return r, fmt.Errorf("failed to send GET request: %w", err)
	}
	if resp.StatusCode != 200 {
		return r, handleErrorResponse(resp)
	}
	if err = decodeResponse(resp, &r); err != nil {
		return r, fmt.Errorf("failed to decode HTTP response: %w", err)
	}
	return r, nil
}

func (c *Client) ShortSellingValue(ctx context.Context, req ShortSellingValueRequest) ([]ShortSellingValue, error) {
	var data = make([]ShortSellingValue, 0)
	var paginationKey *string
	ctx, cancel := context.WithTimeout(ctx, c.loopTimeout)
	defer cancel()
	for {
		param := shortSellingValueParameter{req.Sector33Code, req.Date, req.From, req.To, paginationKey}
		resp, err := c.sendShortSellingValueRequest(ctx, param)
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
		if resp.PaginationKey == nil {
			break
		}
	}
	return data, nil
}
