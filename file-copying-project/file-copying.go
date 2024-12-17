package main

import (
	"fmt"
	"log"
	"os"
)

func CopyFile(src string, dst string) {
	srcfile, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer srcfile.Close()

	srcfilestat, err := srcfile.Stat()
	if err != nil {
		log.Fatal(err)
	}
	buffer := make([]byte, srcfilestat.Size())
	dstfile, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer dstfile.Close()

	count, err := srcfile.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	countdst, err := dstfile.Write(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes: %q\n", count, buffer)
	fmt.Printf("writen %d bytes: %q\n", countdst, buffer)
}

func main() {

	Arguments := os.Args
	if len(Arguments) != 3 {
		log.Fatalf("2 argumetns are needed but provided %d", len(Arguments)-1)
	}
	CopyFile(Arguments[1], Arguments[2])
}
