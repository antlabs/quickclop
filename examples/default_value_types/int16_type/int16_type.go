package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic int16 type with default value
	Int16 int16 `clop:"-i,--int16" usage:"An int16 value" default:"1024"`

	// Pointer int16 type with default value
	Int16Ptr *int16 `clop:"--int16-ptr" usage:"An int16 pointer value" default:"2048"`

	// Slice of int16 type with default value
	Int16Slice []int16 `clop:"--int16-slice" usage:"A slice of int16 values" default:"100,200,300"`

	// Pointer to slice of int16 type with default value
	Int16SlicePtr *[]int16 `clop:"--int16-slice-ptr" usage:"A pointer to a slice of int16 values" default:"400,500,600"`

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
	fmt.Println("Int16 type with default value:")
	fmt.Printf("  Int16: %d\n", opts.Int16)

	fmt.Println("\nInt16 pointer type with default value:")
	if opts.Int16Ptr != nil {
		fmt.Printf("  Int16Ptr: %d\n", *opts.Int16Ptr)
	} else {
		fmt.Println("  Int16Ptr: nil")
	}

	fmt.Println("\nInt16 slice type with default value:")
	fmt.Printf("  Int16Slice: %v\n", opts.Int16Slice)

	fmt.Println("\nInt16 slice pointer type with default value:")
	if opts.Int16SlicePtr != nil {
		fmt.Printf("  Int16SlicePtr: %v\n", *opts.Int16SlicePtr)
	} else {
		fmt.Println("  Int16SlicePtr: nil")
	}
}
