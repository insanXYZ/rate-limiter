package testing

import (
	"fmt"
	"rate-limiter/rl"
	"rate-limiter/utils"
	"strconv"
	"testing"
	"time"
)

func TestFixedWindowCounter(t *testing.T) {
	fixedWindowCounter := rl.NewFixedWindowCounter(5)

	utils.SetInterval(500*time.Millisecond, func() {
		statusCode := fixedWindowCounter.Add()
		fmt.Println("statusCode = " + strconv.Itoa(statusCode))
	})
}
