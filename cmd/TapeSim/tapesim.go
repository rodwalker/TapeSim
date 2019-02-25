package main

import (
	"fmt"

	"github.com/rodwalker/TapeSim"
)

func main() {
	fmt.Println("Hello World")
	// t1 := TapeSim.Tape{}
	TapeSim.LoadTapes()
	files := TapeSim.GetFileList("../TapeSim/jsonFiles/2files.json")
	TapeSim.WriteFiles(files)


	//populateTapes
}
