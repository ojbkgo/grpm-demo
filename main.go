package main

import (
	"context"
	c "github.com/ojbkgo/grpc-demo/cli"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

type action struct {
}

func (a *action) Say(ctx context.Context, msg *c.Msg) (resp *c.MsgResp, err error) {
	resp = &c.MsgResp{}
	resp.Val = msg.Val + time.Now().String()
	return resp, nil
}

func main() {

	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "counter_demo",
		Help: "This is help",
	}, []string{"val", "tag", "age"})

	cputemp := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temprature",
		Help: "This is help",
	})


	//sum := prometheus.NewSummaryVec(prometheus.SummaryOpts{Name: "test_summary"}, []string{"instance"})
	//sum.With().Observe()
	h := prometheus.NewHistogram(prometheus.HistogramOpts{Name: "demo", Buckets: []float64{0, 1, 2, 3, 4}})


	prometheus.MustRegister(counter)
	prometheus.MustRegister(h)
	prometheus.MustRegister(cputemp)

	h.Observe(1.1)
	h.Observe(1.2)
	h.Observe(1.3)
	h.Observe(1.1)
	h.Observe(1.1)
	h.Observe(3.1)

	cputemp.Set(109.9)
	go func() {
		tk := time.NewTicker(time.Second)

		for range tk.C {
			counter.With(prometheus.Labels{"tag" : "hello", "val":"val123", "age":"1234"}).Add(1)
			counter.With(prometheus.Labels{"tag" : "hello2", "val":"val123", "age":"1234"}).Add(2)
			cputemp.Set(rand.Float64())
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe("0.0.0.0:80", nil)
}
