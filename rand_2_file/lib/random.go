package lib

import (
	"math"
	"math/rand"
	"time"
)

// RandomStream : rand stream
func RandomStream(label string, size int) <-chan float64 {
	rand.Seed(time.Now().UnixNano())
	output := make(chan float64)
	go func() {
		tSize := 100 * size
		randAry := make([]float64, 0, 200)
		cnt := 0
		for i := 0; i < tSize; i++ {
			mag, val := genRandNum(label)
			for j := 0; j < mag; j++ {
				randAry = append(randAry, val)
			}
			cnt += mag
			i += (mag - 1)
			if cnt >= 100 || i == tSize-1 {
				rand.Shuffle(len(randAry), func(i, j int) {
					randAry[i], randAry[j] = randAry[j], randAry[i]
				})
				for _, num := range randAry {
					output <- num
				}
				randAry = make([]float64, 0, 200)
				cnt = 0
			}
		}
		close(output)
	}()
	return output
}

// Exponential distribution
func expPdf(x, lambda float64) float64 {
	return lambda * math.Exp(-lambda*x)
}

// Normal distribution
// TODO: try https://blog.csdn.net/BertDai/article/details/78231609
func normPdf(x, mean, sigma float64) float64 {
	num1 := (1 / (math.Sqrt(2*math.Pi) * sigma))
	num2 := (math.Exp(-math.Pow(x-mean, 2) / (2 * math.Pow(sigma, 2))))
	return num1 * num2
}

func genRandNum(label string) (int, float64) {
	switch label {
	case "exp":
		for {
			mag := rand.Float64() * 100.0
			val := rand.Float64() * 10.0
			if mag <= expPdf(val, 1) {
				return int(mag+.5) + 1, val
			}
		}
	case "normal":
		for {
			mag := rand.Float64() * 100.0
			val := rand.Float64() * 10.0
			if mag <= normPdf(val, 5, 1) {
				return int(mag+.5) + 1, val
			}
		}
	}
	return 0, 0.0
}
