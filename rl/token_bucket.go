package rl

import (
	"fmt"
	"rate-limiter/utils"
	"strconv"
	"time"
)

type TokenBucket struct {
	tokens, capacity int
}

func NewTokenBucket(capacity, refilRate int) *TokenBucket {

	tokenBucket := &TokenBucket{
		tokens:   capacity,
		capacity: capacity,
	}

	go utils.SetInterval(1*time.Second, func() {
		tokenBucket.addToken(refilRate)
	})

	return tokenBucket
}

func (t *TokenBucket) addToken(refilRate int) {
	count := t.tokens + refilRate

	if count < t.capacity {
		t.tokens = count
	}

}

func (t *TokenBucket) Remove() int {

	defer func() {
		fmt.Println("tokens have = " + strconv.Itoa(t.tokens))
	}()
	if t.tokens > 0 {
		t.tokens--
		return 200
	}

	return 400
}
