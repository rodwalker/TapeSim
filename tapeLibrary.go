package TapeSim

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var mountTime = 200
var numberDrives = 2
var tapes []SimpleTape

func LoadTapes() {
	for id := 0; id <= 9; id++ {
		//	t = Tape{id: id, capacity: 10, readSpeed: 360, windSpeed: 360,	position:0,mounted: true, catalog: make(map[string]TapeFile)}
		//	t.tapeInfo()
		tapes = append(tapes, SimpleTape{Tape{id: id, capacity: 10, readRate: 360, seekRate: 360, position: 0,
			mounted: true, catalog: make(map[string]TapeFile)}})
	}
}

// write files in order
func WriteFiles(files []File) {
	// set first tape, then next when full
	i := 0
	t := tapes[i]
	// order is not enforced!
	for _, f := range files {
		// first check it fits
		if t.position+f.size > 1000*1000*t.capacity {
			fmt.Println("Tape full. Rewind and mount next.")
			t.tapeInfo()
			t.position = 0
			i++
			t = tapes[i]
		}
		t.writeFile(f)
	}
	t.tapeInfo()
	// rewind last tape
	t.position = 0

}

/* func LocateFile(f string) int {
	for _,t := range tapes {
		if len(t.gotFiles([File{fileName:f}])) == 1 {
			fmt.Printf("Got file on %d \n",t.id)
			return t.id
		}
	}
	fmt.Printf("%s not found on any of %d tapes\n",f,len(tapes))
	return 0
}  */

// Return to read list of files
func ReadFiles(files []File) float64 {
	// need to look up and group by tape. No need to order
	// as tape optimizes this. First tape for now.
	// give full list to each tape, and get back the ones this tape has
	var onTape [][]File
	for _, t := range tapes {
		onTape = append(onTape, t.gotFiles(files))
	}
	timeTaken := 0.0
	// to receive timeTaken from each tape
	ch := make(chan float64)
	var times []float64
	for i, t := range tapes {
		if len(onTape[i]) > 0 {
			fmt.Println("Read tape: ", i)
			files := onTape[i]
			go func(files []File, t SimpleTape) {
				t.readFiles(files, ch)
			}(files, t)
			times = append(times, <-ch)
		}
	}

	fmt.Println(times)
	for _, t := range times {
		timeTaken += t
	}
	return timeTaken
}

// load list of files from json
func GetFileList(file string) map[string][]File {

	// and store datasets for read tests. Not advisable for dump of multi-PB system!
	var files []File
	datasetFiles := make(map[string][]File)

	type Dataset struct {
		Dsname   string `json:"dsname"`
		FileSize int    `json:"fileSize"`
		NFiles   int    `json:"NFiles"`
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
	for i := 0; i < len(datasets.Datasets); i++ {
		ds := datasets.Datasets[i].Dsname
		size := datasets.Datasets[i].FileSize
		for j := 0; j < datasets.Datasets[i].NFiles; j++ {
			fn := ds + "_" + strconv.Itoa(j)
			files = append(files, File{fn, size, ds})
			datasetFiles[ds] = append(datasetFiles[ds], File{fn, size, ds})
		}
	}
	var newfiles []File
	for _, dsfiles := range datasetFiles {
		newfiles = append(newfiles, dsfiles...)
	}
	fmt.Println(len(newfiles))
	return datasetFiles
}
