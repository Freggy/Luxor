package fileutils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// TarGzip tars a given directory with gzip compression.
func TarGzip(src string) error {
	if _, err := os.Stat(src); err != nil {
		return fmt.Errorf("could not tar files %v", err.Error())
	}

	buf := new(bytes.Buffer)

	gzipw := gzip.NewWriter(buf)

	defer gzipw.Close()


	return Tar(src, gzipw)
}

// UntarGzip untars a
func UntarGzip(dest string, archive string) error {
	f, err :=  os.Open(archive)

	defer f.Close()

	if err != nil {
		return err
	}

	gzipr, err := gzip.NewReader(f)

	defer gzipr.Close()

	return Untar(dest, gzipr)
}


// Tar tars all files in a given source directory together using the specified writer.
func Tar(srcDir string, w io.Writer) error {

	tarw := tar.NewWriter(w)

	defer tarw.Close()

	return filepath.Walk(srcDir, func(file string, fi os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(fi, fi.Name())

		if err != nil {
			return err
		}

		if err := tarw.WriteHeader(header); err != nil {
			return err
		}

		if !fi.Mode().IsRegular() {
			return nil
		}

		f, err := os.Open(file)

		if err != nil {
			return err
		}

		if _, err := io.Copy(tarw, f); err != nil {
			return err
		}

		f.Close()

		return nil
	})
}


// Untar untars
func Untar(dest string, r io.Reader) error {
 	tarr := tar.NewReader(r)

}


