package web

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type DURL struct {
}

func (d *DURL) Download(url string) (io.Reader, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	pr, pw := io.Pipe()
	go func() {
		defer response.Body.Close()
		defer pw.Close()
		n, err := io.Copy(pw, response.Body)
		if err != nil {
			log.Printf("error occured while copying : %v", err)
		}
		fmt.Printf("%d bytes copied from web url content\n", n)
	}()
	//defer response.Body.Close()
	return pr, err

}
