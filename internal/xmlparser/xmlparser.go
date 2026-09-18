package xmlparser

import (
	"encoding/xml"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	m "github.com/0xspiderr/gododocs/internal/model"
)

// purpose: Walks the directory and subdirectories of the given path
// and finds files with .xml suffix, reads those files, unmarshalls them
// in the given Class structure and returns a list of classes for the
// renderer to handle.
func Parse(path string) ([]m.Class, error) {
	var classes []m.Class
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".xml") {
			return nil
		}
		f, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		var c m.Class
		if err := xml.Unmarshal(f, &c); err != nil {
			return err
		}
		classes = append(classes, c)
		return nil
	})
	return classes, err
}
