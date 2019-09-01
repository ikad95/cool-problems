package main

import (
	"fmt"
)

var collapsed map[string]string

func smallest(a string,b string,c string,d string) string {
	ab := a
	if len(a) > len(b) {
		ab = b
	}
	cd := c
	if len(c) > len(d) {
		cd = d
	}
	abcd := ab
	if len(ab) > len(cd) {
		abcd = cd
	}
	return abcd
}

func removeStupidity(s string) string {
	res := ""
	lastChar := rune(0)
	count := 0
	for _, char := range s {
		if lastChar != char {
			lastChar = char
			count = 0
		}
		count++
		if count < 3 {
			res = res + string(char)
		}
	}
	return res
}

func collapse(s string) string {
	s = removeStupidity(s)			// because aaaabb is essentially same as aabb. Very Important step!
	if len(s) < 2 {
		return s
	} else if len(s) == 2 {
		if s[0] == s[1] {
			return ""
		}
		return s
	}

	_, existsInCollapsedMap := collapsed[s]
	if existsInCollapsedMap {
		return collapsed[s]
	}

	collapsed[s] = s

	leng := len(s)

	// aX => aY
	newSVal1 := collapse(string(s[0]) + collapse(s[1:]))

	// Xb => Yb
	newSVal2 := collapse(collapse(s[:leng - 1]) + string(s[leng - 1]))

	// aXc => aYc
	newSVal3 := collapse(string(s[0]) + collapse(s[1:leng - 1]) + string(s[leng - 1]))

	collapsed[s] = smallest(newSVal1, newSVal2, newSVal3, collapsed[s])
	return collapsed[s]
}

func main()  {
	for true {
		collapsed = make(map[string]string)
		var s string
		_,err := fmt.Scanln(&s)
		if err != nil{
			fmt.Println("Lel")
		}
		x := collapse(s)
		fmt.Printf("Ans : %d\n\n", len(x))
	}
}
