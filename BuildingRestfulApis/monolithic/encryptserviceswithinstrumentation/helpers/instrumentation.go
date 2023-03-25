/* INstrumentation can measure the number of request oper service and the latency in terms of parameters such
as counter and histogram, respoectively.*/

package helpers

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"time"
)

// InstrumentingMiddleware is a struct representing middleware
type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Next           EncryptService
}

/*
The preceding middleware log sthe reqeust count and latency to the metrics provided by the prometheus client.
*/

func (mw InstrumentingMiddleware) Encrypt(ctx context.Context, key string, text string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "encrypt", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)                                 // when a request is passed through this middleware, this line increments the counter
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds()) // this line observes the latency bycalculating the difference between the request arrival time and filan time
	}(time.Now())
	output, err = mw.Next.Encrypt(ctx, key, text)
	return
}

func (mw InstrumentingMiddleware) Decrypt(ctx context.Context, key string, text string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "decrypt", "error", "false"}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	output, err = mw.Next.Decrypt(ctx, key, text)
	return
}
