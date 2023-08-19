package main

import(
	"bufio"
	"fmt"
	// "io"
	"os"
	"errors"
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
	ByteCount int
	FileInfo
}
type Chunk struct{
	ChunkCount int
	FileInfo
}

func (l Line) Split() error {
	// ファイル操作
    f, err := os.Open(l.file)
    if err != nil{
        return errors.New("Error: Cannot open file : " + l.file)	
    }
    defer f.Close() // ! important

	scanner := bufio.NewScanner(f)
	for  {
		scanner.Scan()
		fmt.Println(scanner.Text())
	}
	return nil
}

func (b Bynary) Split() error {
	// ファイル操作
    f, err := os.Open(b.file)
    if err != nil{
        return errors.New("Error: Cannot open file : " + b.file)	
    }
    defer f.Close() // ! important

    // バイト型スライスの作成
    buf := make([]byte, b.ByteCount)
    for {
        // nはバイト数を示す
        n, err := f.Read(buf)
        // バイト数が0 = 読み取り終了
        if n == 0{ break }
        if err != nil{ return err }

        // バイト型スライスを文字列型に変換してファイルの内容を出力
        fmt.Println(string(buf[:n]))
		fmt.Println("----------------------------")
    }
	return nil
}

func (n Chunk) Split() error {
	// ファイル操作
    f, err := os.Open(n.file)
    if err != nil{
        return errors.New("Error: Cannot open file : " + b.file)	
    }
    defer f.Close() // ! important



	return nil
}


