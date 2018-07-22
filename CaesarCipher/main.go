package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the caesarCipher function below.
func caesarCipher(s string, k int32) string {
	originLower := "abcdefghijklmnopqrstuvwxyz"
	originUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ret := ""

	for _, r := range s {
		switch {
		case strings.IndexRune(originLower, r) >= 0:
			ret = ret + string(rotate(r, k, []rune(originLower)))
		case strings.IndexRune(originUpper, r) >= 0:
			ret = ret + string(rotate(r, k, []rune(originUpper)))
		default:
			ret = ret + string(r)
		}
	}
	return ret
}

func rotate(s rune, delta int32, key []rune) rune {
	idx := strings.IndexRune(string(key), s)
	if idx < 0 {
		fmt.Errorf("Got error in rotate()")
	}
	idx = (idx + int(delta)) % len(key)
	return key[idx]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	_, err = strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

	fmt.Fprintf(writer, "%s\n", result)

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
