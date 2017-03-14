// Sometimes we just want to completely replace the current Go process
// with another (perhaps non-Go) one. To do this we'll use Go's implementation
// of the classic exec function
package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// Go requires an absolute path to the binary we want to execute
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec requires arguments in slice form
	// First argument should be the program name
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec also needs a set of Environment variables to use
	// Here we just provide our current environment
	env := os.Environ()

	// Here's the actual syscall.Exec call. If this call is successful,
	// the execution of our process will end here and be replaced by
	// the /bin/ls -a -l -h process
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
