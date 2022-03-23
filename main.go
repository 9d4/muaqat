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

		line = strings.ReplaceAll(line, string('\n'), "")
		s := strings.Split(line, "/")

		final := ""
		for i, word := range s {
			if i == 0 {
				continue
			}

			f := strings.ToUpper(string(word[0]))
			final += f + word[1:]
		}

		final += " = " + "\"" + line + "\""

		fmt.Println(final)
		// break
	}
}
