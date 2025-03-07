package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic int type with default value
	Int int `clop:"-i,--int" usage:"An int value" default:"42"`

	// Pointer int type with default value
	IntPtr *int `clop:"--int-ptr" usage:"An int pointer value" default:"100"`

	// Slice of int type with default value
	IntSlice []int `clop:"--int-slice" usage:"A slice of int values" default:"1,2,3"`

	// Pointer to slice of int type with default value
	IntSlicePtr *[]int `clop:"--int-slice-ptr" usage:"A pointer to a slice of int values" default:"4,5,6"`

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
	fmt.Println("Int type with default value:")
	fmt.Printf("  Int: %d\n", opts.Int)

	fmt.Println("\nInt pointer type with default value:")
	if opts.IntPtr != nil {
		fmt.Printf("  IntPtr: %d\n", *opts.IntPtr)
	} else {
		fmt.Println("  IntPtr: nil")
	}

	fmt.Println("\nInt slice type with default value:")
	fmt.Printf("  IntSlice: %v\n", opts.IntSlice)

	fmt.Println("\nInt slice pointer type with default value:")
	if opts.IntSlicePtr != nil {
		fmt.Printf("  IntSlicePtr: %v\n", *opts.IntSlicePtr)
	} else {
		fmt.Println("  IntSlicePtr: nil")
	}
}
