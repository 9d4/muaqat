package main

import (
	"os"
)

func main() {
	var hanan Tulisan = "꧋ꦲꦤꦕꦫꦏ꧉"
	hanan.Write();
}

type Tulisan string

func (t Tulisan) Write() {
	os.Stdout.Write([]byte(t))
}
