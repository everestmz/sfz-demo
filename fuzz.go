package sfzdemo

// Import the fuzz package
import (
	"git.fuzzbuzz.io/fuzz"
)

// Put fuzz tests in a <something>_fuzz.go file

func FuzzMyStringAPI(f *fuzz.F) {
	str := f.String("myString").Get()
	MyStringAPI(str)
}

func FuzzMyBytesAPI(f *fuzz.F) {
	// We can also configure the data the fuzzer generates
	bs := f.Bytes("bs").
		Seeds([]byte("seedthefuzzer")). // Add seed inputs
		Range(10, 100)                  // Configure the max and min length of the byte array
	MyBytesAPI(bs.Get())
}

func FuzzDoubleStringAPI(f *fuzz.F) {
	// You can create more than one of a type by using different tags
	MyDoubleStringAPI(f.String("s1").Get(), f.String("s2").Get())
}

func FuzzPrimitive(f *fuzz.F) {
	i := f.Int("i").Get()
	u := f.Uint64("u").Get()
	f32 := f.Float32("f").Get()
	b := f.Bool("b").Get()
	MyPrimitiveAPI(i, u, f32, b)
}
