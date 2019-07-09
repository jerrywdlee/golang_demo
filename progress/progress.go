package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// for i := 0; i < 100; i++ {
	// 	fmt.Printf("\rOn %d/10", i)
	// 	fmt.Printf("\r\t\tOn %d/10", i)
	// 	time.Sleep(10 * time.Millisecond)
	// }
	quits := make([]chan bool, 5)
	for j := 0; j < 5; j++ {
		quit := make(chan bool)
		quits[j] = quit
		go ShowProgressBar(j, quit)
	}
	for _, quit := range quits {
		<-quit
	}
	defer fmt.Println("\nEnd ")
}

// ShowProgressBar show bar
func ShowProgressBar(index int, quit chan bool) {
	for i := 0; i < 100; i++ {
		tempStr := "\r" + strings.Repeat("\t\t", index) + "Chan%d: %d/100"
		fmt.Printf(tempStr, index, i)
		// fmt.Printf("\r\t\tOn %d/10", i)
		time.Sleep(50 * time.Millisecond)
	}
	quit <- true
}
