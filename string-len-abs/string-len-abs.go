package main

import "fmt"

func main() {
	s := "hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	i := lengthOfLongestSubstring(s)
	fmt.Println(i)

}

func lengthOfLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		if len(s)-i < max {
			break
		}
		cache := make(map[byte]bool)
		for j := i; j < len(s); j++ {
			cache[s[j]] = true

			if max < len(cache) {
				max = len(cache)
			}
			if (j - i + 1) > len(cache) {
				break
			}

		}
	}
	return max
}

func lengthOfLongestSubstring1(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if len(s[i:j]) == countDiffRunes(s[i:j]) {
				if max < len(s[i:j]) {
					max = len(s[i:j])
				}
			}
		}
	}
	return max
}

func countDiffRunes(s string) int {
	a := make(map[rune]bool)
	for _, v := range s {
		a[v] = true
	}
	return len(a)
}
