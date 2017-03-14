// Command-line flags are a common way to specify options for command-line programs
// For example, in wc -l the -l is a command-line flag
package main

// Go provides a flag package supporting basic command-line flag parsing
import (
	"flag"
	"fmt"
)

func main() {

	// Basic flag declarations are available for string, integer and boolean options
	// a string flag word with a default value "foo" and a short description
	// flag.String function returns a string pointer (not a string value)
	wordPtr := flag.String("word", "foo", "a string")

	// flag.Int
	numbPtr := flag.Int("numb", 42, "an int")

	// flag.Bool
	boolPtr := flag.Bool("fork", false, "a bool")

	// It's also possible to declare an option that uses an existing var declared
	// elsewhere in the program. Note that we need to pass in a pointer to the flag
	// declaration function
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Once all flags are declared, call flag.Parse() to execute the command-line parsing
	flag.Parse()

	// We'll just dump out the parsed options and any trailing positional arguments. Note
	// that we need to dereference the pointers with e.g *wordPtr to get the actual
	// option values
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

	// use -h or --help flags to get automatically generated help text for the
	// command-line program
	// Trailing positional arguments can be provided after any flags
	// Note that the flag package requires all falgs to appear before positional
	// arguments (otherwise the flags will be interpreted as positional arguments)
}
