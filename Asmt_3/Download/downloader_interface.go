package download

import (
	"dna/Download/filesystem"
	"dna/Download/web"
	"io"
)

type Downloader interface {
	Download(string) (io.Reader, error)
}

func NewDownloader(name string) Downloader {
	switch name {
	case "web":
		{
			return &web.DURL{}
		}
	default:
		{
			return &filesystem.DFILE{}
		}
	}

}
