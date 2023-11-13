package main

import (
	"fmt"

	"DNAFeistelNetwork/xorHelperFunctions"
)

func main() {
	fmt.Println("Demo of XOR on DNA bases")

	fmt.Println("Running example with a single base pair...")
	fmt.Println("Base pair: A, T")
	// fmt.Println("XOR of A on T: ", xor("A", "T"), "\n")

	fmt.Println("Running example of XOR two longer DNA strands...")
	strand1 := "ATGACTGA"
	strand2 := "GCTACTAG"

	xorStrandsRes := xorHelperFunctions.XorStrands(strand1, strand2)

	fmt.Println("Result of xorStrands on strand1 and strand2: ", xorStrandsRes, "\n")

	fmt.Println("Running concurrent XOR on multiple DNA strands...")
	strand3 := "TTCCAGCT"
	strand4 := "CTTAGGCA"
	strand5 := "CAATALGG"
	strand6 := "GCCTAGGA"
	// invoke go routine to run the two xorStrands concurrently
	go xorHelperFunctions.XorStrands(strand3, strand4)
	xorHelperFunctions.XorStrands(strand5, strand6)
}
