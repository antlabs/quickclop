package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic int type
	Int     int  `clop:"-i,--int" usage:"An integer value" default:"42"`
	
	// Pointer int type
	IntPtr  *int `clop:"--int-ptr" usage:"An integer pointer value" default:"100"`
	
	// Slice of int type
	IntSlice []int `clop:"--int-slice" usage:"A slice of int values" default:"1,2,3"`
	
	// Pointer to slice of int type
	IntSlicePtr *[]int `clop:"--int-slice-ptr" usage:"A pointer to a slice of int values" default:"4,5,6"`
	
	// Help flag
	Help    bool `clop:"-h,--help" usage:"Show help information"`
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
	fmt.Println("Int type:")
	fmt.Printf("  Int: %d\n", opts.Int)
	
	fmt.Println("\nInt pointer type:")
	if opts.IntPtr != nil {
		fmt.Printf("  IntPtr: %d\n", *opts.IntPtr)
	} else {
		fmt.Println("  IntPtr: nil")
	}

	fmt.Println("\nInt slice type:")
	fmt.Printf("  IntSlice: %v\n", opts.IntSlice)

	fmt.Println("\nInt slice pointer type:")
	if opts.IntSlicePtr != nil {
		fmt.Printf("  IntSlicePtr: %v\n", *opts.IntSlicePtr)
	} else {
		fmt.Println("  IntSlicePtr: nil")
	}
}
