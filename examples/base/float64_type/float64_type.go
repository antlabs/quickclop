package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Float64TypeOptions struct {
	// Basic float64 type
	Float64 float64 `clop:"-f,--float64" usage:"A float64 value" default:"3.14159265359"`

	// Pointer float64 type
	Float64Ptr *float64 `clop:"--float64-ptr" usage:"A float64 pointer value" default:"2.71828182846"`

	// Slice of float64 type
	Float64Slice []float64 `clop:"--float64-slice" usage:"A slice of float64 values" default:"1.1,2.2,3.3"`

	// Pointer to slice of float64 type
	Float64SlicePtr *[]float64 `clop:"--float64-slice-ptr" usage:"A pointer to a slice of float64 values" default:"4.4,5.5,6.6"`

	// Help flag
	Help bool `clop:"-h,--help" usage:"Show help information"`
}

func main() {
	var opts Float64TypeOptions
	err := opts.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if opts.Help {
		opts.Usage()
		return
	}

	// Display all values
	fmt.Println("Float64 type:")
	fmt.Printf("  Float64: %f\n", opts.Float64)

	fmt.Println("\nFloat64 pointer type:")
	if opts.Float64Ptr != nil {
		fmt.Printf("  Float64Ptr: %f\n", *opts.Float64Ptr)
	} else {
		fmt.Println("  Float64Ptr: nil")
	}

	fmt.Println("\nFloat64 slice type:")
	fmt.Printf("  Float64Slice: %v\n", opts.Float64Slice)

	fmt.Println("\nFloat64 slice pointer type:")
	if opts.Float64SlicePtr != nil {
		fmt.Printf("  Float64SlicePtr: %v\n", *opts.Float64SlicePtr)
	} else {
		fmt.Println("  Float64SlicePtr: nil")
	}
}
