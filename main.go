package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func calc(toValue int) int {
	var res = 0
	var pos = 0
	for pos < toValue {
		if rand.Intn(pos+1) == pos {
			res = pos
		}
		pos++
	}
	return res
}

func main() {
	flag.Parse()

	arraySize := *flag.Int("size", 100, "size of array")
	iters := *flag.Int("iter", 1000, "interations count")

	rand.Seed(time.Now().UnixNano())

	var count = make([]int, arraySize)
	for i := 0; i < arraySize*iters; i++ {
		count[calc(arraySize)]++
	}

	var results = make([]float64, arraySize)
	for i := 0; i < arraySize; i++ {
		results[i] = float64(count[i]) / float64(iters)
	}

	fmt.Printf("%v", results)
}
