package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the camelcase function below.
// func camelcase(s string) int32 {
// 	counter := int32(1)
// 	for _, c := range s {
// 		if strings.ToUpper(string(c)) == string(c) {
// 			counter++
// 		}
// 	}
// 	return counter
// }

func camelcase(s string) int32 {
	counter := int32(1)
	min, max := 'A', 'Z'
	for _, c := range s {
		if c >= min && c <= max {
			counter++
		}
	}
	return counter
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := camelcase(s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
