package sfzdemo

// Import the testing helpers
import (
	"testing"

	fuzztest "git.fuzzbuzz.io/fuzz/testing"
)

func TestFuzzStringAPI(t *testing.T) {
	f := fuzztest.NewChecker(t)

	// Run 1000 random inputs to test interop:
	fuzztest.Randomize(f, FuzzMyStringAPI, 1000)
}
