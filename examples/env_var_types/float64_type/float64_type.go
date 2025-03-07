package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Options struct {
	// Basic float64 type with environment variable
	Float64 float64 `clop:"-f,--float64" usage:"A float64 value" env:"FLOAT64_VALUE"`

	// Pointer float64 type with environment variable
	Float64Ptr *float64 `clop:"--float64-ptr" usage:"A float64 pointer value" env:"FLOAT64_PTR_VALUE"`

	// Slice of float64 type with environment variable
	Float64Slice []float64 `clop:"--float64-slice" usage:"A slice of float64 values" env:"FLOAT64_SLICE_VALUE"`

	// Pointer to slice of float64 type with environment variable
	Float64SlicePtr *[]float64 `clop:"--float64-slice-ptr" usage:"A pointer to a slice of float64 values" env:"FLOAT64_SLICE_PTR_VALUE"`

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
	fmt.Println("Float64 type with environment variable:")
	fmt.Printf("  Float64: %f\n", opts.Float64)

	fmt.Println("\nFloat64 pointer type with environment variable:")
	if opts.Float64Ptr != nil {
		fmt.Printf("  Float64Ptr: %f\n", *opts.Float64Ptr)
	} else {
		fmt.Println("  Float64Ptr: nil")
	}

	fmt.Println("\nFloat64 slice type with environment variable:")
	fmt.Printf("  Float64Slice: %v\n", opts.Float64Slice)

	fmt.Println("\nFloat64 slice pointer type with environment variable:")
	if opts.Float64SlicePtr != nil {
		fmt.Printf("  Float64SlicePtr: %v\n", *opts.Float64SlicePtr)
	} else {
		fmt.Println("  Float64SlicePtr: nil")
	}
}
