package types

import "time"

// SliceOrderBook is a general order book structure which could be used
// for RESTful responses and websocket stream parsing
//
//go:generate callbackgen -type SliceOrderBook
type SliceOrderBook struct {
	Symbol string
	Bids   PriceVolumeSlice
	Asks   PriceVolumeSlice

	lastUpdateTime time.Time

	loadCallbacks   []func(book *SliceOrderBook)
	updateCallbacks []func(book *SliceOrderBook)
}
