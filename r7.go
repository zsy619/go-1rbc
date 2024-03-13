package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
)

func Run7(inputPath string, output io.Writer) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	type item struct {
		key  []byte
		stat *StatsInt32
	}
	const numBuckets = 1 << 17        // number of hash buckets (power of 2)
	items := make([]item, numBuckets) // hash buckets, linearly probed
	size := 0                         // number of active items in items slice
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
			const (
				// FNV-1 64-bit constants from hash/fnv.
				offset64 = 14695981039346656037
				prime64  = 1099511628211
			)
			// Hash the station name and look for ';'.
			var station, after []byte
			hash := uint64(offset64)
			i := 0
			for ; i < len(chunk); i++ {
				c := chunk[i]
				if c == ';' {
					station = chunk[:i]
					after = chunk[i+1:]
					break
				}
				hash ^= uint64(c) // FNV-1a is XOR then *
				hash *= prime64
			}
			if i == len(chunk) {
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

			hashIndex := int(hash & uint64(numBuckets-1))
			for {
				if items[hashIndex].key == nil {
					// Found empty slot, add new item (copying key).
					key := make([]byte, len(station))
					copy(key, station)
					items[hashIndex] = item{
						key: key,
						stat: &StatsInt32{
							Min:   temp,
							Max:   temp,
							Sum:   int64(temp),
							Count: 1,
						},
					}
					size++
					if size > numBuckets/2 {
						panic("too many items in hash table")
					}
					break
				}
				if bytes.Equal(items[hashIndex].key, station) {
					// Found matching slot, add to existing stats.
					s := items[hashIndex].stat
					s.Min = min(s.Min, temp)
					s.Max = max(s.Max, temp)
					s.Sum += int64(temp)
					s.Count++
					break
				}
				// Slot already holds another key, try next slot (linear probe).
				hashIndex++
				if hashIndex >= numBuckets {
					hashIndex = 0
				}
			}
		}

		readStart = copy(buf, remaining)
	}

	// Output stats.
	outItems := make([]item, 0, size)
	for _, item := range items {
		if item.key == nil {
			continue
		}
		outItems = append(outItems, item)
	}
	sort.Slice(outItems, func(i, j int) bool {
		return string(outItems[i].key) < string(outItems[j].key)
	})

	fmt.Fprint(output, "{")
	for i, station := range outItems {
		if i > 0 {
			fmt.Fprint(output, ", ")
		}
		s := station.stat
		mean := float64(s.Sum) / float64(s.Count) / 10000
		s.Avg = mean
		fmt.Fprintf(output, "%s: {min:%f, max:%f, avg:%f}", string(station.key), float64(s.Min)/10000, float64(s.Max)/10000, mean)
	}
	fmt.Fprint(output, "}\n")
	return nil
}
