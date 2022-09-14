package types

import (
	"database/sql"

	"github.com/zixsa/learn_bbgo/pkg/fixedpoint"
)

type Trade struct {
	// GID is the global ID
	GID int64 `json:"gid" db:"gid"`

	// ID is the source trade ID
	ID            uint64           `json:"id" db:"id"`
	OrderID       uint64           `json:"orderID" db:"order_id"`
	Exchange      ExchangeName     `json:"exchange" db:"exchange"`
	Price         fixedpoint.Value `json:"price" db:"price"`
	Quantity      fixedpoint.Value `json:"quantity" db:"quantity"`
	QuoteQuantity fixedpoint.Value `json:"quoteQuantity" db:"quote_quantity"`
	Symbol        string           `json:"symbol" db:"symbol"`

	Side        SideType         `json:"side" db:"side"`
	IsBuyer     bool             `json:"isBuyer" db:"is_buyer"`
	IsMaker     bool             `json:"isMaker" db:"is_maker"`
	Time        Time             `json:"tradedAt" db:"traded_at"`
	Fee         fixedpoint.Value `json:"fee" db:"fee"`
	FeeCurrency string           `json:"feeCurrency" db:"fee_currency"`

	IsMargin   bool `json:"isMargin" db:"is_margin"`
	IsFutures  bool `json:"isFutures" db:"is_futures"`
	IsIsolated bool `json:"isIsolated" db:"is_isolated"`

	// The following fields are null-able fields

	// StrategyID is the strategy that execute this trade
	StrategyID sql.NullString `json:"strategyID" db:"strategy"`

	// PnL is the profit and loss value of the executed trade
	PnL sql.NullFloat64 `json:"pnl" db:"pnl"`
}
