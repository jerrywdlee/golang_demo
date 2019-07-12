package main

import (
	"strconv"

	. "./lib"
)

func main() {
	gzs := NewGzStream("test.gz")
	// output := RandomStream("normal", 50)
	output := RandomStream("exp", 50)
	for randNum := range output {
		str := strconv.FormatFloat(randNum, 'e', 6, 64)
		str += "\n"
		gzs.WriteGZ(str)
	}
	defer gzs.CloseGZ()
}
