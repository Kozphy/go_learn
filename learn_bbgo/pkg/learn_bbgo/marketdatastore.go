package learn_bbgo

import "github.com/zixsa/learn_bbgo/pkg/types"

// MarketDataStore receives and maintain the public market data of a single symbol
//
//go:generate callbackgen -type MarketDataStore
type MarketDataStore struct {
	Symbol string

	// KLineWindows stores all loaded klines per interval
	KLineWindows map[types.Interval]*types.KLineWindow `json:"-"`

	kLineWindowUpdateCallbacks []func(interval types.Interval, klines types.KLineWindow)
	kLineClosedCallbacks       []func(k types.KLine)
}
