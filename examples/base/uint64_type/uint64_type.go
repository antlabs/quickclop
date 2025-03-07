package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic uint64 type
	Uint64 uint64 `clop:"-u,--uint64" usage:"A uint64 value" default:"42"`

	// Pointer uint64 type
	Uint64Ptr *uint64 `clop:"--uint64-ptr" usage:"A uint64 pointer value" default:"100"`

	// Slice of uint64 type
	Uint64Slice []uint64 `clop:"--uint64-slice" usage:"A slice of uint64 values" default:"1,2,3"`

	// Pointer to slice of uint64 type
	Uint64SlicePtr *[]uint64 `clop:"--uint64-slice-ptr" usage:"A pointer to a slice of uint64 values" default:"4,5,6"`

	// Help flag
	Help bool `clop:"-h,--help" usage:"Show help information"`
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
	fmt.Println("Uint64 type:")
	fmt.Printf("  Uint64: %d\n", opts.Uint64)

	fmt.Println("\nUint64 pointer type:")
	if opts.Uint64Ptr != nil {
		fmt.Printf("  Uint64Ptr: %d\n", *opts.Uint64Ptr)
	} else {
		fmt.Println("  Uint64Ptr: nil")
	}

	fmt.Println("\nUint64 slice type:")
	if len(opts.Uint64Slice) > 0 {
		fmt.Print("  Uint64Slice: ")
		for i, val := range opts.Uint64Slice {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	} else {
		fmt.Println("  Uint64Slice: empty")
	}

	fmt.Println("\nUint64 slice pointer type:")
	if opts.Uint64SlicePtr != nil {
		if len(*opts.Uint64SlicePtr) > 0 {
			fmt.Print("  Uint64SlicePtr: ")
			for i, val := range *opts.Uint64SlicePtr {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Print(val)
			}
			fmt.Println()
		} else {
			fmt.Println("  Uint64SlicePtr: empty slice")
		}
	} else {
		fmt.Println("  Uint64SlicePtr: nil")
	}
}
