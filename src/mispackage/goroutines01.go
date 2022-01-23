package mispackage

import (
	"fmt"
	"time"
)

func animacion(retraso time.Duration) {
	for {
		for _, r := range `\|/-` {
			fmt.Printf("\r%c", r)
			time.Sleep(retraso)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
func MainGoroutines() {
	start := time.Now()
	go animacion(100 * time.Millisecond)
	const n = 45
	resultado := fib(n)
	fmt.Printf("\rFibonacci(%d)=%d, tiempo de ejecucion %s\n", n, resultado,
		time.Since(start))
}
