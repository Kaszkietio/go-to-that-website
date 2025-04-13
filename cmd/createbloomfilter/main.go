package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	bloom "github.com/kaszkietio/GoToThatWebsite/internal/bloomfilter"
)

func calculateBloomFilterSize(n int, fp float64) int {
	m := math.Ceil(-(float64(n) * math.Log(fp)) / (math.Pow(math.Log(2), 2)))
	return int(m)
}

func calculateNumberOfHashFunctionsForBloomFilter(fp float64) int {
	k := math.Ceil(-math.Log(fp) / math.Log(2))
	return int(k)
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <data-path> <false-positive-rate>", args[0])
		os.Exit(1)
	}

	file, err := os.Open(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 5715)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fp, err := strconv.ParseFloat(args[2], 64)
	if err != nil {

	}
	n := len(lines)
	m := calculateBloomFilterSize(n, fp)                  // size of the bloom filter
	k := calculateNumberOfHashFunctionsForBloomFilter(fp) // number of hash functions

	filter := bloom.NewBloomFilter(m, k)
	for _, line := range lines {
		filter.Add(line)
	}

}
