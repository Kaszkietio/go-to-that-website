package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	lala := [202]bool{}
	var buffer []byte = nil
	for _, xd := range lala {
		buffer2, err := binary.Append(buffer, binary.NativeEndian, xd)
		if err != nil {
			fmt.Printf("ALE JAJCA TY\n")
			os.Exit(1)
		}
		buffer = buffer2
	}
	fmt.Printf("%d\n", len(buffer))
}
