package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	Headers []string `clop:"-H,--header" usage:"HTTP headers in format key:value"`
	Help    bool     `clop:"-h,--help" usage:"Show help"`
}

func main() {
	var opts Options
	if err := opts.Parse(os.Args[1:]); err != nil {
		log.Fatalf("Error parsing arguments: %v", err)
	}

	if opts.Help {
		opts.Usage()
		return
	}

	fmt.Println("Headers:")
	for i, header := range opts.Headers {
		fmt.Printf("  %d: %s\n", i+1, header)
	}
}
