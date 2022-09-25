package service

type PersistenceServiceFacade struct {
	Memory *MemoryService
}

func (facade *PersistenceServiceFacade) Get() PersistenceService {
	return facade.Memory
}
