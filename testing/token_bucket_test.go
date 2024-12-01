package testing

import (
	"fmt"
	"rate-limiter/rl"
	"rate-limiter/utils"
	"strconv"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	tb := rl.NewTokenBucket(5, 1)

	utils.SetInterval(500*time.Millisecond, func() {
		statusCode := tb.Remove()
		fmt.Println("statusCode = " + strconv.Itoa(statusCode))
	})

	time.Sleep(10 * time.Second)
}
