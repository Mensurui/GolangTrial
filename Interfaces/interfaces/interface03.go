package interfaces

import (
	"io"
	"os"
)

var w io.WriteCloser

type PageInfo struct {
	pages int
}

func (p PageInfo) Pages() int {
	return p.pages
}

var any interface {
	Pages() int
}

func init() {
	w = os.Stdout
	w.Close()

	pages_number := 4
	pn := PageInfo{pages: pages_number}
	any = pn

}
