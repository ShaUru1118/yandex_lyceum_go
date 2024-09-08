package main

import (
	"math"
	"sort"
	"strings"
)

func CalculateSeriesSum(n int) float64 {
	s := 0.0
	for i := 1; i <= n; i++ {
		s += 1 / float64(i)
	}
	return s
}

func FindIntersection(k1, b1, k2, b2 float64) (float64, float64) {
	if k1 == k2 {
		return math.NaN(), math.NaN()
	}
	x := (b2 - b1) / (k1 - k2)
	y := k1*x + b1

	return x, y
}

func CalculateDigitalRoot(n int) int {
	sumOfDigits := func(x int) int {
		sum := 0
		for x > 0 {
			sum += x % 10
			x /= 10
		}
		return sum
	}

	for n >= 10 {
		n = sumOfDigits(n)
	}
	return n
}

func IsPalindrome(input string) bool {
	// Приведение строки к нижнему регистру и удаление пробелов
	cleanedInput := strings.ReplaceAll(strings.ToLower(input), " ", "")

	// Проверка палиндрома
	length := len(cleanedInput)
	for i := 0; i < length/2; i++ {
		if cleanedInput[i] != cleanedInput[length-1-i] {
			return false
		}
	}
	return true
}

func Permutations(input string) []string {
	var results []string
	permuteHelper([]rune(input), 0, len(input)-1, &results)
	sort.Strings(results) // Сортировка в алфавитном порядке
	return results
}

func permuteHelper(s []rune, l, r int, results *[]string) {
	if l == r {
		*results = append(*results, string(s))
	} else {
		for i := l; i <= r; i++ {
			s[l], s[i] = s[i], s[l] // обмен символов
			permuteHelper(s, l+1, r, results)
			s[l], s[i] = s[i], s[l] // возврат к исходному состоянию
		}
	}
}

func AreAnagrams(str1, str2 string) bool {
    cleanedStr1 := strings.ToLower(str1)
    cleanedStr2 := strings.ToLower(str2)

    sortedStr1 := sortString(cleanedStr1)
    sortedStr2 := sortString(cleanedStr2)

    return sortedStr1 == sortedStr2
}

func sortString(s string) string {
    runes := []rune(s)
    sort.Sort(sortRunes(runes))
    return string(runes)
}

type sortRunes []rune

func (s sortRunes) Len() int           { return len(s) }
func (s sortRunes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortRunes) Less(i, j int) bool { return s[i] < s[j] }
