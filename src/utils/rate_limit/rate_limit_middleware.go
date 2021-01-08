package rate_limit

import (
	"github.com/prometheus/client_golang/prometheus"
	"net"
	"net/http"
	"testTools/src/utils/clog"
)

const (
	DefaultRateLimitLabel = "ratelimit"
)

type RateLimitMiddleware struct {
	GetKeyFromRequest func(*http.Request) string
	LimitUnavaible    func(http.ResponseWriter, *http.Request) bool
	c                 *prometheus.CounterVec
	limit             RateLimit
}

type keyFunc func(*http.Request) string
type limitFunc func(http.ResponseWriter, *http.Request) bool

func NewRateLimitMiddleware(qps float64, kf keyFunc, lf limitFunc) *RateLimitMiddleware {
	if kf == nil {
		kf = defaultGetKeyFunc
	}
	if lf == nil {
		lf = defaultLimitFunc
	}

	return &RateLimitMiddleware{
		GetKeyFromRequest: kf,
		LimitUnavaible:    lf,
		c:                 nil,
		limit:             NewRateLimit(qps),
	}
}

func (m *RateLimitMiddleware) Collector(name, help string) *prometheus.CounterVec {
	m.c = prometheus.NewCounterVec(prometheus.CounterOpts{Name: name, Help: help}, []string{DefaultRateLimitLabel})
	return m.c
}

func (m *RateLimitMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		//Get key from request.
		key := m.GetKeyFromRequest(req)
		if m.limit.GetToken(key) != UnavailableToken {
			next.ServeHTTP(w, req)
			return
		}

		//Record unavailable request
		if m.c != nil {
			m.c.With(prometheus.Labels{DefaultRateLimitLabel: key}).Add(1)
		}

		clog.Blog.Warningf("rate-limit : key %s request too fast with uri(%s)", key, req.RequestURI)

		//Do something after unavailable, if return true, middleware will call next function
		if m.LimitUnavaible(w, req) {
			next.ServeHTTP(w, req)
		}
	})
}

//Get key from req.RemoteHost
func defaultGetKeyFunc(req *http.Request) string {
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	return host
}

//Discard request and do nothing
func defaultLimitFunc(w http.ResponseWriter, req *http.Request) bool {
	return false
}
