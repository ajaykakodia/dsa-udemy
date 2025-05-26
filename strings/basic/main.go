package main

import (
	"fmt"
)

func main() {
	fmt.Println(changingCase("DaKu Dada"))
	fmt.Printf("Vowel Counts: %d\n", countVowels("Yaku is a very bad charactEr."))
	fmt.Printf("Words Counts: %d\n", countWords("Yaku is a very   bad  charactEr."))
	fmt.Printf("Is string Palindrome: %v\n", checkingPalindrome("madam"))
	fmt.Printf("Is string Palindrome: %v\n", checkingPalindrome("madamam"))
	fmt.Printf("Is duplicate char in string: %v\n", findDuplicateInStringByBitwise("findings"))
	fmt.Printf("Is duplicate char in string: %v\n", findDuplicateInStringByBitwise("findegs"))
	fmt.Printf("Is strings are anagram: %v\n", isStringsAreAnagram("decimal", "medical"))
	fmt.Printf("Is strings are anagram: %v\n", isStringsAreAnagram("decimal", "medicallly"))
}

func changingCase(str string) string {
	changeCase := ""
	for i := 0; i < len(str); i++ {
		if str[i] >= 95 && str[i] <= 122 {
			changeCase = changeCase + string(str[i]-32)
		} else if str[i] >= 65 && str[i] <= 90 {
			changeCase = changeCase + string(str[i]+32)
		} else {
			changeCase = changeCase + string(str[i])
		}
	}

	return changeCase
}

func countVowels(str string) int {
	vowels := map[int]bool{'a': true, 'A': true, 'e': true, 'E': true, 'i': true, 'I': true, 'o': true, 'O': true, 'u': true, 'U': true}
	count := 0
	for i := 0; i < len(str); i++ {
		if vowels[int(str[i])] {
			count++
		}
	}
	return count
}

func countWords(str string) int {
	if str == "" {
		return 0
	}
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' && str[i-1] != ' ' {
			count++
		}
	}
	return count + 1
}

func checkingPalindrome(str string) bool {
	i, j := 0, len(str)-1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func findDuplicateInStringByBitwise(str string) bool {
	h, x := int64(0), int64(0)
	for i := 0; i < len(str); i++ {
		x = 1
		bitToMove := str[i] - 97
		x = x << bitToMove
		if h&x > 0 {
			return true
		}
		h = x | h
	}
	return false
}

func isStringsAreAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	h, x := 0, 0
	for i := 0; i < len(str1); i++ {
		x = 1
		x = x << (str1[i] - 97)
		h = h | x
	}

	for i := 0; i < len(str2); i++ {
		x = 1
		x = x << (str2[i] - 97)
		h = h & x
	}

	return h == 0
}
