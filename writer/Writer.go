package writer

import (
	"fmt"
	"os"
	"ray-tracer/color"
)

type Writer struct {
	Name string
	file *os.File
}

func NewWriter(name string) *Writer {
	writer := Writer{
		Name: name,
	}

	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	writer.file = file

	return &writer
}

func (w *Writer) WriteHeader(width, height int) {
	w.file.WriteString("P3\n")
	w.file.WriteString(fmt.Sprintf("%d %d\n", width, height))
	w.file.WriteString("255\n\n")
}

func (w *Writer) WriteColor(color *color.Color) {
	r, g, b := color.ToRgbBytes()

	w.file.WriteString(fmt.Sprintf("%d %d %d\n", r, g, b))
}

func (w *Writer) Close() {
	w.file.Close()
}
