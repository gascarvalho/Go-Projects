package main

import (
	"flag"
	"log"
	"time"

	"go-project/restaurant"
)

func main() {

	numCustomers := flag.Int("customers", 5, "Number of customers")
	numChefs := flag.Int("chefs", 2, "Number of chefs")
	numWaiters := flag.Int("waiters", 2, "Number of waiters")
	customerWaitTime := flag.Duration("customer-wait-time", 10*time.Second, "Maximum wait time for customers")
	minPrepTime := flag.Duration("min-prep-time", 2*time.Second, "Minimum dish preparation time")
	maxPrepTime := flag.Duration("max-prep-time", 9*time.Second, "Maximum dish preparation time")

	//Parses the comand line arguments
	flag.Parse()

	config := restaurant.SimulationConfig{
		NumCustomers:       *numCustomers,
		NumChefs:           *numChefs,
		NumWaiters:         *numWaiters,
		CustomerTimeout:    *customerWaitTime,
		MinPreparationTime: *minPrepTime,
		MaxPreparationTime: *maxPrepTime,
	}

	log.Println("-----------Simulation Starting-----------")

	restaurant.RunSimulation(config)

	log.Println("-----------Simulation Ending-----------")
}
