/*
Given a string s, rearrange the characters of s so that any two adjacent characters are not the same.
Return any possible rearrangement of s or return "" if not possible.

Example 1:
Input: s = "aab"
Output: "aba"
Example 2:
Input: s = "aaab"
Output: ""
*/

package main

import (
	"fmt"
	"sort"
)

type CharOccurance struct {
	char  rune
	count int
}

func rearrangeCharacters(s string) string {
	charCountMap := make(map[rune]int)
	var counts []CharOccurance

	// find out the count of each character
	for _, char := range s {
		charCountMap[char]++
	}

	for char, count := range charCountMap {
		counts = append(counts, CharOccurance{char, count})
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i].count > counts[j].count
	})

	var result []rune
	for len(result) < len(s) {
		// if first character in counts not same as the current character in result, appending the character to result
		if len(result) == 0 || result[len(result)-1] != counts[0].char {
			result = append(result, counts[0].char)
			// decreasing the character count as we already append it to result
			counts[0].count--
		} else {
			// If counts have only one character or the next character have 0 count, return empty string
			if len(counts) < 2 || counts[1].count == 0 {
				return ""
			}
			// appending the next character to result and decreasing it's count
			result = append(result, counts[1].char)
			counts[1].count--
		}

		// sort the counts slice with latest character count values
		sort.Slice(counts, func(i, j int) bool {
			return counts[i].count > counts[j].count
		})
	}

	return string(result)
}

func main() {
	fmt.Println(rearrangeCharacters("acbdba"))        // Output: "babacd"
	fmt.Println(rearrangeCharacters("acbdbacccefgk")) // Output: "cacbcbcadefgk"
	fmt.Println(rearrangeCharacters("aaaba"))         // Output: ""
}
