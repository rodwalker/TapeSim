package main

import (
	"fmt"
	"github.com/rodwalker/TapeSim"
)

func main() {
	fmt.Println("Hello World")
	// t1 := TapeSim.Tape{}
	TapeSim.LoadTapes()
	datasetFiles := TapeSim.GetFileList("jsonFiles/2files.json")
	var files []TapeSim.File
	for _,dsfiles := range datasetFiles{
		files = append(files,dsfiles...)
	}
	TapeSim.WriteFiles(files)
	id:=TapeSim.LocateFile("ds1_1")
	id =TapeSim.LocateFile("ds3_1")
	//read 1 file 
	timeTaken := TapeSim.ReadFiles("ds1_0")
	fmt.Println(id,timeTaken)
}
