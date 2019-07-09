package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// string型スライスを作成
	var capacity = 10
	s1 := make([]int32, 0, capacity)
	fmt.Println("len=", len(s1))
	fmt.Println("cap=", cap(s1))

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 20; i++ {
		var num = rand.Int31n(50)
		if len(s1) < capacity {
			s1 = append(s1, num) // 末尾追加
			continue
		}
		sort.Slice(s1, func(prev, curr int) bool {
			return s1[prev] > s1[curr]
		})

		if s1[len(s1)-1] > num {
			s1 = s1[1:]          // 先頭削除
			s1 = append(s1, num) // 末尾追加
		} else if s1[0] < num {
			s1 = s1[:len(s1)-1]                        // 末尾削除
			s1, s1[0] = append(s1[:1], s1[0:]...), num // 先頭追加
		}
		// fmt.Printf("len\tcap\tpointer\n")
		// fmt.Printf("%d\t%d\t%p\n", len(s1), cap(s1), s1)
		// fmt.Println("s1", s1)
		// fmt.Printf("\n")
	}

	// fmt.Println("s1", s1)
	// fmt.Println("len=", len(s1))
	// fmt.Println("cap=", cap(s1))
	fmt.Printf("len\tcap\tpointer\n")
	fmt.Printf("%d\t%d\t%p\n", len(s1), cap(s1), s1)
	fmt.Println("s1", s1)
}
