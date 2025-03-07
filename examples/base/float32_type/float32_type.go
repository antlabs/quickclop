package main

import (
	"fmt"
	"log"
	"os"
)

// :quickclop
type Float32TypeOptions struct {
	// Basic float32 type
	Float32 float32 `clop:"-f,--float32" usage:"A float32 value" default:"3.14"`

	// Pointer float32 type
	Float32Ptr *float32 `clop:"--float32-ptr" usage:"A float32 pointer value" default:"2.718"`

	// Slice of float32 type
	Float32Slice []float32 `clop:"--float32-slice" usage:"A slice of float32 values" default:"1.1,2.2,3.3"`

	// Pointer to slice of float32 type
	Float32SlicePtr *[]float32 `clop:"--float32-slice-ptr" usage:"A pointer to a slice of float32 values" default:"4.4,5.5,6.6"`

	// Help flag
	Help bool `clop:"-h,--help" usage:"Show help information"`
}

func main() {
	var opts Float32TypeOptions
	err := opts.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if opts.Help {
		opts.Usage()
		return
	}

	// Display all values
	fmt.Println("Float32 type:")
	fmt.Printf("  Float32: %f\n", opts.Float32)

	fmt.Println("\nFloat32 pointer type:")
	if opts.Float32Ptr != nil {
		fmt.Printf("  Float32Ptr: %f\n", *opts.Float32Ptr)
	} else {
		fmt.Println("  Float32Ptr: nil")
	}

	fmt.Println("\nFloat32 slice type:")
	fmt.Printf("  Float32Slice: %v\n", opts.Float32Slice)

	fmt.Println("\nFloat32 slice pointer type:")
	if opts.Float32SlicePtr != nil {
		fmt.Printf("  Float32SlicePtr: %v\n", *opts.Float32SlicePtr)
	} else {
		fmt.Println("  Float32SlicePtr: nil")
	}
}
