package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic uint8 type with default value
	Uint8 uint8 `clop:"-u,--uint8" usage:"A uint8 value" default:"42"`

	// Pointer uint8 type with default value
	Uint8Ptr *uint8 `clop:"--uint8-ptr" usage:"A uint8 pointer value" default:"100"`

	// Slice of uint8 type with default value
	Uint8Slice []uint8 `clop:"--uint8-slice" usage:"A slice of uint8 values" default:"1,2,3"`

	// Pointer to slice of uint8 type with default value
	Uint8SlicePtr *[]uint8 `clop:"--uint8-slice-ptr" usage:"A pointer to a slice of uint8 values" default:"4,5,6"`

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
	fmt.Println("Uint8 type with default value:")
	fmt.Printf("  Uint8: %d\n", opts.Uint8)

	fmt.Println("\nUint8 pointer type with default value:")
	if opts.Uint8Ptr != nil {
		fmt.Printf("  Uint8Ptr: %d\n", *opts.Uint8Ptr)
	} else {
		fmt.Println("  Uint8Ptr: nil")
	}

	fmt.Println("\nUint8 slice type with default value:")
	fmt.Printf("  Uint8Slice: %v\n", opts.Uint8Slice)

	fmt.Println("\nUint8 slice pointer type with default value:")
	if opts.Uint8SlicePtr != nil {
		fmt.Printf("  Uint8SlicePtr: %v\n", *opts.Uint8SlicePtr)
	} else {
		fmt.Println("  Uint8SlicePtr: nil")
	}
}
