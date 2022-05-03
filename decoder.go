package main

type UV struct {
	X, Y uint16
}

type Normal struct {
	X, Y, Z int16
}

type Vertex struct {
	X, Y, Z float32
}

type Mesh struct {
	MeshStartOffset   int   // 000010 00000000 00000000 00000000 00000000 01010001 00000000
	StripCountersSize int   // How many strip counters are there
	StripCounters     []int // The size of each strip
	UVSize            int   // The amount of UV points
	UVs               []UV
	NormalsSize       int
	Normals           []Normal
	VertsSize         int
	Verts             []Vertex
}

func DataToMesh(data []byte, meshTolerance int) (Mesh, error) {
	mesh := Mesh{}

	// Find the MeshStartOffset
	{
		meshStartOffsetPattern := []byte{
			0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80,
		}
		pos, err := FindPattern(data, meshTolerance, meshStartOffsetPattern)
		if err != nil {
			return Mesh{}, err
		}
		mesh.MeshStartOffset = pos
	}

	// Find StripCountersSize
	{
		stripsAmount := data[mesh.MeshStartOffset+47]
		mesh.StripCountersSize = int(stripsAmount)
	}

	// Find StripCounters
	{
		for i := 0; i < mesh.StripCountersSize; i++ {
			mesh.StripCounters = append(mesh.StripCounters, int(data[mesh.MeshStartOffset+47+((i+1)*16)]))
		}
	}

	// UNKNOWN UVSIZE
	// UNKNOWN UVs
	// UNKNOWN NormalsSize
	// UNKNOWN Normals

	// Find VertsSize & Verts
	{
		vertPattern := []byte{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x3F, 0x00, 0x00, 0x00, 0x20, 0x40, 0x40, 0x40, 0x40,
		}
		pos, err := FindPattern(data[mesh.MeshStartOffset:], 0, vertPattern)
		if err != nil {
			return Mesh{}, err
		}
		pos += mesh.MeshStartOffset // That way its relative to the file start instead of MeshStartOffset
		mesh.VertsSize = int(data[pos+30])

		firstVertexIndex := pos + 30
		rawVerts := data[firstVertexIndex : firstVertexIndex+mesh.VertsSize*12]
		mesh.Verts = RawVertsToVertexArray(rawVerts)
	}

	return mesh, nil
}
