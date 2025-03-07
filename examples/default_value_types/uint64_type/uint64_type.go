package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic uint64 type with default value
	Uint64 uint64 `clop:"-u,--uint64" usage:"A uint64 value" default:"9876543210"`

	// Pointer uint64 type with default value
	Uint64Ptr *uint64 `clop:"--uint64-ptr" usage:"A uint64 pointer value" default:"1234567890"`

	// Help flag
	Help bool `clop:"-h,--help" usage:"Show help information"`
}

func main() {
	var opts Options
	err := opts.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if opts.Help {
		opts.Usage()
		return
	}

	// Display all values
	fmt.Println("Uint64 type with default value:")
	fmt.Printf("  Uint64: %d\n", opts.Uint64)

	fmt.Println("\nUint64 pointer type with default value:")
	if opts.Uint64Ptr != nil {
		fmt.Printf("  Uint64Ptr: %d\n", *opts.Uint64Ptr)
	} else {
		fmt.Println("  Uint64Ptr: nil")
	}
}
