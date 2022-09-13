package types

import (
	"fmt"

	"github.com/zixas/learn_bbgo/pkg/fixedpoint"
)

// KLine uses binance's kline as the standard structure
type KLine struct {
	GID      uint64       `json:"gid" db:"gid"`
	Exchange ExchangeName `json:"exchange" db:"exchange"`

	Symbol string `json:"symbol" db:"symbol"`

	StartTime Time `json:"startTime" db:"start_time"`
	EndTime   Time `json:"endTime" db:"end_time"`

	Interval Interval `json:"interval" db:"interval"`

	Open                     fixedpoint.Value `json:"open" db:"open"`
	Close                    fixedpoint.Value `json:"close" db:"close"`
	High                     fixedpoint.Value `json:"high" db:"high"`
	Low                      fixedpoint.Value `json:"low" db:"low"`
	Volume                   fixedpoint.Value `json:"volume" db:"volume"`
	QuoteVolume              fixedpoint.Value `json:"quoteVolume" db:"quote_volume"`
	TakerBuyBaseAssetVolume  fixedpoint.Value `json:"takerBuyBaseAssetVolume" db:"taker_buy_base_volume"`
	TakerBuyQuoteAssetVolume fixedpoint.Value `json:"takerBuyQuoteAssetVolume" db:"taker_buy_quote_volume"`

	LastTradeID    uint64 `json:"lastTradeID" db:"last_trade_id"`
	NumberOfTrades uint64 `json:"numberOfTrades" db:"num_trades"`
	Closed         bool   `json:"closed" db:"closed"`
}

func (k KLine) String() string {
	return fmt.Sprintf("%s %s %s %s O: %.4f H: %.4f L: %.4f C: %.4f CHG: %.4f MAXCHG: %.4f V: %.4f QV: %.2f TBBV: %.2f",
		k.Exchange.String(),
		k.StartTime.Time().Format("2006-01-02 15:04"),
		k.Symbol, k.Interval, k.Open.Float64(), k.High.Float64(),
		k.Low.Float64(), k.Close.Float64(), k.GetChange().Float64(),
		k.GetMaxChange().Float64(), k.Volume.Float64(), k.QuoteVolume.Float64(),
		k.TakerBuyBaseAssetVolume.Float64())
}

func (k KLine) GetHigh() fixedpoint.Value {
	return k.High
}

func (k KLine) GetLow() fixedpoint.Value {
	return k.Low
}

func (k KLine) GetOpen() fixedpoint.Value {
	return k.Open
}

func (k KLine) GetClose() fixedpoint.Value {
	return k.Close
}

func (k KLine) GetMaxChange() fixedpoint.Value {
	return k.GetHigh().Sub(k.GetLow())
}

// GetChange returns Close price - Open price.
func (k KLine) GetChange() fixedpoint.Value {
	return k.Close.Sub(k.Open)
}
