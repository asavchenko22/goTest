package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	path := "/home/andrey/password"
	getFile(path)
}

func getFile(path string) {
	f, _ := os.Open(path)
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
