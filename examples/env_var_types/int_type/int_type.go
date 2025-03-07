package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic int type with environment variable
	Int int `clop:"-i,--int" usage:"An int value" env:"INT_VALUE"`

	// Pointer int type with environment variable
	IntPtr *int `clop:"--int-ptr" usage:"An int pointer value" env:"INT_PTR_VALUE"`

	// Slice of int type with environment variable
	IntSlice []int `clop:"--int-slice" usage:"A slice of int values" env:"INT_SLICE_VALUE"`

	// Pointer to slice of int type with environment variable
	IntSlicePtr *[]int `clop:"--int-slice-ptr" usage:"A pointer to a slice of int values" env:"INT_SLICE_PTR_VALUE"`

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
	fmt.Println("Int type with environment variable:")
	fmt.Printf("  Int: %d\n", opts.Int)

	fmt.Println("\nInt pointer type with environment variable:")
	if opts.IntPtr != nil {
		fmt.Printf("  IntPtr: %d\n", *opts.IntPtr)
	} else {
		fmt.Println("  IntPtr: nil")
	}

	fmt.Println("\nInt slice type with environment variable:")
	fmt.Printf("  IntSlice: %v\n", opts.IntSlice)

	fmt.Println("\nInt slice pointer type with environment variable:")
	if opts.IntSlicePtr != nil {
		fmt.Printf("  IntSlicePtr: %v\n", *opts.IntSlicePtr)
	} else {
		fmt.Println("  IntSlicePtr: nil")
	}
}
