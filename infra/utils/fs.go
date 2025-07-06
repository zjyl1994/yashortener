package utils

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func ExtractDataTo(emfs embed.FS, path string) error {
	err := fs.WalkDir(emfs, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if p == "." {
			return nil
		}

		targetPath := filepath.Join(path, p)

		if d.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		}

		data, err := emfs.ReadFile(p)
		if err != nil {
			return fmt.Errorf("read file failed %s: %v", p, err)
		}

		if err := os.WriteFile(targetPath, data, 0644); err != nil {
			return fmt.Errorf("write file failed %s: %v", targetPath, err)
		}
		return nil
	})
	return err
}
