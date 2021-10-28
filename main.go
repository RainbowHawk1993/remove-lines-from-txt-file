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
	//scanner.Split(bufio.ScanLines)  	//doesn't see, like this does anything, but I'll keep it commented in case I'm forgetting what it's useful for

	var bs []byte
	buf := bytes.NewBuffer(bs)

	for scanner.Scan() {
		//finding words longer than 2 that should be kept
		if len(scanner.Text()) > 2 {
			_, err := buf.Write(scanner.Bytes())
			if err != nil {
				exit("Couldn't write line")
			}
			//words that are shorter than 2 are to be removed
			_, err = buf.WriteString("\n")
			if err != nil {
				exit("Couldn't replace line")
			}
		}
	}
	err = os.WriteFile(*txtFilename, buf.Bytes(), 0)
}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
