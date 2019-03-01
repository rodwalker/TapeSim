package main

import (
	"fmt"
	//"math/rand"
	"github.com/rodwalker/TapeSim"
)

func main() {
	fmt.Println("Loading tapes")
	TapeSim.LoadTapes()
	datasetFiles := TapeSim.GetFileList("jsonFiles/2files.json")
	var files []TapeSim.File
	for _, dsfiles := range datasetFiles {
		files = append(files, dsfiles...)
	}
	//rand.Shuffle(len(files), func(i, j int) { files[i], files[j] = files[j], files[i] })
	TapeSim.WriteFiles(files)

	//read 1 dataset
	fmt.Println(len(datasetFiles["ds1"]))
	timeTaken := TapeSim.ReadFiles(datasetFiles["ds1"][0:])
	fmt.Println(timeTaken)
}
