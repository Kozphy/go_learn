package learn_bbgo

import (
	"fmt"
	"sync"
	"time"
)

var LoadedExchangeStrategies = make(map[string]SingleExchangeStrategy)
var LoadedCrossExchangeStrategies = make(map[string]CrossExchangeStrategy)

func RegisterStrategy(key string, s interface{}) {
	loaded := 0
	if d, ok := s.(SingleExchangeStrategy); ok {
		LoadedExchangeStrategies[key] = d
		loaded++
	}

	if d, ok := s.(CrossExchangeStrategy); ok {
		LoadedCrossExchangeStrategies[key] = d
		loaded++
	}

	if loaded == 0 {
		panic(fmt.Errorf("%T does not implement SingleExchangeStrategy or CrossExchangeStrategy", s))
	}
}

type SyncStatus int

const (
	SyncNotStarted SyncStatus = iota
	Syncing
	SyncDone
)

type Environment struct {
	startTime time.Time

	// syncStartTime is the time point we want to start the sync (for trades and orders)
	syncStartTime time.Time
	syncMutex     sync.Mutex

	syncStatusMutex sync.Mutex
	syncStatus      SyncStatus
	syncConfig      *SyncConfig

	sessions map[string]*ExchangeSession
}

func NewEnvironment() *Environment {

	now := time.Now()
	return &Environment{
		// default trade scan time
		syncStartTime: now.AddDate(-1, 0, 0), // defaults to sync from 1 year ago
		sessions:      make(map[string]*ExchangeSession),
		startTime:     now,

		syncStatus: SyncNotStarted,
	}
}
