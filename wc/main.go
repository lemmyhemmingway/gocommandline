package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	count(os.Stdin)
}

func count(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	wc := 0

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		wc++
	}

	return wc
}
