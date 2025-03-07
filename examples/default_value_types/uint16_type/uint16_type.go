package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic uint16 type with default value
	Uint16 uint16 `clop:"-u,--uint16" usage:"A uint16 value" default:"1024"`

	// Pointer uint16 type with default value
	Uint16Ptr *uint16 `clop:"--uint16-ptr" usage:"A uint16 pointer value" default:"2048"`

	// Slice of uint16 type with default value
	Uint16Slice []uint16 `clop:"--uint16-slice" usage:"A slice of uint16 values" default:"100,200,300"`

	// Pointer to slice of uint16 type with default value
	Uint16SlicePtr *[]uint16 `clop:"--uint16-slice-ptr" usage:"A pointer to a slice of uint16 values" default:"400,500,600"`

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
	fmt.Println("Uint16 type with default value:")
	fmt.Printf("  Uint16: %d\n", opts.Uint16)

	fmt.Println("\nUint16 pointer type with default value:")
	if opts.Uint16Ptr != nil {
		fmt.Printf("  Uint16Ptr: %d\n", *opts.Uint16Ptr)
	} else {
		fmt.Println("  Uint16Ptr: nil")
	}

	fmt.Println("\nUint16 slice type with default value:")
	fmt.Printf("  Uint16Slice: %v\n", opts.Uint16Slice)

	fmt.Println("\nUint16 slice pointer type with default value:")
	if opts.Uint16SlicePtr != nil {
		fmt.Printf("  Uint16SlicePtr: %v\n", *opts.Uint16SlicePtr)
	} else {
		fmt.Println("  Uint16SlicePtr: nil")
	}
}
