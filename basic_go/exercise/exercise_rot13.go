package exercise

import (
	"fmt"
	"image"
	"io"
	"os"
	"strings"
)

func exercise_rot13() {
	fmt.Println("Exercise: rot13 Reader")
	s := strings.NewReader("Lbh penpxrq gur phqr!")
	rot := &rot13Reader{s}
	/*
		func Copy(dst Writer, src Reader) (written int64, err error)
		If src implements the WriterTo interface, the copy is implemented by calling src.WriteTo(dst).
		Otherwise, if dst implements the ReaderFrom interface, the copy is implemented by calling dst.ReadFrom(src).
	*/
	io.Copy(os.Stdout, rot)

	fmt.Println()
	fmt.Println()
	fmt.Println("Images")
	im := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(im.Bounds())
	fmt.Println(im.At(0, 0).RGBA())

	fmt.Println()
}

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	for i, c := range p {
		switch {
		case c >= 'A' && c <= 'M' || c >= 'a' && c <= 'm':
			p[i] += 13
		case c >= 'N' && c <= 'Z' || c >= 'n' && c <= 'z':
			p[i] -= 13
		}
	}
	return
}
