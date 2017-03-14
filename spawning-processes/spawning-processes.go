// Sometimes our Go programs need to spawn other, non-Go processes
// For example, the syntax highlighting on this site is implemented
// by spawning a pygmentize (link: http://pygments.org/) process
// from a Go program
package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	// exec.Command helper creates an object to represent this external process
	dateCmd := exec.Command("date")

	// .Output is another helper that handlers the common case of running a command,
	// waiting for it to finish, and collecting its output. If there were no erros,
	// dataOut will hold bytes with the data info
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Here we pipe data to the external process on its stdin and collect the results form its
	// stdout
	grepCmd := exec.Command("grep", "hello")

	// We explicitly grab input/output pipes, start the process, write some input to it,
	// read the resulting output and finally wait for the process to exit
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	// We only collect the StdoutPipe results, but you could collect the StderrPipe in exactly the same way
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Note that when spawning commands we need to provide an explicitly delineated command
	// and argument array vs being able to just pass in one command-line string. If you want to
	// spawn a full command with a string, you can use bash's -c option
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
