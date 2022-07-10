package app

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestCreate(t *testing.T) {
	cwd, _ := os.Getwd()
	path := path.Join(cwd, "assets/meong.txt") 
	fmt.Println(path)

	_, err := os.Stat(path)

	if(os.IsNotExist(err)) {
		fmt.Println("Buat File Baru")
		file, err := os.Create(path)
		
		if err != nil {
			fmt.Println(err.Error())
		}

		file.WriteString("Halo, Dunia\n")

		file.Close()
	}
}
