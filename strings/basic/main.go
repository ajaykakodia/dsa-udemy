package main

import (
	"fmt"
)

func main() {
	fmt.Println(changingCase("DaKu Dada"))
	fmt.Printf("Vowel Counts: %d\n", countVowels("Yaku is a very bad charactEr."))
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
