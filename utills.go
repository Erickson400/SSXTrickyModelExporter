package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"strings"
	"unsafe"
)

func RawVertsToVertexArray(rawVerts []byte) []Vertex {
	// Raw Verts to []Vertex
	verts := make([]Vertex, 0, len(rawVerts)/12)
	//fmt.Println(cap(verts))
	for i := 2; i < len(rawVerts); i += 12 {
		// Painfully make the values big endian for the damn Float32frombits
		b1, b2, b3 := make([]byte, 4), make([]byte, 4), make([]byte, 4)
		binary.BigEndian.PutUint32(b1, *(*uint32)(unsafe.Pointer(&rawVerts[i])))
		binary.BigEndian.PutUint32(b2, *(*uint32)(unsafe.Pointer(&rawVerts[i+4])))
		binary.BigEndian.PutUint32(b3, *(*uint32)(unsafe.Pointer(&rawVerts[i+8])))
		x := binary.BigEndian.Uint32(b1)
		y := binary.BigEndian.Uint32(b2)
		z := binary.BigEndian.Uint32(b3)
		x1 := math.Float32frombits(x)
		y1 := math.Float32frombits(y)
		z1 := math.Float32frombits(z)

		vec := Vertex{x1, y1, z1}
		verts = append(verts, vec)
	}
	return verts
}

func RawStripCountersToIncremental(rawCounters []int) []int {
	splits := []int{0}
	for _, v := range rawCounters {
		splits = append(splits, splits[len(splits)-1]+v)
	}
	return splits
}

func ReadMpf(filename string) ([]byte, error) {
	if !strings.HasSuffix(filename, ".mpf") {
		filename += ".mpf"
	}

	content, err := os.ReadFile(filename) // the file is inside the local directory
	if err != nil {
		return nil, fmt.Errorf("Failed to read .mpf file")
	}
	return content, nil
}

func FindPattern(data []byte, tolerance int, pattern []byte) (position int, err error) {
	// Iterate through the whole file
	numsCorrect := 0
	for i := range data {
		for {
			if data[i+numsCorrect] == pattern[numsCorrect] {
				numsCorrect++
				if numsCorrect == len(pattern) {
					if tolerance <= 0 {
						return i, nil
					}
					tolerance--
					numsCorrect = 0
					break
				}
				continue
			}
			numsCorrect = 0
			break
		}
	}
	return 0, fmt.Errorf("Could not find pattern")
}

func StringToStruct(filename string, data string) error {
	dataSlice := strings.Split(data, " ")

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Could not create string file %v", filename)
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
