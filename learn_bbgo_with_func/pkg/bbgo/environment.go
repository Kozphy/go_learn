package enviroment

import (
	"fmt"
	"strings"
)

type MemoryService struct {
	Slots map[string]interface{}
}

type MemoryStore struct {
	Key    string
	memory *MemoryService
}

func NewMemoryService() *MemoryService {
	return &MemoryService{
		Slots: make(map[string]interface{}),
	}
}

func (s *MemoryService) NewStore(id string, subIDs ...string) interface{} {
	key := strings.Join(append([]string{id}, subIDs...), ":")
	return &MemoryStore{
		Key:    key,
		memory: s,
	}
}

func ConfigureNotificationSystem() {
	ms := NewMemoryService()
	ms.NewStore("bbgo", "auth", "default")
	fmt.Println(ms)
}
