package chp3

import "fmt"

// Queue - Array of Orders Type
type Queue []*Order

// Order class
type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

func (order *Order) New(priority int, quantity int, product string, customerName string) {
	order.priority = priority
	order.quantity = quantity
	order.product = product
	order.customerName = customerName
}

func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		var appended bool
		appended = false
		var i int
		var addedOrder *Order
		for i, addedOrder = range *queue {
			// order of priority is big to small.
			if order.priority > addedOrder.priority {
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}

		if !appended {
			*queue = append(*queue, order)
		}
	}
}

func Exec_Queue() {
	var queue Queue
	queue = make(Queue, 0)

	var order1 *Order = &Order{}

	var priority1 int = 2
	var quantity1 int = 20
	var product1 string = "Computer"
	var customerName1 string = "Greg White"

	order1.New(priority1, quantity1, product1, customerName1)

	var order2 *Order = &Order{}

	var priority2 int = 1
	var quantity2 int = 10
	var product2 string = "Monitor"
	var customerName2 string = "John Smith"

	order2.New(priority2, quantity2, product2, customerName2)
	queue.Add(order1)
	queue.Add(order2)

	var i int
	for i = 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}
}
