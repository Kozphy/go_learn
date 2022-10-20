package array_hashing

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var mapping_t map[rune]int = make(map[rune]int)
	for _, str := range t {
		mapping_t[str] += 1
	}

	for _, str := range s {
		if mapping_t[str] == 0 {
			return false
		}
		mapping_t[str] -= 1
	}
	return true
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sn := strings.Split(s, "")
	sort.Strings(sn)
	tn := strings.Split(t, "")
	sort.Strings(tn)

	for pos, str := range tn {
		if sn[pos] != str {
			return false
		}
	}
	return true
}

func Execute_isAnagram() {
	s := "anagram"
	// s = "a"
	t := "nagaram"
	// t = "ab"
	ans := isAnagram2(s, t)
	fmt.Println(ans)
}
