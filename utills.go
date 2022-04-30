package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"strings"
)

func StringToStruct(filename string, data string) error {
	dataSlice := strings.Split(data, " ")

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Could not open string file %v", filename)
	}
	defer f.Close()

	fmt.Fprintln(f, "strips := []int {")
	for _, v := range dataSlice {

		str := fmt.Sprintf("\t0X%s,", v)
		fmt.Fprintln(f, str)
	}
	fmt.Fprintln(f, "}")

	return nil
}

func RawStripsToVectors(rawStrips []uint32) []Vector {
	// Raw Strips to float32 vectors
	strips := make([]Vector, 0, 55)
	for i := 0; i < len(rawStrips); i += 3 {
		// Make the values little endian
		b1, b2, b3 := make([]byte, 4), make([]byte, 4), make([]byte, 4)
		binary.LittleEndian.PutUint32(b1, rawStrips[i])
		binary.LittleEndian.PutUint32(b2, rawStrips[i+1])
		binary.LittleEndian.PutUint32(b3, rawStrips[i+2])
		x := binary.BigEndian.Uint32(b1)
		y := binary.BigEndian.Uint32(b2)
		z := binary.BigEndian.Uint32(b3)
		x1 := math.Float32frombits(x)
		y1 := math.Float32frombits(y)
		z1 := math.Float32frombits(z)

		vec := Vector{x1, y1, z1}
		strips = append(strips, vec)
	}
	return strips
}

func RawSplitsToIncremental(rawSplits []int) []int {
	splits := []int{0}
	for _, v := range rawSplits {
		splits = append(splits, splits[len(splits)-1]+v)
	}
	return splits
}
