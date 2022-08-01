// This is a generated source file. DO NOT EDIT
// Source: httpx_test/httpx_mines__generated_test.go

package httpx_test

import (
	"fmt"

	"github.com/saitofun/qkit/kit/httptransport/httpx"
)

func ExampleAudioWebm() {
	m := httpx.NewAudioWebm()
	fmt.Println(m.ContentType())
	// Output:
	// audio/webm
}

func ExampleAudioOgg() {
	m := httpx.NewAudioOgg()
	fmt.Println(m.ContentType())
	// Output:
	// audio/ogg
}

func ExampleAudioMp3() {
	m := httpx.NewAudioMp3()
	fmt.Println(m.ContentType())
	// Output:
	// audio/mpeg
}

func ExamplePlain() {
	m := httpx.NewPlain()
	fmt.Println(m.ContentType())
	// Output:
	// text/plain
}

func ExampleCSS() {
	m := httpx.NewCSS()
	fmt.Println(m.ContentType())
	// Output:
	// text/css
}

func ExampleImageGIF() {
	m := httpx.NewImageGIF()
	fmt.Println(m.ContentType())
	// Output:
	// image/gif
}

func ExampleImageSVG() {
	m := httpx.NewImageSVG()
	fmt.Println(m.ContentType())
	// Output:
	// image/svg+xml
}

func ExampleImageJPEG() {
	m := httpx.NewImageJPEG()
	fmt.Println(m.ContentType())
	// Output:
	// image/jpeg
}

func ExampleAudioWave() {
	m := httpx.NewAudioWave()
	fmt.Println(m.ContentType())
	// Output:
	// audio/wav
}

func ExampleVideoWebm() {
	m := httpx.NewVideoWebm()
	fmt.Println(m.ContentType())
	// Output:
	// video/webm
}

func ExampleApplicationOgg() {
	m := httpx.NewApplicationOgg()
	fmt.Println(m.ContentType())
	// Output:
	// application/ogg
}

func ExampleHTML() {
	m := httpx.NewHTML()
	fmt.Println(m.ContentType())
	// Output:
	// text/html
}

func ExampleImagePNG() {
	m := httpx.NewImagePNG()
	fmt.Println(m.ContentType())
	// Output:
	// image/png
}

func ExampleImageBmp() {
	m := httpx.NewImageBmp()
	fmt.Println(m.ContentType())
	// Output:
	// image/bmp
}

func ExampleImageWebp() {
	m := httpx.NewImageWebp()
	fmt.Println(m.ContentType())
	// Output:
	// image/webp
}

func ExampleAudioMidi() {
	m := httpx.NewAudioMidi()
	fmt.Println(m.ContentType())
	// Output:
	// audio/midi
}

func ExampleVideoOgg() {
	m := httpx.NewVideoOgg()
	fmt.Println(m.ContentType())
	// Output:
	// video/ogg
}
