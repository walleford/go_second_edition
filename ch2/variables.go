package main

import (
	"fmt"
	"math"
)

func main() {
	var twenty int32 = 20
	var floatingPoint float32 = 12.365
	twenty = int32(floatingPoint)
	fmt.Println("exercise problem 1:")
	fmt.Println(twenty, floatingPoint)

	const value = 3
	var floatConst float32 = value
	var intConstant int32 = value
	fmt.Println("exercise problem 2:")
	fmt.Println(floatConst, intConstant)

	var b byte = 255
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64
	b = b + 1
	smallI += 1
	bigI += 1
	fmt.Println("exercise problem 3:")
	fmt.Println(b, smallI, bigI)
}
