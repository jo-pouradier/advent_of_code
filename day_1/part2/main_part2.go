package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

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

var digit_map = map[string]int{
	"one":   1,
	"1":     1,
	"two":   2,
	"2":     2,
	"three": 3,
	"3":     3,
	"four":  4,
	"4":     4,
	"5":     5,
	"five":  5,
	"6":     6,
	"six":   6,
	"7":     7,
	"seven": 7,
	"8":     8,
	"eight": 8,
	"9":     9,
	"nine":  9,
}

var all_digit_str = reflect.ValueOf(digit_map).MapKeys()

func main() {
	file := "./day_1/part2/input2.txt"
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Errorf("ERROR: %s", err)
		return
	}

	// scanner to read file by lines
	scanner := bufio.NewScanner(bytes.NewReader(data))
	// scanner := bufio.NewScanner(strings.NewReader(data))
	total := 0
	for scanner.Scan() {
		index_each_digit := map[int]string{}
		txt := scanner.Text()
		fmt.Println("txt=", txt)
		// save surrent line in variable because txt will be modified later
		init_txt := txt
		for _, digit_str := range all_digit_str {
			txt := init_txt
			counter_appearance := strings.Count(txt, digit_str.String())
			// loop to get every occurence of a digit
			for i := 0; i < counter_appearance; i++ {
				// index_each_digit = append(index_each_digit, strings.Index(txt, digit_str))
				substring_idx := strings.Index(txt, digit_str.String())
				if substring_idx > -1 {
					index_each_digit[substring_idx] = digit_str.String()
					// replace the current digit by a string of same size otherwise we dont have same index after
					txt = strings.Replace(
						txt,
						digit_str.Interface().(string),
						strings.Repeat("-", len(digit_str.String())),
						1,
					)
				}
			}
		}
		// read all index in order to get first and last
		keys := []int{}
		for key := range index_each_digit {
			keys = append(keys, key)
		}
		sort.Ints(keys)
		// get the first number aka tens
		tens := index_each_digit[keys[0]]
		// get last number aka unit
		unit := index_each_digit[keys[len(keys)-1]]
		// create the correct number based on tens and unit
		number := digit_map[tens]*10 + digit_map[unit]
		fmt.Println("number=", number)
		total += number
	}

	fmt.Println(total)
}
