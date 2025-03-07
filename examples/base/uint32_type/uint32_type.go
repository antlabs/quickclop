package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic uint32 type
	Uint32 uint32 `clop:"-u,--uint32" usage:"A uint32 value" default:"42"`

	// Pointer uint32 type
	Uint32Ptr *uint32 `clop:"--uint32-ptr" usage:"A uint32 pointer value" default:"100"`

	// Help flag
	Help bool `clop:"-h,--help" usage:"Show help information"`

	F32Slice []float32 `clop:"--f32-slice" usage:"A slice of float32 values" default:"1.1,2.2,3.3"`

	F32SlicePtr *[]float32 `clop:"--f32-slice-ptr" usage:"A pointer to a slice of float32 values" default:"4.4,5.5,6.6"`
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

	// Display all values
	fmt.Println("Uint32 type:")
	fmt.Printf("  Uint32: %d\n", opts.Uint32)

	fmt.Println("\nUint32 pointer type:")
	if opts.Uint32Ptr != nil {
		fmt.Printf("  Uint32Ptr: %d\n", *opts.Uint32Ptr)
	} else {
		fmt.Println("  Uint32Ptr: nil")
	}
}
