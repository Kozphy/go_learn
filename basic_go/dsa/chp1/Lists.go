package chp1

import (
	"container/list"
	"fmt"
)

func Execute_lists() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)

	for elem := intList.Front(); elem != nil; elem = elem.Next() {
		fmt.Println(elem.Value.(int))
	}
}
