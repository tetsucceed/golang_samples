package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadSampleCsv() {

	f, err := os.Open("example.csv")
	if err != nil {
		log.Fatal("ERROR csv")
	}
	defer f.Close()

	f_reader := bufio.NewReader(f)

	for {
		csv_reader := csv.NewReader(f_reader)
		csv_reader.Comma = ';'
		csv_reader.Comment = '#'

		record, err := csv_reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("csv FATAL")
		}
		fmt.Println(record)
	}
}

func WriteSampleCsv(name string) {

	// init
	ex_csv := make([][]string, 10)

	for i := range ex_csv {
		ex_csv[i] = make([]string, 10)
	}

	for i := 0; i < 10; i++ {
		for z := 0; z < 10; z++ {
			ex_csv[i][z] = fmt.Sprintf("name%d", (i + z))
			fmt.Println(ex_csv[i][z])
		}
	}

	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f_writer := bufio.NewWriter(f)
	csv_writer := csv.NewWriter(f_writer)

	csv_writer.Comma = ';'

	err = csv_writer.WriteAll(ex_csv)
	if err != nil {
		log.Fatal(err)
	}
}
