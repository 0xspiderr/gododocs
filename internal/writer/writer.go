package writer

import (
	"os"
	"path/filepath"
)

func Write(outDir, filename, content string) error {
	if err := os.MkdirAll(outDir, 0777); err != nil {
		return err
	}
	name := filepath.Base(filename)
	finalPath := filepath.Join(outDir, name+".md")
	return os.WriteFile(finalPath, []byte(content), 0777)
}
