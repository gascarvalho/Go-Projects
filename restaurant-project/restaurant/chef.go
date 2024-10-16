package restaurant

import (
	"log"
	"sync"
	"time"
)

type Chef struct {
	id               int
	orderChan        <-chan *Order
	chefToWaiterChan chan<- *Order
}

func NewChef(id int, orderChan <-chan *Order, chefToWaiterChan chan<- *Order) *Chef {
	log.Printf("Creating a new chef with id: %d\n", id)
	return &Chef{
		id:               id,
		orderChan:        orderChan,
		chefToWaiterChan: chefToWaiterChan,
	}
}

func (chef *Chef) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	for order := range chef.orderChan {
		chef.handleOrder(order)
	}

	log.Printf("Chef %d has no more orders and is leaving.\n", chef.id)
}

func (chef *Chef) handleOrder(order *Order) {
	log.Printf("Chef %d started preparing order %d.\n", chef.id, order.id)

	select {
	case <-time.After(order.preparation):
		log.Printf("Chef %d finished preparing order %d.\n", chef.id, order.id)
		chef.chefToWaiterChan <- order
	case <-order.ctxt.Done():
		log.Printf("Chef %d stopped preparing order %d due to cancellation.\n", chef.id, order.id)
	}
}
