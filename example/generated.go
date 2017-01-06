// Package example was automatically generated by Shenzhen Go.
package example // import "github.com/google/shenzhen-go/example"

import (
	"fmt"
	"sync"
)

var (
	div2 = make(chan int, 0)
	div3 = make(chan int, 0)
	out  = make(chan int, 0)
	raw  = make(chan int, 0)
)

// Run executes all the goroutines associated with the graph that generated
// this package, and waits for any that were marked as "wait for this to
// finish" to finish before returning.
func Run() {
	var wg sync.WaitGroup

	// Filter divisible by 2
	wg.Add(1)
	go func() {
		defer wg.Done()

		for n := range raw {
			if n > 2 && n%2 == 0 {
				continue
			}
			div2 <- n
		}
		close(div2)
	}()

	// Filter divisible by 3
	wg.Add(1)
	go func() {
		defer wg.Done()

		for n := range div2 {
			if n > 3 && n%3 == 0 {
				continue
			}
			div3 <- n
		}
		close(div3)
	}()

	// Filter divisible by 5
	wg.Add(1)
	go func() {
		defer wg.Done()

		for n := range div3 {
			if n > 5 && n%5 == 0 {
				continue
			}
			out <- n
		}
		close(out)
	}()

	// Generate integers ≥ 2
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 2; i < 50; i++ {
			raw <- i
		}
		close(raw)
	}()

	// Print output
	wg.Add(1)
	go func() {
		defer wg.Done()

		for n := range out {
			fmt.Println(n)
		}
	}()

	// foooo
	wg.Add(1)
	go func() {
		defer wg.Done()

		lkasdjfklajsdf
	}()

	// Wait for the end
	wg.Wait()
}
