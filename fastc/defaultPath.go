package fastc

import (
	"github.com/spf13/afero"

	"github.com/tsingson/bytecache/pkg/vtils"
)

// defaultLogPath
func defaultLogPath(path string) string {
	var err error
	p, _ := vtils.GetCurrentExecDir()

	logPath := p + "/" + path

	afs := afero.NewOsFs()
	check, _ := afero.DirExists(afs, logPath)
	if !check {
		err = afs.MkdirAll(logPath, 0755)
		if err != nil {
			panic("can't make path" + logPath)
		}
	}

	return logPath
}
