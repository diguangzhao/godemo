package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type MetricValue struct {
	Value float64
	Time  time.Time
}

func main() {
	metric := MetricValue{
		Value: 1.0,
		Time:  time.Now(),
	}

	// Store a value

	m0 := map[string]MetricValue{}
	m0["foo"] = metric

	m1 := sync.Map{}
	m1.Store("foo", metric) // not type-checked

	// Load a value and print its square

	foo0 := m0["foo"].Value // rely on zero-value hack if not present
	fmt.Printf("Foo square = %f\n", math.Pow(foo0, 2))

	foo1 := 0.0
	if x, ok := m1.Load("foo"); ok { // have to make sure it's present (not bad, actually)
		foo1 = x.(MetricValue).Value // cast interface{} value
	}
	fmt.Printf("Foo square = %f\n", math.Pow(foo1, 2))

	// Sum all elements

	sum0 := 0.0
	for _, v := range m0 { // built-in range iteration on map
		sum0 += v.Value
	}
	fmt.Printf("Sum = %f\n", sum0)

	sum1 := 0.0
	m1.Range(func(key, value interface{}) bool { // no 'range' for you! Provide a function
		sum1 += value.(MetricValue).Value // with untyped interface{} parameters
		return true                       // continue iteration
	})
	fmt.Printf("Sum = %f\n", sum1)
}
