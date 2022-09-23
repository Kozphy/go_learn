package bbgo

import (
	"fmt"

	"github.com/k0kubun/pp/v3"
	"github.com/zixsa/learn_bbgo_with_func/pkg/service"
)

func ConfigureNotificationSystem_setupInteraction_getAuthStore() {
	// case 1, without facade
	fmt.Println("without Facade")
	nms1 := service.NewMemoryService()
	fmt.Printf("Address of nms1 = %p\n", &nms1)
	mst1 := nms1.NewStore("bbgo", "auth", "id")
	fmt.Println("mst1 = ", mst1)
	// fmt.Printf("Address of mst1 = %p\n", mst1)

	// case2, have facade
	fmt.Println("\nhave Facade")
	var persistence = PersistenceServiceFacade.Get()
	fmt.Printf("Adress of persistence = %p\n", &persistence)
	ms := persistence.NewStore("bbgo", "auth", "default")
	fmt.Printf("value of ms = %v\n", ms)
	// fmt.Printf("Address of ms = %p\n", ms)
	pp.Print(ms)
}
