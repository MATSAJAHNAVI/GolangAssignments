package zipper

import (
	"archive/zip"
	"io"
	"log"
)

type Ziparchiver struct {
}

func (z *Ziparchiver) Archive(names []string, readers ...io.Reader) (io.Reader, error) {
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		zipWriter := zip.NewWriter(pw)
		for n, reader := range readers {
			filename := names[n]
			f, _ := zipWriter.Create(filename)

			_, err := io.Copy(f, reader)
			if err != nil {
				log.Printf("error occured in copying file")
			}
		}
		err := zipWriter.Close()
		if err != nil {
			log.Printf("error in closing zip file")
		}
	}()
	return pr, nil
}

// 	fmt.Println("creating archive file")
// 	archive, err := os.Create("archive.zip")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer archive.Close()

// 	//zipWriter := zip.NewWriter(archive)

// 	fmt.Println("opening first file")
// 	f1, err := os.Open(file1)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f1.Close()

// 	fmt.Println("writing to first file")
// 	w1, err := zipWriter.Create(file1)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if _, err := io.Copy(w1, f1); err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("opening second file")
// 	f2, err := os.Open(file2)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f2.Close()

// 	fmt.Println("writing second file to the archive")
// 	w2, err := zipWriter.Create(file2)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if _, err := io.Copy(w2, f2); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("closing zip archive")
// 	zipWriter.Close()

// 	//pr, pw := io.Pipe()
// 	io.Copy(pw, archive)
// 	defer pw.Close()
// 	return pr, err

// }
