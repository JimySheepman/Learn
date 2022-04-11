package main

import (
	"fmt"
	"runtime/metrics"
)

func main() {
	const myMetric = "/memory/classes/heap/free:bytes"

	sample := make([]metrics.Sample, 1)
	sample[0].Name = myMetric

	metrics.Read(sample)

	if sample[0].Value.Kind() == metrics.KindBad {
		panic(fmt.Sprintf("metric %q no longer supported", myMetric))
	}

	freeBytes := sample[0].Value.Uint64()

	fmt.Printf("free but not released memory: %d\n", freeBytes)
}
