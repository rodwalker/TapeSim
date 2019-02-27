package TapeSim

import ( 
	"fmt"
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
	fmt.Println(tf)
	// always going forward
	seekTime := float64(tf.startMB - t.position) / float64(t.seekRate)
	readTime := float64(tf.endMB - tf.startMB) / float64(t.readRate)
	fmt.Println(seekTime,readTime,t.position)
	// set new position
	t.position = tf.endMB
	return seekTime+readTime
}

func (t *SimpleTape) readFiles(f []File) float64 {
	fmt.Println(len(f))
	return 0.0
}

// order requested files by position
func (t *SimpleTape) orderRequestedFiles() {

}