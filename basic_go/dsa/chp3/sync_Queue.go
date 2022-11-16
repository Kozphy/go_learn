package chp3

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
)

type Queue_sync struct {
	waitPass    int
	waitTicket  int
	playPass    bool
	playTicket  bool
	queuePass   chan int
	queueTicket chan int
	message     chan int
}

func (queue *Queue_sync) New() {
	queue.message = make(chan int)
	queue.queuePass = make(chan int)
	queue.queueTicket = make(chan int)

	// "Go" routine handles selecting the message based on the
	// type of message and the respective queue to process it:
	go func() {
		var message int
		for {
			select {
			case message = <-queue.message:
				switch message {
				case messagePassStart:
					queue.waitPass++
				case messagePassEnd:
					queue.playPass = false
				case messageTicketStart:
					queue.waitTicket++
				case messageTicketEnd:
					queue.playTicket = false
				}
			}
			if queue.waitPass > 0 && queue.waitTicket > 0 && !queue.playPass && !queue.playTicket {
				queue.playPass = true
				queue.playTicket = true
				queue.waitTicket--
				queue.waitPass--
				queue.queuePass <- 1
				queue.queueTicket <- 1
			}
		}
	}()
}

// StartTicketIssue starts the ticket issue
func (queue *Queue_sync) StartTicketIssue() {
	queue.message <- messageTicketStart
	<-queue.queueTicket
}

// EndTicketIssue ends the ticket issue
func (queue *Queue_sync) EndTicketIssue() {
	queue.message <- messageTicketEnd
}

// ticketIssue starts and ends the ticket issue
func ticketIssue(Queue *Queue_sync) {
	for {
		// Sleep up to 10 seconds.
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		Queue.StartTicketIssue()
		fmt.Println("Ticket Issue starts")
		// Sleep up to 2 seconds.
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("Ticket Issue ends")
		Queue.EndTicketIssue()
	}
}

// StartPass ends the Pass Queue
func (Queue *Queue_sync) StartPass() {
	Queue.message <- messagePassStart
	<-Queue.queuePass
}

// EndPass ends the Pass Queue
func (Queue *Queue_sync) EndPass() {
	Queue.message <- messagePassEnd
}

// passengr method starts and ends the pass Queue
func passenger(Queue *Queue_sync) {
	fmt.Println("starting the passenger Queue")
	for {
		fmt.Println("starting the processing")
		// Sleep up to 10 seconds.
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		Queue.StartPass()
		fmt.Println("Passenger starts")
		// Sleep up to 2 seconds.
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("Passenger ends")
		Queue.EndPass()
	}
}

func Exec_sync_Queue() {
	var queue *Queue_sync
	queue.New()
	fmt.Println(queue)

	var i int
	for i = 0; i < 10; i++ {
		go passenger(queue)
	}

	var j int
	for j = 0; j < 5; j++ {
		go ticketIssue(queue)
	}

	select {}
}
