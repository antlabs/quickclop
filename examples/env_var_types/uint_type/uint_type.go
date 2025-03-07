package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic uint type with environment variable
	Uint uint `clop:"-u,--uint" usage:"A uint value" env:"UINT_VALUE"`

	// Pointer uint type with environment variable
	UintPtr *uint `clop:"--uint-ptr" usage:"A uint pointer value" env:"UINT_PTR_VALUE"`

	// Slice of uint type with environment variable
	UintSlice []uint `clop:"--uint-slice" usage:"A slice of uint values" env:"UINT_SLICE_VALUE"`

	// Pointer to slice of uint type with environment variable
	UintSlicePtr *[]uint `clop:"--uint-slice-ptr" usage:"A pointer to a slice of uint values" env:"UINT_SLICE_PTR_VALUE"`

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
	fmt.Println("Uint type with environment variable:")
	fmt.Printf("  Uint: %d\n", opts.Uint)

	fmt.Println("\nUint pointer type with environment variable:")
	if opts.UintPtr != nil {
		fmt.Printf("  UintPtr: %d\n", *opts.UintPtr)
	} else {
		fmt.Println("  UintPtr: nil")
	}

	fmt.Println("\nUint slice type with environment variable:")
	fmt.Printf("  UintSlice: %v\n", opts.UintSlice)

	fmt.Println("\nUint slice pointer type with environment variable:")
	if opts.UintSlicePtr != nil {
		fmt.Printf("  UintSlicePtr: %v\n", *opts.UintSlicePtr)
	} else {
		fmt.Println("  UintSlicePtr: nil")
	}
}
