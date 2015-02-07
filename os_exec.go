package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Roughly equivalent to backticks or system in Ruby
	output, err := exec.Command("ls", "-l", "/").CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}
