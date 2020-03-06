package lesson7

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func TimeServer() {

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)

		for {
			reader := bufio.NewReader(conn)
			var buffer bytes.Buffer

			line, prefix, err := reader.ReadLine()
			if err != nil {
				fmt.Println(err)
				break
			}
			buffer.Write(line)
			if !prefix {
				fmt.Println(prefix)
			}

			fmt.Println(buffer.String())

			if string(buffer.String()) == "exit" {
				_, err := io.WriteString(conn, "Bye")
				// just in case to not miss BYE
				time.Sleep(time.Second * 10)
				if err != nil {
					fmt.Println(err)
				}
				err = conn.Close()
				if err != nil {
					fmt.Println(err)
				}
				break
			}
		}
		return
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
