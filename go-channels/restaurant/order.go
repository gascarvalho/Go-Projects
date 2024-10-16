package restaurant

import (
	"context"
	"time"
)

type Order struct {
	id          int
	preparation time.Duration
	customer    *Customer
	ctxt        context.Context
	ackChan     chan bool
}
