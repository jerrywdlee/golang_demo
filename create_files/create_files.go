package main

// Main Pkg
import (
	"bufio"
	"compress/gzip"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

// GzipWriteStream Gzip stream writer, See https://gist.github.com/mchirico/6147687
type GzipWriteStream struct {
	f     *os.File
	gf    *gzip.Writer
	fw    *bufio.Writer
	money int
}

// NewGzipWriteStream Init GzipWriteStream
func NewGzipWriteStream(s string) *GzipWriteStream {
	f := new(GzipWriteStream)
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
func (f *GzipWriteStream) WriteGZ(s string) {
	(f.fw).WriteString(s)
}

// CloseGZ Close file stream
func (f *GzipWriteStream) CloseGZ() {
	f.fw.Flush()
	// Close the gzip first.
	f.gf.Close()
	f.f.Close()
}

func main() {
	dirPath := filepath.Join(".", "data")
	os.MkdirAll(dirPath, os.ModePerm)
	filePath := filepath.Join(dirPath, "test.txt")
	//os.O_RDWRを渡していると、同時に読み込みも可能
	streamMode := os.O_WRONLY | os.O_CREATE | os.O_APPEND
	file, err := os.OpenFile(filePath, streamMode, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	rand.Seed(time.Now().UnixNano())
	fmt.Fprintln(file, rand.Int31())

	gzf := NewGzipWriteStream("data/Append.gz")
	gzf.WriteGZ("AAAAAAA\n")
	defer gzf.CloseGZ()
}
