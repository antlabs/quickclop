package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic uint type
	Uint     uint  `clop:"-u,--uint" usage:"A uint value" default:"42"`
	
	// Pointer uint type
	UintPtr  *uint `clop:"--uint-ptr" usage:"A uint pointer value" default:"100"`
	
	// Slice of uint type
	UintSlice []uint `clop:"--uint-slice" usage:"A slice of uint values" default:"1,2,3"`
	
	// Pointer to slice of uint type
	UintSlicePtr *[]uint `clop:"--uint-slice-ptr" usage:"A pointer to a slice of uint values" default:"4,5,6"`
	
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
	fmt.Println("Uint type:")
	fmt.Printf("  Uint: %d\n", opts.Uint)
	
	fmt.Println("\nUint pointer type:")
	if opts.UintPtr != nil {
		fmt.Printf("  UintPtr: %d\n", *opts.UintPtr)
	} else {
		fmt.Println("  UintPtr: nil")
	}

	fmt.Println("\nUint slice type:")
	fmt.Printf("  UintSlice: %v\n", opts.UintSlice)

	fmt.Println("\nUint slice pointer type:")
	if opts.UintSlicePtr != nil {
		fmt.Printf("  UintSlicePtr: %v\n", *opts.UintSlicePtr)
	} else {
		fmt.Println("  UintSlicePtr: nil")
	}
}
