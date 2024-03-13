package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func Run8(inputPath string, output io.Writer) error {
	// 获取cpu核心数
	maxGoroutines := runtime.NumCPU()
	// 分割文件
	parts, err := splitFile(inputPath, maxGoroutines)
	if err != nil {
		return nil
	}
	// 创建一个channel
	resultsCh := make(chan map[string]Stats)
	// 读取文件，并发处理，构建生产者消费者模型
	for _, part := range parts {
		go r8ProgressPart(inputPath, part.offset, part.size, resultsCh)
	}

	// 创建一个map，用于存储结果
	totals := make(map[string]Stats)
	// 从channel中读取结果
	partLen := len(parts)
	for i := 0; i < partLen; i++ {
		results := <-resultsCh
		for station, stats := range results {
			ts, ok := totals[station]
			if !ok {
				totals[station] = Stats{
					Min:   stats.Min,
					Max:   stats.Max,
					Sum:   stats.Sum,
					Count: stats.Count,
				}
				continue
			}
			ts.Min = min(ts.Min, stats.Min)
			ts.Max = max(ts.Max, stats.Max)
			ts.Sum += stats.Sum
			ts.Count += stats.Count
			totals[station] = ts
		}
	}

	stations := make([]string, 0, len(totals))
	for station := range totals {
		stations = append(stations, station)
	}
	sort.Strings(stations)

	fmt.Fprint(output, "{")
	for i, station := range stations {
		if i > 0 {
			fmt.Fprint(output, ", ")
		}
		s := totals[station]
		mean := s.Sum / float64(s.Count)
		s.Avg = mean
		fmt.Fprintf(output, "%s: {min:%f, max:%f, avg:%f}", station, s.Min, s.Max, mean)
	}
	fmt.Fprint(output, "}\n")
	return nil
}

func r8ProgressPart(inputPath string, fileOffset, fileSize int64, resultsCh chan map[string]Stats) {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Seek(fileOffset, io.SeekStart)
	if err != nil {
		panic(err)
	}
	f := io.LimitedReader{R: file, N: fileSize}

	stationStats := make(map[string]Stats)

	scanner := bufio.NewScanner(&f)
	for scanner.Scan() {
		line := scanner.Text()
		station, tempStr, hasSemi := strings.Cut(line, ";")
		if !hasSemi {
			continue
		}

		temp, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			panic(err)
		}
		s, ok := stationStats[station]
		if !ok {
			s.Min = temp
			s.Max = temp
			s.Sum = temp
			s.Count = 1
		} else {
			s.Min = min(s.Min, temp)
			s.Max = max(s.Max, temp)
			s.Sum += temp
			s.Count++
		}
		stationStats[station] = s
	}
	resultsCh <- stationStats
}

type part struct {
	offset, size int64
}

func splitFile(inputPath string, numParts int) ([]part, error) {
	const maxLineLength = 100

	// Open the input file
	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	// Get the file stats
	st, err := f.Stat()
	if err != nil {
		return nil, err
	}
	// Get the file size
	size := st.Size()
	// Calculate the split size
	splitSize := size / int64(numParts)
	// Create a buffer to read the file
	buf := make([]byte, maxLineLength)
	// Create a slice to store the parts
	parts := make([]part, 0, numParts)
	// Set the offset to 0
	offset := int64(0)
	// Loop through the file
	for offset < size {
		// Calculate the seek offset
		seekOffset := max(offset+splitSize-maxLineLength, 0)
		// If the seek offset is greater than the file size, break
		if seekOffset > size {
			break
		}
		// Seek to the seek offset
		_, err := f.Seek(seekOffset, io.SeekStart)
		if err != nil {
			return nil, err
		}
		// Read the buffer
		n, _ := io.ReadFull(f, buf)
		// Create a chunk of the buffer
		chunk := buf[:n]
		// Find the last new line
		newLine := bytes.LastIndexByte(chunk, '\n')
		// If no new line is found, return an error
		if newLine < 0 {
			return nil, fmt.Errorf("newline not found at offset %d", offset+splitSize-maxLineLength)
		}
		// Calculate the remaining bytes
		remaining := len(chunk) - newLine - 1
		// Calculate the next offset
		nextOffset := seekOffset + int64(len(chunk)) - int64(remaining)
		// Append the part to the parts slice
		parts = append(parts, part{offset: offset, size: nextOffset - offset})
		// Set the offset to the next offset
		offset = nextOffset
	}
	return parts, nil
}
