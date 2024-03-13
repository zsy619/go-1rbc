package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func Run4(inputPath string, output io.Writer) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	stationStats := make(map[string]*StatsInt32)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		station, tempStr, hasSemi := strings.Cut(line, ";")
		if !hasSemi {
			continue
		}

		temp, _ := ToInt32(tempStr)

		s := stationStats[station]
		if s == nil {
			stationStats[station] = &StatsInt32{
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
