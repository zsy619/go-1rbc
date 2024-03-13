package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func Run3(inputPath string, output io.Writer) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	stationStats := make(map[string]*Stats)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		station, tempStr, hasSemi := strings.Cut(line, ";")
		if !hasSemi {
			continue
		}

		temp := ToFloat64_1(tempStr)

		s := stationStats[station]
		if s == nil {
			stationStats[station] = &Stats{Min: temp, Max: temp, Sum: temp, Count: 1}
		} else {
			s.Min = min(s.Min, temp)
			s.Max = max(s.Max, temp)
			s.Sum += temp
			s.Count++
		}
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
		mean := s.Sum / float64(s.Count)
		s.Avg = mean
		fmt.Fprintf(output, "%s: {min:%f, max:%f, avg:%f}", station, s.Min, s.Max, mean)
	}
	fmt.Fprint(output, "}\n")
	return nil
}
