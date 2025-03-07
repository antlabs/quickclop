package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic int32 type
	Int32     int32  `clop:"-i,--int32" usage:"An int32 value" default:"42"`
	
	// Pointer int32 type
	Int32Ptr  *int32 `clop:"--int32-ptr" usage:"An int32 pointer value" default:"100"`
	
	// Slice of int32 type
	Int32Slice []int32 `clop:"--int32-slice" usage:"A slice of int32 values" default:"1,2,3"`
	
	// Pointer to slice of int32 type
	Int32SlicePtr *[]int32 `clop:"--int32-slice-ptr" usage:"A pointer to a slice of int32 values" default:"4,5,6"`
	
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
	fmt.Println("Int32 type:")
	fmt.Printf("  Int32: %d\n", opts.Int32)
	
	fmt.Println("\nInt32 pointer type:")
	if opts.Int32Ptr != nil {
		fmt.Printf("  Int32Ptr: %d\n", *opts.Int32Ptr)
	} else {
		fmt.Println("  Int32Ptr: nil")
	}

	fmt.Println("\nInt32 slice type:")
	fmt.Printf("  Int32Slice: %v\n", opts.Int32Slice)

	fmt.Println("\nInt32 slice pointer type:")
	if opts.Int32SlicePtr != nil {
		fmt.Printf("  Int32SlicePtr: %v\n", *opts.Int32SlicePtr)
	} else {
		fmt.Println("  Int32SlicePtr: nil")
	}
}
