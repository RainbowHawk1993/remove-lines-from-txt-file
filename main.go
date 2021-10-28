package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
)

func main() {
	txtFilename := flag.String("txt", "words.txt",
		"a txt file with 1 word per line")
	flag.Parse()

	file, err := os.Open(*txtFilename)
	if err != nil {
		exit("Failed to open txt file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	//scanner.Split(bufio.ScanLines)

	var bs []byte
	buf := bytes.NewBuffer(bs)

	var text string
	for scanner.Scan() {
		text = scanner.Text()
		length := len(text)
		if length > 3 {
			_, err := buf.Write(scanner.Bytes())
			if err != nil {
				exit("Couldn't write line")
			}
			_, err = buf.WriteString("aaaa")
			if err != nil {
				exit("Couldn't replace line")
			}
		}
	}
}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
