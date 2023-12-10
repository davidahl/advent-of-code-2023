package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numbers = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func substring(str string, start int, end int) string {
	return strings.TrimSpace(str[start:end])
}

func FindFirst(word string) string {
	for i := 0; i <= len(word); i++ {
		for k, v := range numbers {
			subs := substring(word, 0, i)
			if strings.Contains(subs, k) || strings.Contains(subs, v) {
				return v
			}
		}
	}
	return ""
}

func FindLast(word string) string {
	for i := 0; i <= len(word); i++ {
		for k, v := range numbers {

			subs := substring(word, len(word)-i, len(word))
			if strings.Contains(subs, k) || strings.Contains(subs, v) {
				return v
			}
		}
	}
	return ""
}

func main() {
	fmt.Println("Day one starts")

	day1, err := os.Open("./day1/day1.txt")

	if err != nil {
		log.Println(err.Error())
	}
	defer day1.Close()

	scanner := bufio.NewScanner(day1)
	// var words []string
	words := make([]string, 0)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	sum := 0
	for i := 0; i < len(words); i++ {
		num := FindFirst(words[i]) + FindLast(words[i])
		fmt.Println(num)
		lineValue, _ := strconv.Atoi(num)
		sum += lineValue

	}
	fmt.Print(sum)

}
