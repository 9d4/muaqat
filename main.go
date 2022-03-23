package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./endpoints.txt")
	if err != nil {
		panic(err)
	}

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}

		s := strings.Split(line, "/")

		for i, word := range s {
			if i == 0 {
				continue
			}

			f := strings.ToUpper(string(word[0]))
			final := f + word[1:]
			fmt.Print(final)
		}
	}
}
