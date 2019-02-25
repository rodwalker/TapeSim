package TapeSim

import "fmt"

type Tape struct {
	id int
	// size in GB
	capacity int
	// read speed MB/s
	readSpeed int
	// wind speed MB/s
	windSpeed int
	// the files
//	catalog []TapeFile
	catalog map[string]TapeFile
	position int
	mounted bool
}

func (t *Tape) addFile(f File) {
	// find the last MB of the last file
	// OR assume we do not mix read and write, so position is the end 

	tf := TapeFile{t.position ,t.position+f.size}
//	t.catalog = append(t.catalog, tf)
	t.catalog[f.fileName] = tf
	t.position = tf.endMB
}

func (t *Tape) readFile(f File){
	//locate,wind,read
	tf := t.catalog[f.fileName]
	fmt.Println(tf)
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

func AddTape() {
	t1 := Tape{capacity: 10, readSpeed: 400, windSpeed: 600, position:0, mounted: true}
	t1.catalog = make(map[string]TapeFile)
	fmt.Println(t1)
	t1.addFile(File{"f1",2000,"ds1"})
	t1.addFile(File{"f2",2000,"ds1"})
	t1.tapeInfo()
}

func (t Tape) tapeInfo(){	
	fmt.Printf("Id: %d No. of files: %d ",t.id,len(t.catalog))
	fmt.Printf("Tape position: %d MB\n",t.position)
}