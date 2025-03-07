package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic uint64 type with environment variable
	Uint64 uint64 `clop:"-u,--uint64" usage:"A uint64 value" env:"UINT64_VALUE"`

	// Pointer uint64 type with environment variable
	Uint64Ptr *uint64 `clop:"--uint64-ptr" usage:"A uint64 pointer value" env:"UINT64_PTR_VALUE"`

	// Slice of uint64 type with environment variable
	Uint64Slice []uint64 `clop:"--uint64-slice" usage:"A slice of uint64 values" env:"UINT64_SLICE_VALUE"`

	// Pointer to slice of uint64 type with environment variable
	Uint64SlicePtr *[]uint64 `clop:"--uint64-slice-ptr" usage:"A pointer to a slice of uint64 values" env:"UINT64_SLICE_PTR_VALUE"`

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
	fmt.Println("Uint64 type with environment variable:")
	fmt.Printf("  Uint64: %d\n", opts.Uint64)

	fmt.Println("\nUint64 pointer type with environment variable:")
	if opts.Uint64Ptr != nil {
		fmt.Printf("  Uint64Ptr: %d\n", *opts.Uint64Ptr)
	} else {
		fmt.Println("  Uint64Ptr: nil")
	}

	fmt.Println("\nUint64 slice type with environment variable:")
	fmt.Printf("  Uint64Slice: %v\n", opts.Uint64Slice)

	fmt.Println("\nUint64 slice pointer type with environment variable:")
	if opts.Uint64SlicePtr != nil {
		fmt.Printf("  Uint64SlicePtr: %v\n", *opts.Uint64SlicePtr)
	} else {
		fmt.Println("  Uint64SlicePtr: nil")
	}
}
