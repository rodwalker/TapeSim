package TapeSim

import ( 
	"fmt"
	"math"
)

// A fictitous 1-dimensional tape
type SimpleTape struct {
	Tape
}

func (t *SimpleTape) writeFile(f File) {
	// find the last MB of the last file
	// OR assume we do not mix read and write, so position is the end 

	tf := TapeFile{t.position ,t.position+f.size}
	t.catalog[f.fileName] = tf
	t.position = tf.endMB
}

func (t *SimpleTape) readFile(f File) float64 {
	//locate,seek,read
	tf := t.catalog[f.fileName]
	//fmt.Println(tf)
	// only going forward, cos we can`t find reverse!
	seekTime := math.Abs(float64(tf.startMB - t.position)) / float64(t.seekRate)
	readTime := float64(tf.endMB - tf.startMB) / float64(t.readRate)
	//fmt.Println(seekTime,readTime,t.position)
	// set new position
	t.position = tf.endMB
	return seekTime+readTime
}

func (t *SimpleTape) readFiles(files []File,ch chan float64) {
	fmt.Printf("Will read %d files from tape %d\n",len(files),t.id)
	// assume in order for now
	timeTaken := 0.0
	size := 0
	for _,f := range files{
		timeTaken += t.readFile(f)
		size += f.size
	}
	fmt.Printf("Read %d MB in %d files from tape %d at %f MB/s \n",
		size,len(files),t.id,float64(size)/timeTaken)
	ch <- timeTaken
}

// order requested files by position
func (t *SimpleTape) orderRequestedFiles() {

}