package util

import (
	"fmt"
	"math"
	"time"
)

var (
	baseDelay = 100 * time.Millisecond
	maxDelay  = 10 * time.Second
)

// Retry 指数退避重试简单实现
func Retry(cb func() error, times int) (err error) {
	for i := 0; i < times; i++ {
		err = cb()
		if err == nil {
			return nil
		}
		sleepTime := baseDelay * time.Duration(math.Pow(2, float64(times)))
		if sleepTime > maxDelay {
			sleepTime = maxDelay
		}
		time.Sleep(sleepTime)
	}
	return fmt.Errorf("failed after %d retries, %w", times, err)
}
