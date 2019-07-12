package main

import (
	"strconv"
	"time"

	. "./lib"
)

func main() {
	// label := "exp"
	label := "normal"
	numHect := 50 // 50 * 100 = 5000

	fileName := label + "_"
	fileName += strconv.FormatInt(time.Now().Unix(), 10)
	fileName += ".gz"

	gzs := NewGzStream(fileName)
	// output := RandomStream("normal", 50)
	output := RandomStream(label, numHect)
	for randNum := range output {
		str := strconv.FormatFloat(randNum, 'e', 6, 64)
		str += "\n"
		gzs.WriteGZ(str)
	}
	defer gzs.CloseGZ()
}
