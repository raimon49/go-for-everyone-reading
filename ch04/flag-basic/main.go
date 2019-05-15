package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

const defaultPort = 3000
var port int

func main() {
	envPort, _ := strconv.Atoi(os.Getenv("PORT"))

	if envPort > 0 {
		flag.IntVar(&port, "port", envPort, "port to use")
		flag.IntVar(&port, "p", envPort, "port to use(short)")

	} else {
		flag.IntVar(&port, "port", defaultPort, "port to use")
		flag.IntVar(&port, "p", defaultPort, "port to use(short)")
	}

	flag.Parse()

	fmt.Printf("port: %d", port)
	fmt.Println()
}
