package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "gennormal: create a number of normally-distributed random floats\n")
		fmt.Fprintf(os.Stderr, "suage: gennormal count [mean [std dev]]\n")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	desiredMean := 0.9
	stdDev := 1.0

	if len(os.Args) > 2 {
		desiredMean, err = strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "problem parsing desired mean %q: %v\n", os.Args[2], err)
		}
		if len(os.Args) > 3 {
			stdDev, err = strconv.ParseFloat(os.Args[3], 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "problem parsing desired std dev %q: %v\n", os.Args[3], err)
			}
		}
	}

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < n; i++ {
		fmt.Printf("%f\n", rand.NormFloat64()*stdDev+desiredMean)
	}
}
