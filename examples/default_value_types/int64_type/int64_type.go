package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic int64 type with default value
	Int64 int64 `clop:"-i,--int64" usage:"An int64 value" default:"9876543210"`

	// Pointer int64 type with default value
	Int64Ptr *int64 `clop:"--int64-ptr" usage:"An int64 pointer value" default:"1234567890"`

	// Slice of int64 type with default value
	Int64Slice []int64 `clop:"--int64-slice" usage:"A slice of int64 values" default:"1000000,2000000,3000000"`

	// Pointer to slice of int64 type with default value
	Int64SlicePtr *[]int64 `clop:"--int64-slice-ptr" usage:"A pointer to a slice of int64 values" default:"4000000,5000000,6000000"`

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
	fmt.Println("Int64 type with default value:")
	fmt.Printf("  Int64: %d\n", opts.Int64)

	fmt.Println("\nInt64 pointer type with default value:")
	if opts.Int64Ptr != nil {
		fmt.Printf("  Int64Ptr: %d\n", *opts.Int64Ptr)
	} else {
		fmt.Println("  Int64Ptr: nil")
	}
}
