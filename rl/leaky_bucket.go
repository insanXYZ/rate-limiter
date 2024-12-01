package rl

import (
	"fmt"
	"rate-limiter/utils"
	"strconv"
	"time"
)

type LeakyBucket struct {
	tokens, capacity int
}

func NewLeakyBucket(capacity, outFlowRate int) *LeakyBucket {
	leakyBucket := &LeakyBucket{
		tokens:   0,
		capacity: capacity,
	}

	go utils.SetInterval(1*time.Second, func() {
		leakyBucket.remove(outFlowRate)
	})

	return leakyBucket
}

func (l *LeakyBucket) Add() int {
	fmt.Println("token have = " + strconv.Itoa(l.tokens))
	if l.tokens >= l.capacity {
		return 429
	}

	l.tokens++

	return 200
}

func (l *LeakyBucket) remove(outFlowRate int) {
	count := l.tokens - outFlowRate

	if l.tokens > 0 {
		l.tokens = count
	}
}
