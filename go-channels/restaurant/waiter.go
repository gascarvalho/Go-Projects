package restaurant

import (
	"log"
	"sync"
	"time"
)

type Waiter struct {
	id               int
	chefToWaiterChan <-chan *Order
}

func NewWaiter(id int, chefToWaiterChan <-chan *Order) *Waiter {
	log.Printf("Creating a new waiter with id: %d\n", id)
	return &Waiter{
		id:               id,
		chefToWaiterChan: chefToWaiterChan,
	}
}

func (waiter *Waiter) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	for order := range waiter.chefToWaiterChan {
		waiter.deliverOrder(order)
	}

	log.Printf("Waiter %d has no more orders to deliver and is leaving.\n", waiter.id)
}

func (waiter *Waiter) deliverOrder(order *Order) {
	log.Printf("Waiter %d is delivering order %d to customer %d.\n", waiter.id, order.id, order.customer.id)

	select {
	case <-order.ctxt.Done():
		log.Printf("Waiter %d found out that customer %d left. Discarding order %d.\n", waiter.id, order.customer.id, order.id)
	default:
		order.customer.deliveryChan <- order
		log.Printf("Waiter %d delivered order %d to customer %d.\n", waiter.id, order.id, order.customer.id)

		log.Printf("Waiter %d is waiting for acknowledgment from customer %d for order %d.\n", waiter.id, order.customer.id, order.id)
		select {
		case ola := <-order.ackChan:
			if ola {
				log.Printf("Waiter %d received acknowledgment from customer %d for order %d.\n", waiter.id, order.customer.id, order.id)
			}
		case <-time.After(5 * time.Second):
			log.Printf("Waiter %d did not receive acknowledgment from customer %d for order %d, proceeding.\n", waiter.id, order.customer.id, order.id)
		}
	}
}
