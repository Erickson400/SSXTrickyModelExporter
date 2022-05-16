package main

import (
	"fmt"
	"os"
	"strconv"
)

type Face struct {
	V1, V2, V3 int
}

func ModelFromMeshArray(filepath string, meshes []Mesh) error {
	f, err := os.Create(filepath + "/mesh.obj")
	if err != nil {
		return fmt.Errorf("Failed to create file")
	}
	defer f.Close()

	for i, m := range meshes {
		fmt.Fprintf(f, "o mesh%v\n", strconv.FormatInt(int64(i), 10))
		splits := RawStripCountersToIncremental(m.StripCounters)
		verts, faces := MakeSplitStripModel(splits, m.Verts)
		content := MakeModel(verts, faces)
		fmt.Fprintln(f, content)
	}
	return nil
}

// func ModelFromMesh(filename string, mesh Mesh) {
// 	splits := RawStripCountersToIncremental(mesh.StripCounters)
// 	MakeSplitStripModel(filename, splits, mesh.Verts)
// }

func MakeSplitStripModel(splits []int, vertices []Vertex) ([]Vertex, []Face) {
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
	return vertices, finalFs
}

func MakeModel(vertices []Vertex, faces []Face) (out string) {
	for _, v := range vertices {
		out += fmt.Sprintln("v", v.X, v.Y, v.Z)
	}
	for _, v := range faces {
		out += fmt.Sprintln("f", v.V1+1, v.V2+1, v.V3+1)
	}
	return
}
