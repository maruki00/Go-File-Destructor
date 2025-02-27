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
		file, err := os.OpenFile(filePath, os.O_RDWR, 0777)
		if err != nil {
			continue
		}
		file.Write(getRandomData())
	}
	os.Remove(filePath)
}

func run(basePath string, writeTimes int) {
	dir, err := os.ReadDir(basePath)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return
	}
	counter := 0
	for _, file := range dir {
		if file.IsDir() {
			run(path.Join(basePath, file.Name()), writeTimes)
			os.Remove(path.Join(basePath, file.Name()))
			continue
		}
		distroy(path.Join(basePath, file.Name()), writeTimes)
		counter++
		print("\r\rfiles : ", counter)
	}
}

func main() {
	var basePath string
	var timeToWrite int
	flag.StringVar(&basePath, "path", "", "specify the path of files")
	flag.IntVar(&timeToWrite, "times", 1000, "specify how many times you want to write on file")
	flag.Parse()
	fmt.Println("Process Started")
	run(basePath, timeToWrite)
	fmt.Println("Finished")

}
