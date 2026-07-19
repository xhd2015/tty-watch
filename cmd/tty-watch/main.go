package main

import (
	"fmt"
	"os"

	"github.com/xhd2015/tty-watch/cli"
)

func main() {
	if err := cli.Main(os.Args[1:], os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
