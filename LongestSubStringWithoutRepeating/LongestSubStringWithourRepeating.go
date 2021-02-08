package leetcode

func lengthOfLongestSubstring(s string) (length int) {
	length = 0
	if len(s) == 0 {
		return
	}
	var frequence [256]int
	left, right := 0, 0
	for left < len(s) {
		if right < len(s) && frequence[s[right]] == 0 {
			frequence[s[right]]++
			right++
		} else {
			frequence[s[left]]--
			left++
		}
		length = max(length, right-left)
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
