package main

import(
	"fmt"
	// "io"
	// "os"
)

type FileInfo struct {
	file string
	prefix string
	saffix int
}

type Line struct{
	LineCount int
	FileInfo
}
type Bynary struct{
	ByteCount string
	FileInfo
}
type Chunk struct{
	ChunkCount int
	FileInfo
}

func (l Line) Split() string {
	return fmt.Sprintf("%d", l.LineCount)
}

func (b Bynary) Split() string {
	return fmt.Sprintf("%s", b.file[:6])
}

func (n Chunk) Split() string {
	return fmt.Sprint("hai, donmai!")
}

func readFile() error {
	return nil
}

