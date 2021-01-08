package rate_limit

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestSimplyGetToken(t *testing.T) {
	l := NewRateLimit(DefaultQPS)
	qps := int(DefaultQPS)
	for i := 0; i < qps-1; i++ {
		if l.GetToken("test") == UnavailableToken {
			t.Fatal("Get Token return zero when the number of requests less than qps")
		}
	}
}

func TestGC(t *testing.T) {
	DefaultGCTime = time.Second

	r := &rate_limit{cache: make(map[string]*Item), lock: sync.Mutex{}, qps: DefaultQPS}
	go r.gcLoop()

	for key := 0; key < 10; key++ {
		r.GetToken(strconv.Itoa(key))
	}

	time.Sleep(2 * time.Second)

	if len(r.cache) != 0 {
		t.Fatal("RateLimit GC did not clean cache")
	}
}

func TestMultipleGetToken(t *testing.T) {
	l := NewRateLimit(DefaultQPS)
	qps := int(DefaultQPS)

	for i := 0; i < qps-1; i++ {
		if l.GetToken("test1") == UnavailableToken || l.GetToken("test2") == UnavailableToken {
			t.Fatal("Get Token with multiple key return zero when the number of requests less than qps")
		}
	}
}

func TestFastGetToken(t *testing.T) {
	l := NewRateLimit(DefaultQPS)
	qps := int(DefaultQPS)
	for i := 0; i < qps; i++ {
		if l.GetToken("test") == UnavailableToken {
			t.Fatal("Get Token return zero when the number of requests less than qps")
		}
	}
	for i := 0; i < qps; i++ {
		if l.GetToken("test") > UnavailableToken {
			t.Fatal("Get Token recover too fast")
		}
	}
	time.Sleep(time.Second)
	for i := 0; i < qps; i++ {
		if l.GetToken("test") == UnavailableToken {
			t.Fatal("Get Token return zero when the number of requests less than qps")
		}
	}
	for i := 0; i < qps; i++ {
		if l.GetToken("test") > UnavailableToken {
			t.Fatal("Get Token recover too fast")
		}
	}
}
