package files

import (
	"fmt"
	"github.com/slham/toolbelt/l"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(workingDirectory, fileName string) ([]byte, error) {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, workingDirectory) {
		wd = filepath.Dir(wd)
	}
	path := fmt.Sprintf("%s/%s", wd, fileName)
	l.Debug(nil, "path:%s", path)
	return ioutil.ReadFile(path)
}
