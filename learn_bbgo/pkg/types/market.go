package types

import "github.com/zixsa/learn_bbgo/pkg/fixedpoint"

type Market struct {
	Symbol string `json:"symbol"`

	// LocalSymbol is used for exchange's API (exchange package internal)
	LocalSymbol string `json:"localSymbol,omitempty"`

	// PricePrecision is the precision used for formatting price, 8 = 8 decimals
	// can be converted from price tick step size, e.g.
	//    int(math.Log10(price step size))
	PricePrecision int `json:"pricePrecision"`

	// VolumePrecision is the precision used for formatting quantity and volume, 8 = 8 decimals
	// can be converted from step size, e.g.
	//    int(math.Log10(quantity step size))
	VolumePrecision int `json:"volumePrecision"`

	// QuoteCurrency is the currency name for quote, e.g. USDT in BTC/USDT, USDC in BTC/USDC
	QuoteCurrency string `json:"quoteCurrency"`

	// BaseCurrency is the current name for base, e.g. BTC in BTC/USDT, ETH in ETH/USDC
	BaseCurrency string `json:"baseCurrency"`

	// The MIN_NOTIONAL filter defines the minimum notional value allowed for an order on a symbol.
	// An order's notional value is the price * quantity
	MinNotional fixedpoint.Value `json:"minNotional,omitempty"`
	MinAmount   fixedpoint.Value `json:"minAmount,omitempty"`

	// The LOT_SIZE filter defines the quantity
	MinQuantity fixedpoint.Value `json:"minQuantity,omitempty"`

	// MaxQuantity is currently not used in the code
	MaxQuantity fixedpoint.Value `json:"maxQuantity,omitempty"`

	// StepSize is the step size of quantity
	// can be converted from precision, e.g.
	//    1.0 / math.Pow10(m.BaseUnitPrecision)
	StepSize fixedpoint.Value `json:"stepSize,omitempty"`

	MinPrice fixedpoint.Value `json:"minPrice,omitempty"`
	MaxPrice fixedpoint.Value `json:"maxPrice,omitempty"`

	// TickSize is the step size of price
	TickSize fixedpoint.Value `json:"tickSize,omitempty"`
}
