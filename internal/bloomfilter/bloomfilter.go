package bloomfilter

import (
	"encoding/binary"
	"hash/maphash"
	"os"
)

type BloomFilter struct {
	array []bool
	seeds []maphash.Seed
}

func NewBloomFilter(arrayLength int, numHashFunctions int) *BloomFilter {
	seeds := make([]maphash.Seed, numHashFunctions)
	for i := range numHashFunctions {
		seeds[i] = maphash.MakeSeed()
	}

	return &BloomFilter{
		array: make([]bool, arrayLength),
		seeds: seeds,
	}
}

func (bf *BloomFilter) GetIndexes(s string) []uint64 {
	m := uint64(len(bf.array))
	k := len(bf.seeds)
	indexes := make([]uint64, k)
	for i, seed := range bf.seeds {
		idx := maphash.String(seed, s) % m
		indexes[i] = idx
	}
	return indexes
}

func (bf *BloomFilter) Add(s string) {
	indexes := bf.GetIndexes(s)
	for _, idx := range indexes {
		bf.array[idx] = true
	}
}

func (bf *BloomFilter) Contains(s string) bool {
	indexes := bf.GetIndexes(s)
	for _, idx := range indexes {
		if !bf.array[idx] {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) Probability(s string) float32 {
	indexes := bf.GetIndexes(s)
	den := float32(len(indexes))
	num := 0
	for _, idx := range indexes {
		if bf.array[idx] {
			num += 1
		}
	}

	return float32(num) / den
}

func (bf *BloomFilter) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	m := len(bf.array)
	k := len(bf.seeds)

	err = binary.Write(file, binary.NativeEndian, int64(m))
	if err != nil {
		return err
	}

	for _, a := range bf.array {
		err = binary.Write(file, binary.NativeEndian, a)
		if err != nil {
			return err
		}
	}

	err = binary.Write(file, binary.NativeEndian, int64(k))
	if err != nil {
		return nil
	}

	for _, seed := range bf.seeds {
		err = binary.Write(file, binary.NativeEndian, seed)
		if err != nil {
			return err
		}
	}

	return nil
}

func Load(path string) (bf *BloomFilter, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	m := int64(-1)
	err = binary.Read(file, binary.NativeEndian, &m)
	if err != nil {
		return nil, err
	}

	array := make([]bool, m)
	for i := range m {
		err = binary.Read(file, binary.NativeEndian, &array[i])
		if err != nil {
			return nil, err
		}
	}

	k := int64(-1)
	err = binary.Read(file, binary.NativeEndian, &k)
	if err != nil {
		return nil, err
	}

	seeds := make([]maphash.Seed, m)
	for i := range k {
		err = binary.Read(file, binary.NativeEndian, &seeds[i])
		if err != nil {
			return nil, err
		}
	}

	return &BloomFilter{
		array: array,
		seeds: seeds,
	}, nil
}
