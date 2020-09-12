package fs

import (
	"fmt"
	"os"
	"path"
)

/**
ensure the file exist
*/
func FileExists(filepath string) (err error) {
	var (
		file *os.File
	)
	fmt.Println("file: " + filepath)
	fmt.Println("dir : " + path.Dir(filepath))
	// ensure dir exist
	if err = IsDir(path.Dir(filepath)); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("create file")
	// ensure file exist
	if _, err = os.Stat(filepath); os.IsNotExist(err) {
		file, err = os.Create(filepath)
		defer func() {
			file.Close()
		}()
	}

	return
}
