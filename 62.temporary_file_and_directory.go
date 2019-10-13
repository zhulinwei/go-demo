package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main () {
	file1, err := ioutil.TempFile("./", "samle1")

	fmt.Println(err)
	fmt.Println(file1.Name())

	_, err = file1.Write([]byte{1, 2, 3, 4})
	dir, err := ioutil.TempDir("./", "samle2")
	fmt.Println(dir)

	file2 := filepath.Join(dir, "file2")
	_ = ioutil.WriteFile(file2, []byte{1, 2}, 0666)


	_ = os.RemoveAll(file1.Name())
	_ = os.RemoveAll(dir)

}