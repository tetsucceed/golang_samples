package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func SampleReadFile() {
	file, err := os.Open("README.md")
	if err != nil {
		return
	}
	defer file.Close()

	// getting size of file
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// will not open huge file
	if stat.Size() > 1024*1024*1024 {
		log.Fatal("File is too large to open")
	}
	// buffered reader
	f_reader := bufio.NewReader(file)

	// reading file
	for {
		// reading by line
		line, _, err := f_reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(line))
	}
}
