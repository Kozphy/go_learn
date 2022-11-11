package chp3

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
