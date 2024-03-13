package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ToFloat64(input string) float64 {
	tempBytes := []byte(input)
	negative := false
	index := 0
	if tempBytes[index] == '-' {
		index++
		negative = true
	}
	temp := float64(tempBytes[index] - '0') // parse first digit
	index++
	if tempBytes[index] != '.' {
		fmt.Println(float64(tempBytes[index] - '0'))
		temp = temp*10 + float64(tempBytes[index]-'0') // parse optional second digit
		index++
	}
	index++                                    // skip '.'
	temp += float64(tempBytes[index]-'0') / 10 // parse decimal digit
	if negative {
		temp = -temp
	}
	return temp
}

func ToFloat64_1(input string) float64 {
	inputs := strings.Split(input, ".")
	intPart := inputs[0]
	decPart := inputs[1]
	negative := intPart[0] == '-'
	// fmt.Println(intPart, decPart)
	decPartLen := len(decPart)
	decPartInt, _ := strconv.Atoi(decPart)
	decPartFloat := float64(decPartInt) / math.Pow10(decPartLen)
	intPartInt, _ := strconv.Atoi(intPart)
	if negative {
		intPartInt = -intPartInt
	}
	if negative {
		return -(float64(intPartInt) + decPartFloat)
	}
	return float64(intPartInt) + decPartFloat
}
