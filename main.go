package main

import (
	"github.com/d3z41k/blockchain-boilerplate/cli"
	"os"
)

func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}
	cli.Run()
}
