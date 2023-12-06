package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"

	"golang.org/x/exp/constraints"
)

func sum[T constraints.Signed](array []T) int64 {
	var result int64
	result = 0
	for _, v := range array {
		result += int64(v)
	}
	return result
}

func LineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func main() {
	file := "./day_1/input.txt"
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Errorf("ERROR: %s", err)
	}

	// scanner to read file by lines
	scanner := bufio.NewScanner(bytes.NewReader(data))

	// init counter of value
	values := []int8{}
	var current_number_str string

	for scanner.Scan() {
		current_number_str = ""
		txt := scanner.Text()
		// loop each char forward
		for _, char := range txt {
			if unicode.IsDigit(char) {
				current_number_str = current_number_str + string(char)
				break
			}
		}
		// loop each char backward
		for i := len(txt) - 1; i > -1; i-- {
			if unicode.IsDigit(rune(txt[i])) {
				current_number_str = current_number_str + string(txt[i])
				break
			}
		}
		int_value, _ := strconv.Atoi(current_number_str)
		values = append(values, int8(int_value))

	}

	fmt.Println(len(values)) // verify that the number of values match the number of lines
	total_sum := sum(values)
	fmt.Println(total_sum)
}
