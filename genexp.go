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
		fmt.Fprintf(os.Stderr, "genexp: create a number of exponentially-distributed random floats\n")
		fmt.Fprintf(os.Stderr, "usage: genexp count [mean]\n")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	desiredMean := 0.9

	if len(os.Args) > 2 {
		desiredMean, err = strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "problem parsing desired mean %q: %v\n", os.Args[2], err)
		}
	}

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < n; i++ {
		fmt.Printf("%f\n", rand.ExpFloat64()*desiredMean)
	}
}
