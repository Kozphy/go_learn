package service

import "strings"

type MemoryService struct {
	Slots map[string]interface{}
}

type MemoryStore struct {
	Key    string
	memory *MemoryService
}

func (s *MemoryService) NewStore(id string, subIDs ...string) interface{} {
	key := strings.Join(append([]string{id}, subIDs...), ":")
	return &MemoryStore{
		Key:    key,
		memory: s,
	}
}

func NewMemoryService() *MemoryService {
	return &MemoryService{
		Slots: make(map[string]interface{}),
	}
}
