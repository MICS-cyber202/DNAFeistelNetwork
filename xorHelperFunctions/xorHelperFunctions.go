package xorHelperFunctions

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

// DNA base XOR helper function, to use in xorStrands function
func XorBasePair(base1 string, base2 string) (xorRes string, err error) {

	// error checking for base1 or base2 != "A" "T" "C" or "G"
	var bases string = "ACTG"
	if strings.Index(bases, base1) == -1 || strings.Index(bases, base2) == -1 || len(base1) != 1 || len(base2) != 1 {
		errorMessage := fmt.Sprintf("Either %s or %s is not a single DNA base of value A, T, C, or G", base1, base2)
		return "", errors.New(errorMessage)
	}

	// use a hashmap to store four way DNA XOR rules
	// the hashmap is much faster than an if/else logic for retrieiving the XOR result
	var baseXorMap = make(map[string]string)
	baseXorMap["AA"] = "A"
	baseXorMap["AT"] = "C"
	baseXorMap["AG"] = "G"
	baseXorMap["AC"] = "T"
	baseXorMap["TA"] = "G"
	baseXorMap["TT"] = "T"
	baseXorMap["TG"] = "C"
	baseXorMap["TC"] = "G"
	baseXorMap["GA"] = "C"
	baseXorMap["GT"] = "G"
	baseXorMap["GG"] = "T"
	baseXorMap["GC"] = "C"
	baseXorMap["CA"] = "T"
	baseXorMap["CT"] = "A"
	baseXorMap["CG"] = "A"
	baseXorMap["CC"] = "A"

	// concatenate the two arguments
	var basePair string = base1 + base2

	// retrieve value from the hashmap
	xorRes = baseXorMap[basePair]

	// return the xorRes (defined in function signature)
	return
}

// function to XOR two DNA strands
// export function from module (enable exporting by starting the function name with a capital letter)
func XorStrands(strand1 string, strand2 string) (strand1XORstrands2 string) {
	// use Sleep to enable concurrent invocation
	time.Sleep(1 * time.Microsecond)

	fmt.Println("Loop through each base in the two strands and perform XOR")
	fmt.Println("first strand: ", strand1)
	fmt.Println("second strand: ", strand2)

	// iterate through the two strands and call the helper xorBasePair function
	// concatenates result into strand1XORstrands2
	for i := 0; i < len(strand1); i++ {
		base1 := string(strand1[i])
		base2 := string(strand2[i])
		base1XORbase2, err := XorBasePair(base1, base2)
		if (err) != nil {
			log.Fatal(err)
		}
		fmt.Println(base1XORbase2)
		strand1XORstrands2 += base1XORbase2
	}

	fmt.Println("Result: ", strand1XORstrands2)
	return
}
