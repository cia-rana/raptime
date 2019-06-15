package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
)

var bufSize = (int)(*flag.Uint64("size", 2147483648, "buffer size"))

func main() {
	log.SetFlags(log.Lmicroseconds | log.Flags())

	r := bufio.NewReaderSize(os.Stdin, bufSize)
	for {
		l, _, err := r.ReadLine()
		if err != nil || err == io.EOF {
			break
		}
		log.Println(string(l))
	}
}
