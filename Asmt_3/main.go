package main

import (
	archive "dna/Archive"
	download "dna/Download"
	"fmt"
	"io"
	"os"
)

func main() {
	//url1 := "https://in.pinterest.com/pin/602286150160436158/"
	//url2 := "https://in.pinterest.com/pin/712061391077012491/"
	//url1 := "http://www.yahoo.com/image_to_read.jpg"
	//url2 := "https://images.app.goo.gl/CiFfz13ZcWHzAhGs8"

	d := download.NewDownloader("web")
	url1 := "https://images.app.goo.gl/EYuUQFuyFLpZPLHE9"
	r1, _ := d.Download(url1)

	fd := download.NewDownloader("filesystem")
	url2 := "/Users/matsa.jahnavi/Downloads/manali.jpeg"
	r2, _ := fd.Download(url2)

	zipper := archive.New()
	zipR, _ := zipper.Archive([]string{"f1.jpg", "f2.jpg"}, r1, r2)
	zipW, _ := os.Create("result.zip")
	n, _ := io.Copy(zipW, zipR)

	fmt.Printf("%d bytes ", n)
	//fmt.Printf("%v error occured", err)

}
