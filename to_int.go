package main

import (
	"math"
	"strconv"
	"strings"
)

func ToInt32(input string) (int32, int32) {
	inputs := strings.Split(input, ".")
	intPart := inputs[0]
	decPart := inputs[1]
	negative := intPart[0] == '-'
	// fmt.Println(intPart, decPart)
	decPartLen := len(decPart)
	decPartInt, _ := strconv.Atoi(decPart)
	intPartInt, _ := strconv.Atoi(intPart)
	if negative {
		intPartInt = -intPartInt
	}
	if negative {
		return -int32((intPartInt*int(math.Pow10(decPartLen)) + decPartInt)), int32(decPartLen)
	}
	return int32((intPartInt*int(math.Pow10(decPartLen)) + decPartInt)), int32(decPartLen)
}
