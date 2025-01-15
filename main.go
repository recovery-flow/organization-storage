package main

import "os"

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
