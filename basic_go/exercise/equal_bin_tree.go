package exercise

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/tour/tree"
)

func equalvalent_binary_tree() {
	fmt.Println("Exercise: Equivalent Binary Trees")
	var ch3 = make(chan int)
	go Walk(tree.New(1), ch3)
	for {
		v, ok := <-ch3
		fmt.Println(v)
		fmt.Println(ok)
		if !ok {
			fmt.Println("exit")
			break
		}
	}
	fmt.Println()
	fmt.Println("Compare tree")
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))

	fmt.Println()
	fmt.Println("sync.Mutex")
	c_safe := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c_safe.inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c_safe.Value("somekey"))

	fmt.Println()

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if v1 != v2 || ok1 != ok2 {
			return false
		}

		// if identical tree will trigger following to break out infinite for loop
		if !ok1 {
			break
		}

	}
	return true
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
}

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}
