package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	intList := []int{rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000)}
	factors := make(map[int][]int)
	primes := []int{}

	for _, i := range intList {
		for j := 1; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				factors[i] = append(factors[i], j)
			}
		}
	}

	writer := tabwriter.NewWriter(os.Stdout, 1, 4, 1, '\t', tabwriter.AlignRight)
	for x, y := range factors {
		fmt.Fprintf(writer, "\nFactors of %v:\t%v", x, y)
	}

	for x, y := range factors {
		if len(y) == 2 {
			primes = append(primes, x)
		}
	}

	if len(primes) == 0 {
		fmt.Println("\nNo prime numbers found")
	} else {
		fmt.Fprintf(writer, "\nPrime numbers:\t%v", primes)
	}

	writer.Flush()
}
