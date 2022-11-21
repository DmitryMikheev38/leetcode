package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("aab"))
}

//func lengthOfLongestSubstring(s string) int {
//	longestSubstringLength := 0
//	curSubstringLength := 0
//	substringMap := map[rune]interface{}{}
//	for _, char := range s {
//		if _, ok := substringMap[char]; !ok {
//			substringMap[char] = struct{}{}
//			curSubstringLength++
//		} else {
//			if longestSubstringLength <= curSubstringLength {
//				longestSubstringLength = curSubstringLength
//				curSubstringLength = 0
//				substringMap = map[rune]interface{}{}
//				substringMap[char] = struct{}{}
//				curSubstringLength++
//			}
//		}
//	}
//
//	if longestSubstringLength <= curSubstringLength {
//		longestSubstringLength = curSubstringLength
//	}
//
//	return longestSubstringLength
//}

//func lengthOfLongestSubstring(s string) int {
//	substringMap := map[uint8]int{}
//
//	longestSubstringLength := 0
//	curSubstringLength := 0
//
//	idx := 0
//	for ; idx < len(s); idx++ {
//		char := s[idx]
//		if i, ok := substringMap[char]; !ok {
//			substringMap[char] = idx
//			curSubstringLength++
//		} else {
//			idx = i
//			if longestSubstringLength <= curSubstringLength {
//				longestSubstringLength = curSubstringLength
//			}
//			curSubstringLength = 0
//			substringMap = map[uint8]int{}
//		}
//	}
//
//	if longestSubstringLength <= curSubstringLength {
//		longestSubstringLength = curSubstringLength
//	}
//
//	return longestSubstringLength
//}

func lengthOfLongestSubstring(s string) int {
	substringMap := map[rune]int{}
	l := 0

	res := 0
	for i, char := range s {
		if idx, ok := substringMap[char]; ok {
			l = int(math.Max(float64(l), float64(idx+1)))
		}
		substringMap[char] = i

		res = int(math.Max(float64(res), float64(i-l+1)))

	}

	return res
}
