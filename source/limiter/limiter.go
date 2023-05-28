package limiter

import "strings"

func LimiterTem(svr string) string {
	t :=
		`
package limiter

import (
	"sample/source/cache"
	"time"

	"github.com/go-redis/redis_rate/v10"
)

func GetLimter() *redis_rate.Limiter {
	limiter := redis_rate.NewLimiter(cache.GetRedis().GetClient())
	return limiter
}

func PerDuration(rate int, duration time.Duration) redis_rate.Limit {
	return redis_rate.Limit{
		Rate:   rate,
		Burst:  rate,
		Period: duration / time.Duration(rate),
	}
}


`
	t = strings.ReplaceAll(t, "sample", svr)
	return t
}
