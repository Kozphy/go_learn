package exercise

import (
	"fmt"
	"io"
	"math"
	"strings"
)

func exercise_error() {
	fmt.Println("Exercise: Errors")
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	fmt.Println()
	fmt.Println("Reader")
	r := strings.NewReader("Hello, Reader")

	b := make([]byte, 8)
	for {
		/*
			type Reader interface {
				Read(p []byte) (n int, err error)
			}
			Read reads up to len(p) bytes into p.
			It returns the number of bytes read (0 <= n <= len(p)) and any error encountered.
			Even if Read returns n < len(p), it may use all of p as scratch space during the call

			When Read encounters an error or end-of-file condition after successfully reading n > 0 bytes, it returns the number of bytes read and
			Reader may return either err == EOF or err == nil. The next Read should return 0, EOF.
		*/
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
	fmt.Println()

}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	return math.Sqrt(x), nil

}
