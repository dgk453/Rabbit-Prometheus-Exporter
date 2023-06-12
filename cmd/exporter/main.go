package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

func main() {
	c := collectors.NewGoCollector()
	prometheus.MustRegister(c)
}
