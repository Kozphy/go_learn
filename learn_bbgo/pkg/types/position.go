package types

import "github.com/zixsa/learn_bbgo/pkg/fixedpoint"

type ExchangeFee struct {
	MakerFeeRate fixedpoint.Value
	TakerFeeRate fixedpoint.Value
}

type PositionRisk struct {
	Leverage         fixedpoint.Value `json:"leverage"`
	LiquidationPrice fixedpoint.Value `json:"liquidationPrice"`
}

type FuturesPosition struct {
	Symbol        string `json:"symbol"`
	BaseCurrency  string `json:"baseCurrency"`
	QuoteCurrency string `json:"quoteCurrency"`

	Market Market `json:"market"`

	Base        fixedpoint.Value `json:"base"`
	Quote       fixedpoint.Value `json:"quote"`
	AverageCost fixedpoint.Value `json:"averageCost"`

	// ApproximateAverageCost adds the computed fee in quote in the average cost
	// This is used for calculating net profit
	ApproximateAverageCost fixedpoint.Value `json:"approximateAverageCost"`

	FeeRate          *ExchangeFee                 `json:"feeRate,omitempty"`
	ExchangeFeeRates map[ExchangeName]ExchangeFee `json:"exchangeFeeRates"`

	// Futures data fields
	Isolated     bool  `json:"isolated"`
	UpdateTime   int64 `json:"updateTime"`
	PositionRisk *PositionRisk
}
