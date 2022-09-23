package service

type PersistenceServiceFacade struct {
	Memory *MemoryService
}

func (facade *PersistenceServiceFacade) Get() MemoryService {
	return *facade.Memory
}
