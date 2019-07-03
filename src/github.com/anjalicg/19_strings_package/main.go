package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Compare 'abc' and 'ABC' %d\n", strings.Compare("abc", "ABC"))
	fmt.Printf("Compare 'abc' and 'abc' %d\n", strings.Compare("abc", "abc"))
	fmt.Printf("Compare 'ABC' and 'abc' %d\n", strings.Compare("ABC", "abc"))

	fmt.Printf(" 'abc' contains 'ab' %t\n", strings.Contains("abc", "ab"))
	fmt.Printf(" 'abc' contains 'ba' %t\n", strings.Contains("abc", "ba"))

	fmt.Printf(" 'abc' ContainsAny 'ab' %t\n", strings.ContainsAny("abc", "ab"))
	fmt.Printf(" 'abc' ContainsAny 'ba' %t\n", strings.ContainsAny("abc", "ba"))

	fmt.Printf("Count 'abc' contains 'abcaabcbcabcb' %d\n", strings.Count("abcaabcbcabcb", "abc"))

}
