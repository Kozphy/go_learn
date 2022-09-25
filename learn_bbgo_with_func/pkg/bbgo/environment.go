package bbgo

import (
	"fmt"

	"github.com/zixsa/learn_bbgo_with_func/pkg/service"
)

func ConfigureNotificationSystem_setupInteraction_getAuthStore() {
	// case 1, without facade
	fmt.Println("without Facade")
	var MemoryService1 = service.NewMemoryService() // MemoryService
	fmt.Printf("Address of MemroyService1 = %p\n", &MemoryService1)
	fmt.Printf("Address of MemoryService1 value = %p\n", MemoryService1)
	MemoryStore1 := MemoryService1.NewStore("bbgo", "auth", "id") // MemoryStore{key, MemoryService}
	fmt.Println("MemoryStore1 = ", MemoryStore1)

	// case 2, have facade
	fmt.Println("\nhave Facade")
	var MemoryService2 = PersistenceServiceFacade.Get() // MemoryService
	fmt.Printf("Adress of MemoryService2 = %p\n", MemoryService2)
	MemoryStore2 := MemoryService2.NewStore("bbgo", "auth", "id") // MemoryStore{key, MemoryService}
	fmt.Printf("value of MemoryStore2 = %v\n", MemoryStore2)

}
