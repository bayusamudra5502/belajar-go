package app

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestBaca(t *testing.T) {
	cwd, _ := os.Getwd()
	pathFile := path.Join(cwd, "assets/meong.txt") 
	fmt.Println(pathFile)

	// Mode banyak ko. Lihat aja os.O_blablabla
	// v Artinya Read-Write + Create. Create klo emang gaada aja!! 
	fileFile, err := os.OpenFile(pathFile, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	defer fileFile.Close()

	data := [64]byte{}
	fileFile.Read(data[:])

	fmt.Println(string(data[:]))

	// CARA II:
	// BUKA FILE
	pathFile = path.Join(cwd, "assets/file.txt")
	fileFile2, err := os.Open(pathFile) // Ini otomatis Readonly

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	defer fileFile2.Close()

	fileFile2.Read(data[:])

	fmt.Println(string(data[:]))
}