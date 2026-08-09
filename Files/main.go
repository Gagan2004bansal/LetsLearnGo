package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("is Folder:", fileInfo.IsDir())
	fmt.Println("File Size:", fileInfo.Size(), "bytes")
	fmt.Println("File permission:", fileInfo.Mode())
	fmt.Println("Last Modified:", fileInfo.ModTime())

	buf := make([]byte, fileInfo.Size())

	d, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	println("File content:", string(buf[:d]))

	// Other method to Read File
	data, err := os.ReadFile("example.txt") // Good for small files
	if err != nil {
		panic(err)
	}
	fmt.Println("File content:", string(data))

	// File Creation & Write Operation
	newFile, err := os.Create("newfile.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// newFile.WriteString("Hi GoLang\n")

	// Second type to write file
	bytes := []byte("Hi GoLang\n")
	newFile.Write(bytes)

	// Read & Write to another file (streaming fashion)
	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("dest.txt")
	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		bytee, err := reader.ReadByte()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		e := writer.WriteByte(bytee)
		if e != nil {
			panic(e)
		}
	}

	writer.Flush()
	fmt.Println("File copied successfully")

	// Delete a file
	errror := os.Remove("newfile.txt")
	if errror != nil {
		panic(errror)
	}
}
