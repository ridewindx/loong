package fs

import (
	"github.com/spf13/afero"
)

type File interface {
	afero.File
}

type Fs struct {
	*afero.Afero
}

func NewFs() *Fs {
	return &Fs{
		Afero: &afero.Afero{
			Fs: afero.NewOsFs(),
		},
	}
}
