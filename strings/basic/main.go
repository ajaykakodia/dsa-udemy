package main

import (
	"fmt"
)

func main() {
	fmt.Println(changingCase("DaKu Dada"))
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
