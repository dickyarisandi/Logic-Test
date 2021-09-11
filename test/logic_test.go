package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	result := IsPalindrome("Radar")

	assert.Equal(t, true, result, "Palindrome")
}

func TestLeapYear(t *testing.T) {
	result := LeapYear(1900, 2020)

	assert.ElementsMatch(t, []int{1904, 1908, 1912, 1916, 1920, 1924, 1928, 1932, 1936, 1940, 1944, 1948, 1952, 1956, 1960, 1964, 1968, 1972, 1976, 1980, 1984, 1988, 1992, 1996, 2000, 2004, 2008, 2012, 2016, 2020}, result)
}

func TestReverseWords(t *testing.T) {
	result := ReverseWords("I am A Great human")

	assert.Equal(t, "I ma A Taerg namuh", result, "Equal")
}

func TestNearestFibonaci(t *testing.T) {
	result := NearestFibonaci([]int{15, 1, 3})

	assert.Equal(t, 2, result, "Equal")
}

func TestFizzBuzz(t *testing.T) {
	result := FizzBuzz(15)

	assert.ElementsMatch(t, []string{"1", "2", "Fizz", "4", "Buzz"}, result)
}
