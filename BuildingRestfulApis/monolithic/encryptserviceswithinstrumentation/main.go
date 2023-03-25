package main

import (
	kitlog "github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	transport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"instrumentation/helpers"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := kitlog.NewLogfmtLogger(os.Stderr)
	fieldKeys := []string{"method", "error"}
	// This is a request count
	//operation for my microservice called Encryption.
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "encryption",
		Subsystem: "my_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "encryption",
		Subsystem: "my_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	var svc helpers.EncryptService
	svc = helpers.EncryptServiceInstance{}
	svc = helpers.LoggingMiddleware{Logger: logger, Next: svc}
	svc = helpers.InstrumentingMiddleware{RequestCount: requestCount, RequestLatency: requestLatency, Next: svc}

	encryptHandler := transport.NewServer(helpers.MakeEncryptEndpoint(svc), helpers.DecodeEncryptRequest, helpers.EncodeResponse)
	decryptHandler := transport.NewServer(helpers.MakeDecryptEndpoint(svc), helpers.DecodeDecryptRequest, helpers.EncodeResponse)

	http.Handle("/encrypt", encryptHandler)
	http.Handle("/decrypt", decryptHandler)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":1205", nil))
}
