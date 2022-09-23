package bbgo

import (
	"github.com/zixsa/learn_bbgo_with_func/pkg/service"
)

var DefaultPersistenceServiceFacade = &service.PersistenceServiceFacade{
	Memory: service.NewMemoryService(),
}

var PersistenceServiceFacade = DefaultPersistenceServiceFacade
