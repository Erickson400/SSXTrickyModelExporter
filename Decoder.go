package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadMpf(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic("Failed to open file")
	}
	defer f.Close()

}

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
