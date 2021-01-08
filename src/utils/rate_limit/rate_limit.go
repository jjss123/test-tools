package rate_limit

import (
	"github.com/juju/ratelimit"
	"sync"
	"time"
)

var DefaultQPS float64 = 100
var DefaultGCTime = time.Hour

const (
	UnavailableToken = 0
)

type RateLimit interface {
	GetToken(string) int64
}

type rate_limit struct {
	cache map[string]*Item
	lock  sync.Mutex
	qps   float64
}

type Item struct {
	*ratelimit.Bucket
	lastUpdate time.Time
}

func newItem(qps float64) *Item {
	return &Item{ratelimit.NewBucketWithRate(qps, int64(qps)), time.Now()}
}

func NewRateLimit(qps float64) RateLimit {
	if qps <= DefaultQPS {
		qps = DefaultQPS
	}

	r := &rate_limit{cache: make(map[string]*Item), lock: sync.Mutex{}, qps: qps}
	go r.gcLoop()
	return r
}

func (r *rate_limit) GetToken(key string) int64 {
	r.lock.Lock()
	i, ok := r.cache[key]
	if !ok {
		r.cache[key] = newItem(r.qps)
		i = r.cache[key]
	}
	i.lastUpdate = time.Now()
	r.lock.Unlock()

	return i.TakeAvailable(1)
}

func (r *rate_limit) gcLoop() {
	for {
		time.Sleep(DefaultGCTime)
		r.lock.Lock()
		for k, i := range r.cache {
			if time.Now().Sub(i.lastUpdate) > DefaultGCTime {
				delete(r.cache, k)
			}
		}
		r.lock.Unlock()
	}
}
