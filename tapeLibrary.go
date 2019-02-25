package TapeSim

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

var mountTime = 200
var numberDrives = 2
var tapes[] Tape

func LoadTapes() {
	for id:=1;id<=10;id++ {
//	t = Tape{id: id, capacity: 10, readSpeed: 400, windSpeed: 600,	position:0,mounted: true, catalog: make(map[string]TapeFile)}
//	t.tapeInfo()
	tapes = append(tapes,Tape{id: id, capacity: 10, readSpeed: 400, windSpeed: 600,	position:0,
		mounted: true, catalog: make(map[string]TapeFile)})
	}
}

// write files in order 
func WriteFiles(files []File){
	// set first tape, then next when full
	i := 0
	t := tapes[i]
	for _,f := range files{
		// first check it fits
		if t.position+f.size > 1000*1000*t.capacity {
			fmt.Println("Tape full. Mounting next.")
			t.tapeInfo()
			i++
			t = tapes[i]
		}
		t.addFile(f)
	}
	t.tapeInfo()
}

// load list of files from json
func GetFileList(file string) []File {
	var files []File

	type Dataset struct {
		Dsname string `json:"dsname"`
		FileSize int `json:"fileSize"`
		NFiles int `json:"NFiles"`
	}
	type Datasets struct {
		Datasets []Dataset `json:"datasets"`
	}

    jsonFile, err := os.Open(file)
    defer jsonFile.Close()
    if err != nil {
        fmt.Println(err.Error())
	}
	
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var datasets Datasets
	json.Unmarshal(byteValue, &datasets)
	//fmt.Println(datasets)
	for i := 0; i<len(datasets.Datasets);i++ {
		ds := datasets.Datasets[i].Dsname
		size := datasets.Datasets[i].FileSize
		for j := 0;j<datasets.Datasets[i].NFiles;j++{		
			fn := ds+"_"+strconv.Itoa(j)
			files = append(files,File{fn,size,ds})
		}
	}
	
    return files
}
