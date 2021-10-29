package removeL

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func removeLines(txtFilename string) {

	file, err := os.Open(txtFilename)
	if err != nil {
		exit("Failed to open txt file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

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
	os.WriteFile(txtFilename, buf.Bytes(), 0666)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
