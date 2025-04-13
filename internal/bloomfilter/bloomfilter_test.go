package bloomfilter_test

import (
	"fmt"
	"testing"

	bloom "github.com/kaszkietio/GoToThatWebsite/internal/bloomfilter"
)

func TestBloomFilter(t *testing.T) {
	bf := bloom.NewBloomFilter(100, 10)
	s := "Jasna dupa"
	bf.Add(s)
	fmt.Printf("%s in bloom filter?: %v\n", s, bf.Contains(s))
	fmt.Println(bf.GetIndexes(s))
	s = "Dupa"
	fmt.Printf("%s in bloom filter?: %f\n", s, bf.Probability(s))
}

// TODO: add tests
