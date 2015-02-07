package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:] // exclude program name
	s := strings.Join(args, " ")
	fmt.Println(s)
}
