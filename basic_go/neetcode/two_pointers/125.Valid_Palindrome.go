package two_pointers

import (
	"fmt"
	"strings"
)

// FIXME: improve
func isPalindrome(s string) bool {
	new_s := ""
	for _, str := range s {
		if (str >= 48 && str <= 57) || (str >= 65 && str <= 90) || (str >= 97 && str <= 122) {
			new_s += string(str)
		}
	}

	if len(new_s) == 0 {
		return true
	}
	new_s = strings.ToLower(new_s)

	p1 := 0
	p2 := len(new_s) - 1

	for {
		if p1 == p2 || p2 < p1 || p1 > p2 {
			break
		}

		if new_s[p2] != new_s[p1] {
			return false
		}

		p1++
		p2--

	}
	return true
}

func Execute_isPalindrome() {
	s := "A man, a plan, a canal: Panama"
	s = "aa"
	ans := isPalindrome(s)
	fmt.Println(ans)
}
