package restaurant

import (
	"context"
	"go-project/utils"
	"log"
	"sync"
	"time"
)

type Customer struct {
	id           int
	orderChan    chan<- *Order
	deliveryChan chan *Order
	timeout      time.Duration
}

func NewCustomer(id int, orderChan chan<- *Order, timeout time.Duration) *Customer {
	log.Printf("Creating a new customer with id: %d\n", id)
	return &Customer{
		id:           id,
		orderChan:    orderChan,
		deliveryChan: make(chan *Order, 1),
		timeout:      timeout,
	}
}

func (customer *Customer) Run(wg *sync.WaitGroup, ctxt context.Context) {
	defer wg.Done()

	ctxt, cancel := context.WithTimeout(ctxt, customer.timeout)
	defer cancel()

	prepTime := utils.RandomDishPreparationTime()

	order := &Order{
		id:          customer.id,
		preparation: prepTime,
		customer:    customer,
		ctxt:        ctxt,
		ackChan:     make(chan bool, 1),
	}

	log.Printf("Customer %d is placing an order with preparation time %.0f seconds.\n", customer.id, prepTime.Seconds())
	customer.orderChan <- order

	select {
	case <-ctxt.Done():
		log.Printf("Customer %d got tired of waiting and is leaving.\n", customer.id)
	case deliveredOrder := <-customer.deliveryChan:
		log.Printf("Customer %d received their order.\n", deliveredOrder.customer.id)

		log.Printf("Customer %d is sending acknowledgment for order %d.\n", customer.id, deliveredOrder.id)
		deliveredOrder.ackChan <- true
	}
}
