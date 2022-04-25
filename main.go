package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Vector struct {
	x, y, z float32
}

type Face struct {
	v1, v2, v3 int
}

func main() {
	strip1 := []Vector{
		{1.0, -1.0, 1.0},
		{1.0, -1.0, -1.0},
		{1.0, 1.0, 1.0},
		{1.0, 1.0, -1.0},
		{-1.0, 1.0, 1.0},
		{-1.0, 1.0, -1.0},
		{-1.0, -1.0, 1.0},
		{-1.0, -1.0, -1.0},
	}

	strip2 := []Vector{
		{-1.0, 1.0, 1.0},
		{1.0, 1.0, 1.0},
		{-1.0, -1.0, 1.0},
		{1.0, -1.0, 1.0},
		{-1.0, -1.0, -1.0},
		{1.0, -1.0, -1.0},
		{-1.0, 1.0, -1.0},
		{1.0, 1.0, -1.0},
	}

	MakeStripModel(strip1, strip2)
}

func MakeStripModel(strip ...[]Vector) {
	// Check every vartice slice if they are greater than 3 len
	for _, v := range strip {
		if len(v) < 3 {
			panic("ERROR: There must at least 3 vertices in the strip")
		}
	}

	finalVecs := []Vector{}
	fs := []Face{}

	for i, v := range strip {
		// Put all the vertices in one Slice
		finalVecs = append(finalVecs, v...)

		// For each strip position
		for j := range v {
			// Skip if position is too small for a face
			if j < 2 {
				continue
			}

			// Make face
			h := j + (i * len(v))
			h++ // first vertex must be 1
			fs = append(fs, Face{h, h - 1, h - 2})
		}
	}

	MakeModel(finalVecs, fs)
}

func MakeModel(vertices []Vector, faces []Face) {
	f, err := os.Create("Model.obj")
	if err != nil {
		panic("Failed to create file")
	}
	defer f.Close()

	for _, v := range vertices {
		fmt.Fprintln(f, "v", v.x, v.y, v.z)
	}
	for _, v := range faces {
		fmt.Fprintln(f, "f", v.v1, v.v2, v.v3)
	}

	content, err := ioutil.ReadFile("Model.obj")
	fmt.Println(string(content))
	fmt.Println("SUCCESS: Succesfully made the Model")
}
