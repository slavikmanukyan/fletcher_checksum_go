package main

import (
	"fmt"
)

func fletcher(data []byte) uint16 {
	sum1 := uint16(0)
	sum2 := uint16(0)

	for index := 0; index < len(data); index++ {
		sum1 = (sum1 + uint16(data[index])) % 255
		sum2 = (sum2 + sum1) % 255
	}

	return (sum2 << 8) | sum1
}

func main() {
	fmt.Printf("%xd\n", fletcher([]byte("abcdeabcde"))) // 46e1d --> 18145
}
