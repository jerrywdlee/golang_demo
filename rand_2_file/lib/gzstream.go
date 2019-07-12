package lib

import (
	"bufio"
	"compress/gzip"
	"log"
	"os"
)

// GzStream : Gzip writer stream
type GzStream struct {
	f     *os.File
	gf    *gzip.Writer
	fw    *bufio.Writer
	money int
}

// NewGzStream Init GzStream
func NewGzStream(s string) *GzStream {
	f := new(GzStream)
	fi, err := os.OpenFile(s, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Printf("Error in Create\n")
		panic(err)
	}
	f.f = fi
	f.gf = gzip.NewWriter(fi)
	f.fw = bufio.NewWriter(f.gf)
	return f
}

// WriteGZ Write into stream
func (f *GzStream) WriteGZ(s string) {
	(f.fw).WriteString(s)
}

// CloseGZ Close file stream
func (f *GzStream) CloseGZ() {
	f.fw.Flush()
	// Close the gzip first.
	f.gf.Close()
	f.f.Close()
}
