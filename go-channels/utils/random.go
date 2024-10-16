package utils

import (
	"math/rand"
	"time"
)

var (
	minPreparationTime = 2 * time.Second
	maxPreparationTime = 9 * time.Second
)

func RandomDishPreparationTime() time.Duration {

	delta := maxPreparationTime - minPreparationTime

	return minPreparationTime + time.Duration(rand.Int63n(int64(delta)))
}

func SetPreparationTimes(min, max time.Duration) {
	minPreparationTime = min
	maxPreparationTime = max
}

func MinPreparationTime() time.Duration {
	return minPreparationTime
}

func MaxPreparationTime() time.Duration {
	return maxPreparationTime
}
