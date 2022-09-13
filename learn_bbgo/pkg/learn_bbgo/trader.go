package learn_bbgo

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type StrategyID interface {
	ID() string
}

type SingleExchangeStrategy interface {
	StrategyID
	Run(ctx context.Context, orderExecutor OrderExecutor, session *ExchangeSession) error
}

type CrossExchangeStrategy interface {
	StrategyID
	CrossRun(ctx context.Context, orderExecutionRouter OrderExecutionRouter, sessions map[string]*ExchangeSession) error
}

type Logger interface {
	Warnf(message string, args ...interface{})
	Errorf(message string, args ...interface{})
	Infof(message string, args ...interface{})
}

type Trader struct {
	environment *Environment

	riskControls *RiskControls

	crossExchangeStrategies []CrossExchangeStrategy
	exchangeStrategies      map[string][]SingleExchangeStrategy

	logger Logger
}

func NewTrader(environ *Environment) *Trader {
	return &Trader{
		environment:        environ,
		exchangeStrategies: make(map[string][]SingleExchangeStrategy),
		logger:             log.StandardLogger(),
	}
}
