package types

import "context"

type Depth string
type Speed string

type SubscribeOptions struct {
	// TODO: change to Interval type later
	Interval Interval `json:"interval,omitempty"`
	Depth    Depth    `json:"depth,omitempty"`
	Speed    Speed    `json:"speed,omitempty"`
}

type Subscription struct {
	Symbol  string           `json:"symbol"`
	Channel Channel          `json:"channel"`
	Options SubscribeOptions `json:"options"`
}

type Stream interface {
	StandardStreamEventHub

	Subscribe(channel Channel, symbol string, options SubscribeOptions)
	GetSubscriptions() []Subscription
	SetPublicOnly()
	GetPublicOnly() bool
	Connect(ctx context.Context) error
	Close() error
}

//go:generate callbackgen -type StandardStream -interface
type StandardStream struct {
	parser     Parser
	dispatcher Dispatcher

	endpointCreator EndpointCreator

	// Conn is the websocket connection
	Conn *websocket.Conn

	// ConnCtx is the context of the current websocket connection
	ConnCtx context.Context

	// ConnCancel is the cancel funcion of the current websocket connection
	ConnCancel context.CancelFunc

	// ConnLock is used for locking Conn, ConnCtx and ConnCancel fields.
	// When changing these field values, be sure to call ConnLock
	ConnLock sync.Mutex

	PublicOnly bool

	// ReconnectC is a signal channel for reconnecting
	ReconnectC chan struct{}

	// CloseC is a signal channel for closing stream
	CloseC chan struct{}

	Subscriptions []Subscription

	startCallbacks []func()

	connectCallbacks []func()

	disconnectCallbacks []func()

	// private trade update callbacks
	tradeUpdateCallbacks []func(trade Trade)

	// private order update callbacks
	orderUpdateCallbacks []func(order Order)

	// balance snapshot callbacks
	balanceSnapshotCallbacks []func(balances BalanceMap)

	balanceUpdateCallbacks []func(balances BalanceMap)

	kLineClosedCallbacks []func(kline KLine)

	kLineCallbacks []func(kline KLine)

	bookUpdateCallbacks []func(book SliceOrderBook)

	bookTickerUpdateCallbacks []func(bookTicker BookTicker)

	bookSnapshotCallbacks []func(book SliceOrderBook)

	marketTradeCallbacks []func(trade Trade)

	// Futures
	FuturesPositionUpdateCallbacks []func(futuresPositions FuturesPositionMap)

	FuturesPositionSnapshotCallbacks []func(futuresPositions FuturesPositionMap)
}
