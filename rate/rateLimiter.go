package rate

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mutex      sync.Mutex
	lastAccess time.Time
	interval   time.Duration
	maxCount   int
	count      int
}

// contructor
func NewRateLimiter(interval time.Duration, maxCount int) *RateLimiter {
	return &RateLimiter{
		lastAccess: time.Now(),
		interval:   interval,
		maxCount:   maxCount,
		count:      0,
	}
}

// methods
func (limiter *RateLimiter) Allow() bool {
	limiter.mutex.Lock()
	defer limiter.mutex.Unlock()

	now := time.Now()

	if now.Sub(limiter.lastAccess) >= limiter.interval {
		limiter.count = 0
		limiter.lastAccess = now
	}

	if limiter.count < limiter.maxCount {
		limiter.count++
		return true
	}

	return false
}
