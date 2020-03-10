package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"utils"
)

func main() {
	var lessonNum, _ = strconv.Atoi(os.Getenv("LESSON"))

	if lessonNum == 1 {
		fmt.Println("Lesson 1")
		utils.Isdiv(-11)
		utils.Isdiv3(333)
		utils.Fibb100()
		utils.Primes()
		os.Exit(0)
	}

	if lessonNum == 2 {
		fmt.Println("Lesson 2")

		fifo := utils.NewFifo()

		var vehs1 = utils.Initialize()
		for v := 0; v < len(vehs1); v++ {
			fmt.Println(vehs1[v])
			fifo.Push(&vehs1[v])
		}

		dataJson, _ := json.Marshal(vehs1)
		err := ioutil.WriteFile("Vehicles.json", dataJson, os.FileMode(0644))
		if err != nil {
			log.Fatal(err)
		}

		var vehs2 = utils.Initialize()
		for v := 0; v < len(vehs2); v++ {
			fmt.Println(vehs2[v])
			fifo.Push(&vehs2[v])
		}

		dataJson, _ = json.Marshal(vehs2)
		err = ioutil.WriteFile("Vehicles.json", dataJson, os.FileMode(0644))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(fifo.Pop(), fifo.Pop()) // pop-pop
		os.Exit(0)
	}

	if lessonNum == 4 {
		utils.ChessKnight()
		utils.AddressBook()
		utils.Calculator()
	}

	if lessonNum == 5 {
		utils.SampleReadFile()
		utils.ReadSampleCsv()
		utils.WriteSampleCsv("out.csv")

		flag.Parse()
		if len(flag.Args()) < 3 {
			fmt.Println(flag.Arg(1))
			log.Fatal("usage: cp [fileName] [newFileName]")
		} else if len(flag.Args()) == 3 {
			switch flag.Arg(0) {
			case "cp":
				filename1 := flag.Arg(1)
				filename2 := flag.Arg(2)
				utils.CopyFile(filename1, filename2)
			}
		}
	}

	fmt.Println("Hello World")
}
