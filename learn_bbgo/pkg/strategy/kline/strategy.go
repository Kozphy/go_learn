package kline

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/zixas/learn_bbgo/pkg/learn_bbgo"

	"github.com/zixas/learn_bbgo/pkg/types"
)

const ID = "kline"

var log = logrus.WithField("strategy", ID)

func init() {
	learn_bbgo.RegisterStrategy(ID, &Strategy{})
}

type Strategy struct {
	Symbol        string               `json:"symbol"`
	MovingAverage types.IntervalWindow `json:"movingAverage"`
}

func (s *Strategy) ID() string {
	return ID
}

func (s *Strategy) Subscribe(session *learn_bbgo.ExchangeSession) {
	session.Subscribe(types.KLineChannel, s.Symbol, types.SubscribeOptions{Interval: s.MovingAverage.Interval})
}

func (s *Strategy) Run(ctx context.Context, orderExecutor learn_bbgo.OrderExecutor, session *learn_bbgo.ExchangeSession) error {
	session.MarketDataStream.OnKLineClosed(func(kline types.KLine) {
		// skip k-lines from other symbols
		if kline.Symbol != s.Symbol {
			return
		}

		log.Infof("%s", kline.String())
	})
	return nil
}
