package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func CopyFile(name1, name2 string) {
	f_origin, err := os.Open(name1)
	if err != nil {
		log.Fatal("ERROR csv")
	}
	defer f_origin.Close()

	f_new, err := os.OpenFile(name2, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f_new.Close()

	f_reader := bufio.NewReader(f_origin)
	f_writer := bufio.NewWriter(f_new)

	for {
		// reading by line
		// TODO: rewrite with scanner
		line, _, err := f_reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		line_string := fmt.Sprintf("%s\n", string(line))

		_, err = f_writer.Write([]byte(line_string))
		if err != nil {
			log.Fatal(err)
		}
		_ = f_writer.Flush()
	}

}
