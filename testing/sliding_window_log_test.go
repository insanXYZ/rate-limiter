package testing

import (
	"fmt"
	"rate-limiter/rl"
	"rate-limiter/utils"
	"strconv"
	"testing"
	"time"
)

func TestSlidingWindowLog(t *testing.T) {
	slidingWindowLog := rl.NewSlidingWindowLog(10)

	utils.SetInterval(1*time.Second, func() {
		statusCode := slidingWindowLog.Add()
		fmt.Println("statusCode = " + strconv.Itoa(statusCode))
	})
}
