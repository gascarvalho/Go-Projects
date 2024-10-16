package restaurant

import (
	"context"
	"log"
	"sync"
	"time"

	"go-project/utils"
)

type SimulationConfig struct {
	NumCustomers       int
	NumChefs           int
	NumWaiters         int
	CustomerTimeout    time.Duration
	MinPreparationTime time.Duration
	MaxPreparationTime time.Duration
}

func (config *SimulationConfig) print() {
	log.Println("Simulation Config {")
	log.Printf("  Number of customers: %d", config.NumCustomers)
	log.Printf("  Number of chefs: %d", config.NumChefs)
	log.Printf("  Number of waiters: %d", config.NumWaiters)
	log.Printf("  Customer maximum wait time: %v", config.CustomerTimeout)
	log.Printf("  Dish preparation time range: %v to %v", config.MinPreparationTime, config.MaxPreparationTime)
	log.Println("}")
}

func RunSimulation(config SimulationConfig) {

	config.print()

	var customerWg sync.WaitGroup
	var chefWg sync.WaitGroup
	var waiterWg sync.WaitGroup

	orderChannel := make(chan *Order)
	chefToWaiterChannel := make(chan *Order)

	ctxt := context.Background()

	utils.SetPreparationTimes(config.MinPreparationTime, config.MaxPreparationTime)

	for i := 1; i <= config.NumChefs; i++ {
		chef := NewChef(i, orderChannel, chefToWaiterChannel)
		chefWg.Add(1)
		go chef.Run(&chefWg)
	}

	for i := 1; i <= config.NumWaiters; i++ {
		waiter := NewWaiter(i, chefToWaiterChannel)
		waiterWg.Add(1)
		go waiter.Run(&waiterWg)
	}

	for i := 1; i <= config.NumCustomers; i++ {
		customer := NewCustomer(i, orderChannel, config.CustomerTimeout)
		customerWg.Add(1)
		go customer.Run(&customerWg, ctxt)
	}

	customerWg.Wait()
	close(orderChannel)

	chefWg.Wait()
	close(chefToWaiterChannel)

	waiterWg.Wait()
}
