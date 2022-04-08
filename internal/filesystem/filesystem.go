package filesystem

import "github.com/spf13/afero"

type FileSystemManager struct {
	Fs   afero.Fs
	Util *afero.Afero
}

var (
	mgr *FileSystemManager
)

func NewFileSystemManager(fs afero.Fs) *FileSystemManager {
	return &FileSystemManager{fs, &afero.Afero{Fs: fs}}
}

func Mgr() *FileSystemManager {
	if mgr == nil {
		mgr = NewFileSystemManager(afero.NewOsFs())
	}
	return mgr
}
