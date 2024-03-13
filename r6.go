package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
)

func Run6(inputPath string, output io.Writer) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	stationStats := make(map[string]*StatsInt32)

	buf := make([]byte, 1024*1024)
	readStart := 0
	for {
		n, err := f.Read(buf[readStart:])
		if err != nil && err != io.EOF {
			return err
		}
		if readStart+n == 0 {
			break
		}
		chunk := buf[:readStart+n]

		newline := bytes.LastIndexByte(chunk, '\n')
		if newline < 0 {
			break
		}
		remaining := chunk[newline+1:]
		chunk = chunk[:newline+1]

		for {
			station, after, hasSemi := bytes.Cut(chunk, []byte(";"))
			if !hasSemi {
				break
			}

			index := 0
			negative := false
			if after[index] == '-' {
				negative = true
				index++
			}
			temp := int32(after[index] - '0')
			index++
			if after[index] == '.' { // N.NNNN
				index++
			} else if after[index+1] == '.' { // NN.NNNN
				temp = temp*10 + int32(after[index]-'0')
				index += 2
			}
			temp = temp*10000 +
				int32(after[index]-'0')*1000 +
				int32(after[index+1]-'0')*100 +
				int32(after[index+2]-'0')*10 +
				int32(after[index+3]-'0')
			index += 5 // skip last digit and '\n'
			if negative {
				temp = -temp
			}
			chunk = after[index:]

			s := stationStats[string(station)]
			if s == nil {
				stationStats[string(station)] = &StatsInt32{
					Min:   temp,
					Max:   temp,
					Sum:   int64(temp),
					Count: 1,
				}
			} else {
				s.Min = min(s.Min, temp)
				s.Max = max(s.Max, temp)
				s.Sum += int64(temp)
				s.Count++
			}
		}

		readStart = copy(buf, remaining)
	}

	stations := make([]string, 0, len(stationStats))
	for station := range stationStats {
		stations = append(stations, station)
	}
	sort.Strings(stations)

	fmt.Fprint(output, "{")
	for i, station := range stations {
		if i > 0 {
			fmt.Fprint(output, ", ")
		}
		s := stationStats[station]
		mean := float64(s.Sum) / float64(s.Count) / 10000
		s.Avg = mean
		fmt.Fprintf(output, "%s: {min:%f, max:%f, avg:%f}", station, float64(s.Min)/10000, float64(s.Max)/10000, mean)
	}
	fmt.Fprint(output, "}\n")
	return nil
}
