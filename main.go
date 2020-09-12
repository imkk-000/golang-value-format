package main

import (
	"fmt"
)

func main() {
	// reference: https://golang.org/pkg/fmt/
	fmt.Printf("%d\n", 1_0_0_0_0_0_0)
	fmt.Printf("%v\n", 1_0_0_0_0_0_0)
	fmt.Printf("%.11F\n", float64(9_8_7_6_5_4_3_2_1_0.0_1_2_3_4_5_6_7_8_9))
	fmt.Printf("%d\n", 0b1_1111_1111)
	fmt.Printf("%b\n", 512)
	fmt.Printf("%d\n", 0o0777)
	fmt.Printf("%02o\n", 511)

	fmt.Printf("%d\n", 1<<3)
	fmt.Printf("%d\n", 8>>3)
	fmt.Printf("%04b\n", 0b1010&0b0101)
	fmt.Printf("%04b\n", 0b1010|0b0101)
	fmt.Printf("%04b\n", 0b1010^0b0101)
	fmt.Printf("%04b\n", 0b1010&^0b0101)
	fmt.Printf("%04b\n", 0b1010^^0b0101)
	fmt.Printf("%04b\n", ^uint8(0b0011)&0x0F)
}
