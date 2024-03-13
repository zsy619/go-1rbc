package main

import "fmt"

func ToCut(line string) {
	end := len(line)
	tenths := int32(line[end-1] - '0')
	ones := int32(line[end-3] - '0') // line[end-2] is '.'
	var temp int32
	var semicolon int
	if line[end-4] == ';' { // positive N.N temperature
		temp = ones*10 + tenths
		semicolon = end - 4
	} else if line[end-4] == '-' { // negative -N.N temperature
		temp = -(ones*10 + tenths)
		semicolon = end - 5
	} else {
		tens := int32(line[end-4] - '0')
		if line[end-5] == ';' { // positive NN.N temperature
			temp = tens*100 + ones*10 + tenths
			semicolon = end - 5
		} else { // negative -NN.N temperature
			temp = -(tens*100 + ones*10 + tenths)
			semicolon = end - 6
		}
	}
	station := line[:semicolon]
	fmt.Println(station, temp)
}

func ToCut_Split(line string) (string, int32) {
	end := len(line)
	tenths := int32(line[end-6] - '0')
	ones := int32(line[end-7] - '0') // line[end-5] is '.'
	ext := int32(line[end-4]-'0')*1000 + int32(line[end-3]-'0')*100 + int32(line[end-2]-'0')*10 + int32(line[end-1]-'0')
	// fmt.Println("ones,tenths-->", ones, tenths)
	var temp int32
	var semicolon int
	char7 := line[end-7]
	switch char7 {
	case ';': // positive N.NNNN temperature
		temp = tenths*10000 + ext
		semicolon = end - 7
	case '-':
		temp = -(tenths*10000 + ext)
		semicolon = end - 8
	default:
		char8 := line[end-8]
		switch char8 {
		case ';': // positive NN.NNNN temperature
			temp = ones*100000 + tenths*10000 + ext
			semicolon = end - 8
		case '-': // negative -NN.NNNN temperature
			temp = -(ones*100000 + tenths*10000 + ext)
			semicolon = end - 9
		}
	}
	station := line[:semicolon]
	// fmt.Println(station, temp)
	return station, temp
}
