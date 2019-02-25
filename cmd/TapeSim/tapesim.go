package main

import (
	"fmt"

	"github.com/rodwalker/TapeSim"
)

func main() {
	fmt.Println("Hello World")
	// t1 := TapeSim.Tape{}
	TapeSim.LoadTapes()
	files := TapeSim.GetFileList("theFiles.json")
	TapeSim.WriteFiles(files)


	//populateTapes
}
