package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// :quickclop
type Options struct {
	// Basic types
	String     string        `clop:"-s,--string" usage:"A string value" default:"default string"`
	Int        int           `clop:"-i,--int" usage:"An integer value" default:"42"`
	Float      float64       `clop:"-f,--float" usage:"A float value" default:"3.14"`
	Bool       bool          `clop:"-b,--bool" usage:"A boolean flag"`
	Duration   time.Duration `clop:"-d,--duration" usage:"A duration value" default:"1h30m"`
	
	// Pointer types
	StringPtr  *string       `clop:"--string-ptr" usage:"A string pointer value" default:"default ptr"`
	IntPtr     *int          `clop:"--int-ptr" usage:"An integer pointer value" default:"100"`
	FloatPtr   *float64      `clop:"--float-ptr" usage:"A float pointer value" default:"2.718"`
	BoolPtr    *bool         `clop:"--bool-ptr" usage:"A boolean pointer flag" default:"true"`
	
	// Slice types
	StringSlice []string     `clop:"--strings" usage:"A string slice" default:"a,b,c"`
	
	// Environment variable
	EnvValue    string       `clop:"--env" usage:"Value from environment" env:"QUICKCLOP_TEST_ENV"`
	
	// Help flag
	Help       bool          `clop:"-h,--help" usage:"Show help information"`
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
	fmt.Println("Basic types:")
	fmt.Printf("  String: %q\n", opts.String)
	fmt.Printf("  Int: %d\n", opts.Int)
	fmt.Printf("  Float: %f\n", opts.Float)
	fmt.Printf("  Bool: %v\n", opts.Bool)
	fmt.Printf("  Duration: %v\n", opts.Duration)
	
	fmt.Println("\nPointer types:")
	if opts.StringPtr != nil {
		fmt.Printf("  StringPtr: %q\n", *opts.StringPtr)
	}
	if opts.IntPtr != nil {
		fmt.Printf("  IntPtr: %d\n", *opts.IntPtr)
	}
	if opts.FloatPtr != nil {
		fmt.Printf("  FloatPtr: %f\n", *opts.FloatPtr)
	}
	if opts.BoolPtr != nil {
		fmt.Printf("  BoolPtr: %v\n", *opts.BoolPtr)
	}
	
	fmt.Println("\nSlice types:")
	fmt.Printf("  StringSlice: %v\n", opts.StringSlice)
	
	fmt.Println("\nEnvironment variable:")
	fmt.Printf("  EnvValue: %q\n", opts.EnvValue)
}
