package test

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

func IsPalindrome(value string) bool {
	for i := 0; i < len(value)/2; i++ {
		startIndex := i
		lastIndex := len(value) - i - 1

		if strings.ToLower(string(value[startIndex])) != strings.ToLower(string(value[lastIndex])) {
			return false
		}
	}

	return true
}

func LeapYear(x, y int) (result []int) {
	for i := x; i <= y; i++ {
		if i%4 == 0 {
			if i%100 == 0 {
				if i%400 == 0 {
					result = append(result, i)
				}
			} else {
				result = append(result, i)
			}
		}
	}

	return
}

func ReverseWords(value string) string {
	words := strings.Fields(value)
	result := ""
	for i, word := range words {
		for i := len(word) - 1; i >= 0; i-- {
			if unicode.IsUpper(rune(word[len(word)-1-i])) {
				result += strings.ToUpper(string(word[i]))
			} else {
				result += strings.ToLower(string(word[i]))
			}
		}

		if i != len(words)-1 {
			result += " "
		}
	}

	return result
}

func NearestFibonaci(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	first, second := 0, 1

	third := first + second

	for third <= sum {
		first = second
		second = third
		third = first + second
	}

	if math.Abs(float64(third)-float64(sum)) >= math.Abs(float64(second)-float64(sum)) {
		return sum - second
	}

	return third - sum
}

func FizzBuzz(n int) []string {
	result := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}

func Kosong() {
	return
}
