package learn_bbgo

import (
	"github.com/zixas/learn_bbgo/pkg/types"
)

type ExchangeSession struct {
	// ---------------------------
	// Session config fields
	// ---------------------------

	// Exchange Session name
	Name         string             `json:"name,omitempty" yaml:"name,omitempty"`
	ExchangeName types.ExchangeName `json:"exchange" yaml:"exchange"`
	EnvVarPrefix string             `json:"envVarPrefix" yaml:"envVarPrefix"`
	Key          string             `json:"key,omitempty" yaml:"key,omitempty"`
	Secret       string             `json:"secret,omitempty" yaml:"secret,omitempty"`
	Passphrase   string             `json:"passphrase,omitempty" yaml:"passphrase,omitempty"`
	SubAccount   string             `json:"subAccount,omitempty" yaml:"subAccount,omitempty"`

	// UserDataStream is the connection stream of the exchange
	UserDataStream   types.Stream `json:"-" yaml:"-"`
	MarketDataStream types.Stream `json:"-" yaml:"-"`

	// Subscriptions
	// this is a read-only field when running strategy
	Subscriptions map[types.Subscription]types.Subscription `json:"-" yaml:"-"`

	usedSymbols map[string]struct{}
}

func (session *ExchangeSession) Subscribe(channel types.Channel, symbol string, options types.SubscribeOptions) *ExchangeSession {
	if channel == types.KLineChannel && len(options.Interval) == 0 {
		panic("subscription interval for kline can not be empty")
	}

	sub := types.Subscription{
		Channel: channel,
		Symbol:  symbol,
		Options: options,
	}

	session.usedSymbols[symbol] = struct{}{}
	session.Subscriptions[sub] = sub
	return session
}
