package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Read
	val := strings.NewReader("line 1\nline 2\n")
	reader := bufio.NewReader(val)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	// Write
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("line 1\n")
	_, _ = writer.WriteString("line 2\n")
	writer.Flush()
}
