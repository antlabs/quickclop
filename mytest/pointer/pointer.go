package pointer

// testPointers represents a struct with pointer fields for testing.
// :quickclop
type testPointers struct {
	StringPtr  *string  `clop:"-s; --string" usage:"A string pointer"`
	IntPtr     *int     `clop:"-i; --int" usage:"An integer pointer"`
	FloatPtr   *float64 `clop:"-f; --float" usage:"A float pointer"`
	BoolPtr    *bool    `clop:"-b; --bool" usage:"A boolean pointer"`
	StringArg  *string  `clop:"args=stringarg" usage:"A string argument as pointer"`
}
