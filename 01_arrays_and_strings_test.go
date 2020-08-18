package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestIsUnique(t *testing.T) {
	IsUnique := func(str string) bool {
		if len(str) > 128 {
			return false
		}

		counter := make(map[int32]bool)
		for _, c := range str {
			_, ok := counter[c]
			if ok {
				return false
			}

			counter[c] = true
		}

		return true
	}

	cases := []struct {
		str    string
		unique bool
	}{
		{"java", false},
		{"python", true},
		{"golang", false},
	}

	for _, c := range cases {
		if IsUnique(c.str) == c.unique {
			continue
		}

		t.Errorf("IsUnique(\"%v\"): expected=%v, actual=%v", c.str, c.unique, IsUnique(c.str))
	}
}

func TestPermutation(t *testing.T) {
	SortString := func(str string) string {
		s := strings.Split(str, "")
		sort.Strings(s)
		return strings.Join(s, "")
	}

	IsPermutation := func(str1, str2 string) bool {
		if len(str1) != len(str2) {
			return false
		}

		if SortString(str1) == SortString(str2) {
			return true
		}

		return false
	}

	cases := []struct {
		str1        string
		str2        string
		permutation bool
	}{
		{"java", "avaj", true},
		{"python", "jython", false},
	}

	for _, c := range cases {
		result := IsPermutation(c.str1, c.str2)
		if result == c.permutation {
			continue
		}

		t.Errorf("IsPermutation(\"%v\", \"%v\"): expected=%v, actual=%v", c.str1, c.str2, c.permutation, result)
	}
}

func TestURLify(t *testing.T) {
	URLify := func(str string, length int) string {
		return strings.Replace(str[:length], " ", "%20", -1)
	}

	cases := []struct {
		input  string
		length int
		output string
	}{
		{"Mr John Smith   ", 13, "Mr%20John%20Smith"},
	}

	for _, c := range cases {
		result := URLify(c.input, c.length)
		if result == c.output {
			continue
		}

		t.Errorf("URLify(\"%v\", \"%v\"): expected=%v, actual=%v", c.input, c.length, c.output, result)
	}
}

func TestPermutationOfPalindrome(t *testing.T) {
	Palindrome := func(str string) bool {
		rep := strings.Replace(str, " ", "", -1)
		low := strings.ToLower(rep)

		counter := make(map[int32]int)
		for _, s := range low {
			_, ok := counter[s]
			if ok {
				counter[s]++
				continue
			}

			counter[s] = 1
		}

		odd := false
		for _, v := range counter {
			if v%2 == 0 {
				continue
			}

			if odd {
				return false
			}
			odd = true
		}

		return true
	}

	cases := []struct {
		input      string
		palindrome bool
	}{
		{"Tact Coa", true},
	}

	for _, c := range cases {
		result := Palindrome(c.input)
		if result == c.palindrome {
			continue
		}

		t.Errorf("Palindrome(\"%v\"): expected=%v, actual=%v", c.input, c.palindrome, result)
	}
}

func TestOneAway(t *testing.T) {
	OneEditReplace := func(s1, s2 string) bool {
		found := false
		for i := range s1 {
			if s1[i] == s2[i] {
				continue
			}

			if found {
				return false
			}
			found = true
		}

		return true
	}

	OneEditInsert := func(s1, s2 string) bool {
		index1, index2 := 0, 0
		for {
			if index2 >= len(s2) || index1 >= len(s1) {
				break
			}

			if s1[index1] == s2[index2] {
				index1++
				index2++
				continue
			}

			if index1 != index2 {
				return false
			}
			index2++
		}

		return true
	}

	OneAway := func(str1, str2 string) bool {
		if len(str1) == len(str2) {
			return OneEditReplace(str1, str2)
		}

		if len(str1)+1 == len(str2) {
			return OneEditInsert(str1, str2)
		}

		if len(str1)-1 == len(str2) {
			return OneEditInsert(str2, str1)
		}

		return false
	}

	cases := []struct {
		str1    string
		str2    string
		oneaway bool
	}{
		{"pale", "ple", true},
		{"pales", "pale", true},
		{"pale", "bale", true},
		{"pale", "bae", false},
	}

	for _, c := range cases {
		result := OneAway(c.str1, c.str2)
		if result == c.oneaway {
			continue
		}

		t.Errorf("OneAway(\"%v\", \"%v\"): expected=%v, actual=%v", c.str1, c.str2, c.oneaway, result)
	}
}

func TestCompression(t *testing.T) {
	Count := func(str string) int {
		length := 0
		consecutive := 0
		for i := range str {
			consecutive++
			if i+1 >= len(str) || str[i] != str[i+1] {
				length = length + len(strconv.Itoa(consecutive))
				consecutive = 0
			}
		}

		return length
	}

	Compress := func(str string) string {
		length := Count(str)
		if length >= len(str) {
			return str
		}

		var sb strings.Builder
		sb.Grow(length)
		consecutive := 0

		for i := range str {
			consecutive++
			if i+1 >= len(str) || str[i] != str[i+1] {
				sb.WriteByte(str[i])
				sb.WriteString(strconv.Itoa(consecutive))
				consecutive = 0
			}
		}

		return sb.String()
	}

	cases := []struct {
		input  string
		output string
	}{
		{"aabcccccaaa", "a2b1c5a3"},
	}

	for _, c := range cases {
		result := Compress(c.input)
		if result == c.output {
			continue
		}

		t.Errorf("Compress(\"%v\"): expected=%v, actual=%v", c.input, c.output, result)
	}
}

func TestRotateMatrix(t *testing.T) {
	// TODO
}

func TestZeroMatrix(t *testing.T) {
	// TODO
}

func TestStringRotation(t *testing.T) {
	IsSubstring := func(str1, str2 string) bool {
		return strings.Contains(str1, str2)
	}

	IsRotation := func(str1, str2 string) bool {
		if len(str1) != len(str2) {
			return false
		}

		if len(str1) < 1 {
			return false
		}

		s1s1 := fmt.Sprintf("%s%s", str1, str2)
		return IsSubstring(s1s1, str2)
	}

	cases := []struct {
		str1     string
		str2     string
		rotation bool
	}{
		{"waterbottle", "erbottlewat", true},
	}

	for _, c := range cases {
		result := IsRotation(c.str1, c.str2)
		if result == c.rotation {
			continue
		}

		t.Errorf("IsRotation(\"%v\", \"%v\"): expected=%v, actual=%v", c.str1, c.str2, c.rotation, result)
	}
}
