package simple_impl

import (
	"time"
	"sync"
	"fmt"
)

// ref: https://google.github.io/guava/releases/19.0/api/docs/index.html?com/google/common/util/concurrent/RateLimiter.html
type RateLimiter struct {
	s *SmoothRateLimiter
	l sync.Mutex
}

func (p *RateLimiter) nowMicros() int64 {
	t := time.Now()
	return t.UnixNano()/int64(time.Microsecond)
}

func (p *RateLimiter) canAcquire(nowMicros int64) bool {
	return p.s.nextFreeTicketMicros <= nowMicros
}

func (p *RateLimiter) TryAcquire() bool {
	nowMicros := p.nowMicros()
	p.l.Lock()
	defer p.l.Unlock()
	if !p.canAcquire(nowMicros) {
		return false
	} else {
		p.s.reserveEarliestAvailable(1, nowMicros)
		return true
	}
}

func (p *RateLimiter) SetRate(permitsPerSecond int64) {
	p.l.Lock()
	defer p.l.Unlock()
	p.s.doSetRate(permitsPerSecond, p.nowMicros())
}

func (p *RateLimiter) Dump() string {
	return fmt.Sprintf("%#v\n", p.s)
}

func Create(permitsPerSecond int64) *RateLimiter {
	limiter := &RateLimiter{
		l: sync.Mutex{},
	}
	srl := &SmoothRateLimiter{}
	srl.nextFreeTicketMicros = 0
	srl.maxPermits = permitsPerSecond
	srl.stableIntervalMicros = int64(time.Second/time.Microsecond) / permitsPerSecond
	srl.doSetRate(permitsPerSecond, limiter.nowMicros())
	limiter.s = srl
	return limiter
}

type SmoothRateLimiter struct {
	maxPermits           int64
	storedPermits        int64
	stableIntervalMicros int64
	nextFreeTicketMicros int64
}

func (p *SmoothRateLimiter) resync(nowMicros int64) {
	if nowMicros > p.nextFreeTicketMicros {
		freshPermits := (nowMicros - p.nextFreeTicketMicros) / p.stableIntervalMicros
		p.storedPermits = min(p.maxPermits, p.storedPermits+freshPermits)
		p.nextFreeTicketMicros = nowMicros
	}
}

func (p *SmoothRateLimiter) doSetRate(permitsPerSecond, nowMicros int64) {
	p.resync(nowMicros)
	p.stableIntervalMicros = int64(time.Second/time.Microsecond) / permitsPerSecond
	p.maxPermits = permitsPerSecond
}

// todo bug? 当
func (p *SmoothRateLimiter) reserveEarliestAvailable(requiredPermits, nowMicros int64) {
	p.resync(nowMicros)
	storedPermitsToSpend := min(requiredPermits, p.storedPermits)
	freshPermits := requiredPermits - storedPermitsToSpend
	waitMicros := p.stableIntervalMicros * freshPermits
	fmt.Printf("%d %d %d\n", requiredPermits, p.storedPermits, waitMicros)
	p.storedPermits -= storedPermitsToSpend
	p.nextFreeTicketMicros += waitMicros
}


func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
