package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic uint32 type with default value
	Uint32 uint32 `clop:"-u,--uint32" usage:"A uint32 value" default:"123456"`

	// Pointer uint32 type with default value
	Uint32Ptr *uint32 `clop:"--uint32-ptr" usage:"A uint32 pointer value" default:"654321"`

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
	fmt.Println("Uint32 type with default value:")
	fmt.Printf("  Uint32: %d\n", opts.Uint32)

	fmt.Println("\nUint32 pointer type with default value:")
	if opts.Uint32Ptr != nil {
		fmt.Printf("  Uint32Ptr: %d\n", *opts.Uint32Ptr)
	} else {
		fmt.Println("  Uint32Ptr: nil")
	}
}
