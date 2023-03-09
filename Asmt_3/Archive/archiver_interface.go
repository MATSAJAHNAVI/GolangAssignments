package archive

import (
	zipper "dna/Archive/zip"
	"io"
)

type Archiver interface {
	Archive([]string, ...io.Reader) (io.Reader, error)
}

func New() Archiver {
	//return &ziparchive.Ziparchiver{}
	//return &zip.ziparchive{}
	return &zipper.Ziparchiver{}
}
