package TapeSim

import "fmt"

// Anything that can store and retrieve a file
type FileStore interface{
	writeFile(f File)
	readFile(f File)
	gotFile(f File) bool
}

type Tape struct {
	id int
	// size in GB
	capacity int
	// read speed MB/s
	readRate int
	// wind speed MB/s
	seekRate int
	// the files
//	catalog []TapeFile
	catalog map[string]TapeFile
	position int
	mounted bool
}

func (t Tape) gotFile(f File) bool{
	if _,ok := t.catalog[f.fileName];ok{
		return true
	}
	return false
}

type File struct {
	fileName string
	size     int // MB
	dataset  string
}

type TapeFile struct {
	startMB  int
	endMB    int
}

func (t Tape) tapeInfo(){	
	fmt.Printf("Id: %d No. of files: %d ",t.id,len(t.catalog))
	fmt.Printf("Tape position: %d MB\n",t.position)
}