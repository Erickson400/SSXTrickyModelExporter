package main

import (
	"fmt"
	"os"
	"strings"
)

type Vector struct {
	X, Y, Z float32
}

type Face struct {
	V1, V2, V3 int
}

func MakeSplitStripModel(filename string, splits []int, vertices []Vector) {
	// splits are the first vertex of a strip. does not make a face that connects to the vertex behind,
	// vertices are all the vertices the model uses from all strips combined.

	// The data that will be written to the file
	finalFs := make([]Face, 0, 10)

	// The vertex index relative to the current strip. this way it does not create a face when its less than 3.
	// It is reset to 0 when a new strip is reached, and added to when a vertex is iterate.

	IsInsideSplits := func(num int) bool {
		for _, v := range splits {
			if v == num {
				return true
			}
		}
		return false
	}

	AddFace := func(lastindex int) {
		finalFs = append(finalFs, Face{lastindex, lastindex - 1, lastindex - 2})
	}

	localIndex := 0

	// Iterate through vertices
	for globalIndex := range vertices {

		// Check if i is in the strip list
		if IsInsideSplits(globalIndex) {
			localIndex = 1
			continue
		}
		// Check if there are more than 3 vertices to make a face
		if localIndex < 2 {
			localIndex++
			continue
		}

		AddFace(globalIndex)
		localIndex++
	}

	MakeModel(filename, vertices, finalFs)
}

func MakeModel(filename string, vertices []Vector, faces []Face) {
	if !strings.HasSuffix(filename, ".obj") {
		filename += ".obj"
	}

	f, err := os.Create(filename)
	if err != nil {
		panic("Failed to create file")
	}
	defer f.Close()

	for _, v := range vertices {
		fmt.Fprintln(f, "v", v.X, v.Y, v.Z)
	}
	for _, v := range faces {
		fmt.Fprintln(f, "f", v.V1+1, v.V2+1, v.V3+1)
	}

	//content, err := ioutil.ReadFile(filename)
	//fmt.Println(string(content))
	fmt.Printf("SUCCESS: Succesfully made '%v'\n", filename)

}
