package app

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestMkdir(t *testing.T) {
	cwd, _ := os.Getwd()
	path := path.Join(cwd, "assets/Hana") 
	fmt.Println(path)

	os.Mkdir(path, 0644);
}