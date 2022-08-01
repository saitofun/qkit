package main

import (
	"sort"

	g "github.com/saitofun/qkit/gen/codegen"
)

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types
var mimes = map[string]string{
	"text/plain":      "Plain",
	"text/css":        "CSS",
	"text/html":       "HTML",
	"image/gif":       "ImageGIF",
	"image/jpeg":      "ImageJPEG",
	"image/png":       "ImagePNG",
	"image/bmp":       "ImageBmp",
	"image/webp":      "ImageWebp",
	"image/svg+xml":   "ImageSVG",
	"audio/wav":       "AudioWave",
	"audio/midi":      "AudioMidi",
	"audio/webm":      "AudioWebm",
	"video/webm":      "VideoWebm",
	"audio/ogg":       "AudioOgg",
	"video/ogg":       "VideoOgg",
	"application/ogg": "ApplicationOgg",
	"audio/mpeg":      "AudioMp3",
}

func main() {
	{
		f := g.NewFile("httpx", g.GenerateFileSuffix("./httpx_mines.go"))
		keys := make([]string, 0)

		for t := range mimes {
			keys = append(keys, t)
		}
		sort.Strings(keys)

		for _, k := range keys {
			name := mimes[k]
			t := g.Type(name)
			f.WriteSnippet(
				f.Expr(`func New?() *? {return &?{}}`, t, t, t),
				f.Expr(`type ? struct {`+f.Use("bytes", "Buffer")+`}`, t),
				f.Expr(`func (?) ContentType() string {return ?}`, t, f.Value(k)),
			)
		}
		if _, err := f.Write(); err != nil {
			panic(err)
		}

	}

	{
		f := g.NewFile("httpx_test", g.GenerateFileSuffix("./httpx_mines_test.go"))

		for k, name := range mimes {
			t := g.Type(name)
			fn := "New" + name
			f.WriteSnippet(
				f.Expr(`func Example?() {
m := `+f.Use("github.com/saitofun/qkit/kit/httptransport/httpx", fn)+`()
`+f.Use("fmt", "Println")+`(m.ContentType())
// Output:
// `+k+`
}`, t, t),
			)
		}

		if _, err := f.Write(); err != nil {
			panic(err)
		}
	}
}
