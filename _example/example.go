package main

import (
	. "spiral_hash"
	"fmt"
)

func main() {
	p := fmt.Println

	p(GetHash(Array(1, 2, 3, 4, 5, 6, 7)))
	p(GetHash(Array(1, 2)))
	p(GetHash(Array(2, 1)))
	p(GetHash(Array(2, 2)))
}
