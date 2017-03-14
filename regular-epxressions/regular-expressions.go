// Go offers built-in support for regex
package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// To tests whether a pattern matches a string
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// For regexp tasks you'll need to compile an optimized Regexp struct
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Many methods are available on these structs
	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))

	fmt.Println(r.FindStringIndex("peach punch"))

	// The Submatch variants include information about both the whole-pattern
	// matches and the submatches within those matches
	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// The All variants of these functions apply to all matches in the input, not just
	// the first
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// Providing a non-negative integer as the second argument to these functions will
	// limit the number of matches
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// We can also provide []byte arguments and drop String from the function name
	fmt.Println(r.Match([]byte("peach")))

	// When creating constants with regex you can use the MustCompile variation of Compile.
	// A plain Compile won't work for constants because it has 2 return values
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// The Func variant allows you to tranform matched text with a given function
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
