package bloomfilter_test

import (
	"testing"

	bloom "github.com/kaszkietio/go-to-that-website/internal/bloomfilter"
)

func TestBloomFilter(t *testing.T) {
	bf := bloom.NewBloomFilter(100, 10)

	bf.Add("Jasna dupa")
	if !bf.Contains("Jasna dupa") {
		t.Errorf("Expected 'Jasna dupa' to be in the bloom filter")
	}

	if bf.Contains("Jasna dupaa") {
		t.Errorf("Expected 'Jasna dupaa' to not be in the bloom filter")
	}

	err := bf.Save("test_bloom_filter.bin")
	if err != nil {
		t.Errorf("Failed to save bloom filter: %v", err)
	}

	bf2, err := bloom.Load("test_bloom_filter.bin")
	if err != nil {
		t.Errorf("Failed to load bloom filter: %v", err)
	}

	if !bf2.Contains("Jasna dupa") {
		t.Errorf("Expected 'Jasna dupa' to be in the bloom filter after loading")
	}
	if bf2.Contains("Jasna dupaa") {
		t.Errorf("Expected 'Jasna dupaa' to be in the bloom filter after loading")
	}
}
