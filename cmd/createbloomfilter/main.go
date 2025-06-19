package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	bloom "github.com/kaszkietio/go-to-that-website/internal/bloomfilter"
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
	if len(args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <data-path> <false-positive-rate> <output-path>", args[0])
		os.Exit(1)
	}

	fp, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing false positive rate: %v", err)
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

	n := len(lines)
	m := calculateBloomFilterSize(n, fp)                  // size of the bloom filter
	k := calculateNumberOfHashFunctionsForBloomFilter(fp) // number of hash functions

	fmt.Printf("Number of lines: %d\n", n)
	fmt.Printf("Bloom filter size (m): %d\n", m)
	fmt.Printf("Number of hash functions (k): %d\n", k)

	filter := bloom.NewBloomFilter(m, k)
	for _, line := range lines {
		filter.Add(line)
	}

	err = filter.Save(args[3])
	if err != nil {
		fmt.Printf("Error saving bloom filter: %v", err)
		os.Exit(1)
	}
	fmt.Printf("SUCCESS: Bloom filter created with size %d and %d hash functions.\n", m, k)
}
