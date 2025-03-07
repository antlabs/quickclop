package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic int16 type with environment variable
	Int16 int16 `clop:"-i,--int16" usage:"An int16 value" env:"INT16_VALUE"`

	// Pointer int16 type with environment variable
	Int16Ptr *int16 `clop:"--int16-ptr" usage:"An int16 pointer value" env:"INT16_PTR_VALUE"`

	// Slice of int16 type with environment variable
	Int16Slice []int16 `clop:"--int16-slice" usage:"A slice of int16 values" env:"INT16_SLICE_VALUE"`

	// Pointer to slice of int16 type with environment variable
	Int16SlicePtr *[]int16 `clop:"--int16-slice-ptr" usage:"A pointer to a slice of int16 values" env:"INT16_SLICE_PTR_VALUE"`

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
	fmt.Println("Int16 type with environment variable:")
	fmt.Printf("  Int16: %d\n", opts.Int16)

	fmt.Println("\nInt16 pointer type with environment variable:")
	if opts.Int16Ptr != nil {
		fmt.Printf("  Int16Ptr: %d\n", *opts.Int16Ptr)
	} else {
		fmt.Println("  Int16Ptr: nil")
	}

	fmt.Println("\nInt16 slice type with environment variable:")
	fmt.Printf("  Int16Slice: %v\n", opts.Int16Slice)

	fmt.Println("\nInt16 slice pointer type with environment variable:")
	if opts.Int16SlicePtr != nil {
		fmt.Printf("  Int16SlicePtr: %v\n", *opts.Int16SlicePtr)
	} else {
		fmt.Println("  Int16SlicePtr: nil")
	}
}
