// From: https://blog.csdn.net/wangyangzhizhou/article/details/80088490

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	length := 200
	resultSlice := make([]float64, 0, length*2)
	for i := 0; i < length; i++ {
		// numAは係数、numBは値
		numA, numB := randExp(20.0, 1.0, func(a float64) float64 {
			return exponentialPdf(a, 1)
		})
		intA := int(numA + .5) // 四捨五入
		for j := 0; j < (intA + 1); j++ {
			// println(numB)
			resultSlice = append(resultSlice, numB)
		}
	}
	rand.Shuffle(len(resultSlice), func(i, j int) {
		resultSlice[i], resultSlice[j] = resultSlice[j], resultSlice[i]
	})
	fmt.Println("Res", resultSlice[:length])
}

// 指数分布
func exponentialPdf(x, lambda float64) float64 {
	return lambda * math.Exp(-lambda*x)
}

func randExp(rangeA, rangeB float64, callback func(float64) float64) (float64, float64) {
	for {
		numA := rand.Float64() * rangeA
		numB := rand.Float64() * rangeB
		if numB <= callback(numA) {
			return numA, numB
		}
	}
}
