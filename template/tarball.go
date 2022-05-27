package template

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func untar(reader io.Reader, target string) error {
	tr := tar.NewReader(reader)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		target := header.Name

		// validate name against path traversal
		if !validRelPath(header.Name) {
			return fmt.Errorf("tarball contained invalid name: %q", target)
		}

		target = filepath.Join(target, header.Name)

		switch header.Typeflag {

		// Directory
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}
		// Files
		case tar.TypeReg:
			fileToWrite, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			if _, err := io.Copy(fileToWrite, tr); err != nil {
				return err
			}
			fileToWrite.Close()
		}
	}

	//
	return nil
}

func validRelPath(p string) bool {
	if p == "" || strings.Contains(p, `\`) || strings.HasPrefix(p, "/") || strings.Contains(p, "../") {
		return false
	}
	return true
}
