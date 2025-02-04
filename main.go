package main

import (
	"math"
	"math/rand"
)

func getRandomData() []byte {
	nRundom := rand.Intn(8096) + 4048
	data := make([]byte, nRundom)
	for i := range nRundom {
		data[i] = byte(rand.Intn(255) + 32)
	}
	return data
}
func distroy(filePath string, writeTimes int) {
	for range writeTimes {

	}
}

func main() {

}
