package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path"
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
		print(".")
		file, err := os.OpenFile(filePath, os.O_RDWR, 0777)
		if err != nil {
			continue
		}
		file.Write(getRandomData())
	}
}

func main() {
	var basePath string
	var timeToWrite int
	flag.StringVar(&basePath, "path", "", "specify the path of files")
	flag.IntVar(&timeToWrite, "times", 1000, "specify how many times you want to write on file")
	flag.Parse()
	fmt.Println("Process Started")
	dir, err := os.ReadDir(basePath)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return
	}
	for _, file := range dir {
		distroy(path.Join(basePath, file.Name()), timeToWrite)
	}

	fmt.Println("Finished")
}
