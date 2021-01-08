package main

import (
	"fmt"
	"math"
)

func simpleSieve(limit int, prime []int) []int {
	// bound is square root of "high"
	var tmp float64 = float64(limit)
	bound := int(math.Sqrt(tmp))

	mark := make([]bool, bound+1, (bound+1)*2)

	for i := 0; i <= bound; i++ {
		mark[i] = true
	}

	for i := 2; i*i <= bound; i++ {
		if mark[i] == true {
			for j := i * i; j <= bound; j += i {
				mark[i] = false
			}
		}
	}

	// adding primes to vector
	for i := 2; i <= bound; i++ {
		if mark[i] == true {
			prime = append(prime, i)
		}
	}
	return prime
}

// Finds all prime numbers in range low to high
// using the primes we obtained from
// simple sieve
func segmentedSieve(low int, high int) {

	prime := make([]int, 0, (high-low)+1)
	prime = simpleSieve(high, prime) // stores primes up to sqrt(high) in prime

	mark := make([]bool, 0, 1)
	for x := 0; x < (high-low)+1; x++ {
		mark = append(mark, true)
	}

	for i := 0; i < len(prime); i++ {
		// find minimum number in [low...high] that is multiple of prime[i]
		loLim := (low / prime[i]) * prime[i]
		if loLim < low {
			loLim += prime[i]
		}
		if loLim == prime[i] {
			loLim += prime[i]
		}

		for j := loLim; j < high; j += prime[i] {
			if j != prime[i] {
				mark[j-low] = false
			}
		}

	}
	fmt.Print("[OUTPUT] ")
	// print all primes in [low...high]
	for k := low + 1; k < high; k++ {
		if mark[k-low] == true {
			fmt.Printf("%d ", k)
		}
	}
	fmt.Println("")
}
func main() {
	low := 10
	high := 20
	fmt.Printf("[INPUT] low:%d\thigh:%d\n", low, high)
	segmentedSieve(low, high)
}
