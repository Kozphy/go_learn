package learn_bbgo

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/zixas/learn_bbgo/pkg/fixedpoint"
	"github.com/zixas/learn_bbgo/pkg/types"
)

type OrderExecutor interface {
	SubmitOrders(ctx context.Context, orders ...types.SubmitOrder) (createdOrders types.OrderSlice, err error)
	CancelOrders(ctx context.Context, orders ...types.Order) error
}

type OrderExecutionRouter interface {
	// SubmitOrdersTo submit order to a specific exchange Session
	SubmitOrdersTo(ctx context.Context, session string, orders ...types.SubmitOrder) (createdOrders types.OrderSlice, err error)
	CancelOrdersTo(ctx context.Context, session string, orders ...types.Order) error
}

type BasicRiskController struct {
	Logger *log.Logger

	MaxOrderAmount      fixedpoint.Value `json:"maxOrderAmount,omitempty" yaml:"maxOrderAmount,omitempty"`
	MinQuoteBalance     fixedpoint.Value `json:"minQuoteBalance,omitempty" yaml:"minQuoteBalance,omitempty"`
	MaxBaseAssetBalance fixedpoint.Value `json:"maxBaseAssetBalance,omitempty" yaml:"maxBaseAssetBalance,omitempty"`
	MinBaseAssetBalance fixedpoint.Value `json:"minBaseAssetBalance,omitempty" yaml:"minBaseAssetBalance,omitempty"`
}

type ExchangeOrderExecutionRouter struct {
	sessions  map[string]*ExchangeSession
	executors map[string]OrderExecutor
}

// ExchangeOrderExecutor is an order executor wrapper for single exchange instance.
//
//go:generate callbackgen -type ExchangeOrderExecutor
type ExchangeOrderExecutor struct {
	// MinQuoteBalance fixedpoint.Value `json:"minQuoteBalance,omitempty" yaml:"minQuoteBalance,omitempty"`

	Notifiability `json:"-" yaml:"-"`

	Session *ExchangeSession `json:"-" yaml:"-"`

	// private trade update callbacks
	tradeUpdateCallbacks []func(trade types.Trade)

	// private order update callbacks
	orderUpdateCallbacks []func(order types.Order)
}
