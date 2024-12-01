package testing

import (
	"fmt"
	"rate-limiter/rl"
	"rate-limiter/utils"
	"testing"
	"time"
)

func TestLeakyBucket(t *testing.T) {
	leakyBucket := rl.NewLeakyBucket(5, 1)

	utils.SetInterval(500*time.Millisecond, func() {
		i := leakyBucket.Add()
		fmt.Println(i)
	})
}
