//not optimized function. Search for the lenght of the longest substring without repeating characters

import "strings"

func lengthOfLongestSubstring(s string) int {

	d := strings.Split(s, "")

	n := 0
	p := 0
	for {
		if len(d) > 0 {
			n = count(d)
			if n > p {
				p = n
			}
			d = sliceChange(n, d)

		} else {
			break
		}

	}
	return p
}

func count(y []string) int {
	q := map[string]int{}
	for i := 0; i < len(y); i++ {
		q[y[i]]++
		if q[y[i]] > 1 {
			return i

		}

	}
	return len(y)
}

func sliceChange(s int, y []string) []string {
	if s >= len(y) {
		y = y[0:0]
		return y
	}

	for i, x := range y {
		if x == y[s] {
			y = y[i+1:]
			return y
		}
	}

	return y
}
