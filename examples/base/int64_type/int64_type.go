package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic int64 type
	Int64     int64  `clop:"-i,--int64" usage:"An int64 value" default:"42"`
	
	// Pointer int64 type
	Int64Ptr  *int64 `clop:"--int64-ptr" usage:"An int64 pointer value" default:"100"`
	
	// Slice of int64 type
	Int64Slice []int64 `clop:"--int64-slice" usage:"A slice of int64 values" default:"1,2,3"`
	
	// Pointer to slice of int64 type
	Int64SlicePtr *[]int64 `clop:"--int64-slice-ptr" usage:"A pointer to a slice of int64 values" default:"4,5,6"`
	
	// Help flag
	Help      bool   `clop:"-h,--help" usage:"Show help information"`
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
	fmt.Println("Int64 type:")
	fmt.Printf("  Int64: %d\n", opts.Int64)
	
	fmt.Println("\nInt64 pointer type:")
	if opts.Int64Ptr != nil {
		fmt.Printf("  Int64Ptr: %d\n", *opts.Int64Ptr)
	} else {
		fmt.Println("  Int64Ptr: nil")
	}

	fmt.Println("\nInt64 slice type:")
	fmt.Printf("  Int64Slice: %v\n", opts.Int64Slice)

	fmt.Println("\nInt64 slice pointer type:")
	if opts.Int64SlicePtr != nil {
		fmt.Printf("  Int64SlicePtr: %v\n", *opts.Int64SlicePtr)
	} else {
		fmt.Println("  Int64SlicePtr: nil")
	}
}
