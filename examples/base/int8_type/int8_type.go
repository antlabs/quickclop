package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic int8 type
	Int8     int8  `clop:"-i,--int8" usage:"An int8 value" default:"42"`
	
	// Pointer int8 type
	Int8Ptr  *int8 `clop:"--int8-ptr" usage:"An int8 pointer value" default:"100"`
	
	// Slice of int8 type
	Int8Slice []int8 `clop:"--int8-slice" usage:"A slice of int8 values" default:"1,2,3"`
	
	// Pointer to slice of int8 type
	Int8SlicePtr *[]int8 `clop:"--int8-slice-ptr" usage:"A pointer to a slice of int8 values" default:"4,5,6"`
	
	// Help flag
	Help     bool  `clop:"-h,--help" usage:"Show help information"`
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
	fmt.Println("Int8 type:")
	fmt.Printf("  Int8: %d\n", opts.Int8)
	
	fmt.Println("\nInt8 pointer type:")
	if opts.Int8Ptr != nil {
		fmt.Printf("  Int8Ptr: %d\n", *opts.Int8Ptr)
	} else {
		fmt.Println("  Int8Ptr: nil")
	}

	fmt.Println("\nInt8 slice type:")
	fmt.Printf("  Int8Slice: %v\n", opts.Int8Slice)

	fmt.Println("\nInt8 slice pointer type:")
	if opts.Int8SlicePtr != nil {
		fmt.Printf("  Int8SlicePtr: %v\n", *opts.Int8SlicePtr)
	} else {
		fmt.Println("  Int8SlicePtr: nil")
	}
}
