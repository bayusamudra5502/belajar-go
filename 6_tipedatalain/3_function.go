package main

import "fmt"

func haha(value int) {
	fmt.Printf("Nilai %d\n", value)
}

func main() {
	defer haha(2) // Defer itu dijalankan sebagaimana stack
	defer haha(3)

	haha(1)

	defer haha(5) 
} // Result 1 -> 5 -> 3 -> 2