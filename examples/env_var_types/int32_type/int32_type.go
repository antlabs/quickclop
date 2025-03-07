package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic int32 type with environment variable
	Int32 int32 `clop:"-i,--int32" usage:"An int32 value" env:"INT32_VALUE"`

	// Pointer int32 type with environment variable
	Int32Ptr *int32 `clop:"--int32-ptr" usage:"An int32 pointer value" env:"INT32_PTR_VALUE"`

	// Slice of int32 type with environment variable
	Int32Slice []int32 `clop:"--int32-slice" usage:"A slice of int32 values" env:"INT32_SLICE_VALUE"`

	// Pointer to slice of int32 type with environment variable
	Int32SlicePtr *[]int32 `clop:"--int32-slice-ptr" usage:"A pointer to a slice of int32 values" env:"INT32_SLICE_PTR_VALUE"`

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
	fmt.Println("Int32 type with environment variable:")
	fmt.Printf("  Int32: %d\n", opts.Int32)

	fmt.Println("\nInt32 pointer type with environment variable:")
	if opts.Int32Ptr != nil {
		fmt.Printf("  Int32Ptr: %d\n", *opts.Int32Ptr)
	} else {
		fmt.Println("  Int32Ptr: nil")
	}

	fmt.Println("\nInt32 slice type with environment variable:")
	fmt.Printf("  Int32Slice: %v\n", opts.Int32Slice)

	fmt.Println("\nInt32 slice pointer type with environment variable:")
	if opts.Int32SlicePtr != nil {
		fmt.Printf("  Int32SlicePtr: %v\n", *opts.Int32SlicePtr)
	} else {
		fmt.Println("  Int32SlicePtr: nil")
	}
}
