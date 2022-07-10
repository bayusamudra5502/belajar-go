package app

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestDelete(t *testing.T) {
	cwd, _ := os.Getwd()
	path := path.Join(cwd, "assets/meong.txt") 
	fmt.Println(path)

	os.Remove(path)
}