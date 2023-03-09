package filesystem

import (
	"fmt"
	"io"
	"log"
	"os"
)

type DFILE struct {
}

func (d *DFILE) Download(url string) (io.Reader, error) {
	file, err := os.Open(url)
	if err != nil {
		return nil, err
	}
	pr, pw := io.Pipe()
	go func() {
		defer file.Close()
		defer pw.Close()
		size, err := io.Copy(pw, file)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Bytes copied from file to pipe : %v \n", size)
		}
	}()
	return pr, nil
}
