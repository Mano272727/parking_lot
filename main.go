package main

import (
	"flag"
)

func main() {
	flag.Parse()

	if len(flag.Args()) > 0 {
		ExecuteFile(flag.Args()[0])
		return
	}

	InteractiveSession()
}
