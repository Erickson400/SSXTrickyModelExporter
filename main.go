package main

import (
	"fmt"
	"os"
)

func main() {

	// objexporter
	// arg1: [mpf file]
	// arg2: [obj destination path]
	// objexporter C:\Users\munoz\Documents\GitHub\SSXTrickyModelExporter\resources\board.mpf C:\Users\munoz\Documents\GitHub\SSXTrickyModelExporter\resources

	// objexporter

	// Help
	if len(os.Args) == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "h" || os.Args[1] == "-help" || os.Args[1] == "help" {
			fmt.Printf(" Command: objexporter [mpf file] [.obj destination path] \n ")
			return
		}
	}
	if len(os.Args) < 3 {
		fmt.Println("You must give 2 arguments: objexporter [.mpf file] [.obj destination path]")
		return
	}

	// Check if file/path exist
	_, err := os.Stat(filepath.ToSlash(os.Args[1]))
	if os.IsNotExist(err) {
		fmt.Println("the .mpf file directory is not valid: ", os.Args[1])
		return
	}
	_, err = os.Stat(filepath.ToSlash(os.Args[2]))
	if os.IsNotExist(err) {
		fmt.Println(".obj destination path does not exist", os.Args[2])
		return
	}

	// If all is gucci then procced.
	data, err := ReadMpf(filepath.ToSlash(os.Args[1]))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read meshes until an error occurs
	meshes := []Mesh{}
	for i := 0; ; i++ {
		myMesh, err := DataToMesh(data, i)
		if err != nil {
			break
		}
		meshes = append(meshes, myMesh)
	}

	err = ModelFromMeshArray(filepath.ToSlash(os.Args[2]), meshes)
	if err != nil {
		fmt.Println("ERROR: Failed to create model: ", err)
	}
	fmt.Printf("SUCCESS: Succesfully made mesh.obj in '%v'\n", filepath.ToSlash(os.Args[2]))
}
