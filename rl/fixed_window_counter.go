package rl

import (
	"fmt"
	"rate-limiter/utils"
	"strconv"
	"time"
)

type FixedWindowCounter struct {
	counter, capacity int
}

func NewFixedWindowCounter(capacity int) *FixedWindowCounter {
	fixedWindowCounter := &FixedWindowCounter{
		capacity: capacity,
	}

	go utils.SetInterval(5*time.Second, func() {
		fixedWindowCounter.counter = 0
	})

	return fixedWindowCounter
}

func (f *FixedWindowCounter) Add() int {
	defer fmt.Println("counter have = " + strconv.Itoa(f.counter))
	if f.counter < f.capacity {
		f.counter++
		return 200
	}

	return 429
}
