package app

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestEdit(t *testing.T) {
	cwd, _ := os.Getwd()
	path := path.Join(cwd, "assets/meong.txt") 
	fmt.Println(path)

	// Mode banyak ko. Lihat aja os.O_blablabla
	// v Artinya Read-Write + Create. Create klo emang gaada aja!! 
	file, err := os.OpenFile(path, os.O_RDWR | os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	
	defer file.Close() // Tutup klo udah ga kepake!!!

	// file.Chmod(0777)
	file.WriteString("Halo, Gaeesss\n")

}