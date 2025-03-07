package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic uint32 type with environment variable
	Uint32 uint32 `clop:"-u,--uint32" usage:"A uint32 value" env:"UINT32_VALUE"`

	// Pointer uint32 type with environment variable
	Uint32Ptr *uint32 `clop:"--uint32-ptr" usage:"A uint32 pointer value" env:"UINT32_PTR_VALUE"`

	// Slice of uint32 type with environment variable
	Uint32Slice []uint32 `clop:"--uint32-slice" usage:"A slice of uint32 values" env:"UINT32_SLICE_VALUE"`

	// Pointer to slice of uint32 type with environment variable
	Uint32SlicePtr *[]uint32 `clop:"--uint32-slice-ptr" usage:"A pointer to a slice of uint32 values" env:"UINT32_SLICE_PTR_VALUE"`

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
	fmt.Println("Uint32 type with environment variable:")
	fmt.Printf("  Uint32: %d\n", opts.Uint32)

	fmt.Println("\nUint32 pointer type with environment variable:")
	if opts.Uint32Ptr != nil {
		fmt.Printf("  Uint32Ptr: %d\n", *opts.Uint32Ptr)
	} else {
		fmt.Println("  Uint32Ptr: nil")
	}

	fmt.Println("\nUint32 slice type with environment variable:")
	fmt.Printf("  Uint32Slice: %v\n", opts.Uint32Slice)

	fmt.Println("\nUint32 slice pointer type with environment variable:")
	if opts.Uint32SlicePtr != nil {
		fmt.Printf("  Uint32SlicePtr: %v\n", *opts.Uint32SlicePtr)
	} else {
		fmt.Println("  Uint32SlicePtr: nil")
	}
}
