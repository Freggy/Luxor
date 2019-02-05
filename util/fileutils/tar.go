package fileutils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// TarGzip tars a given directory with gzip compression.
func TarGzip(src string, dest string) error {
	if _, err := os.Stat(src); err != nil {
		return fmt.Errorf("could not tar files %v", err.Error())
	}

	f, err := os.Create(dest)

	defer f.Close()

	if err != nil {
		return err
	}


	gzipw := gzip.NewWriter(f)

	defer gzipw.Close()

	return Tar(src, gzipw)
}

// UntarGzip untars a gzip compressed tar archive to the given destination.
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

		header.Name = strings.TrimPrefix(strings.Replace(file, srcDir, "", -1), string(filepath.Separator))

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


// Untar untars files using the specified reader to the given destination.
func Untar(dest string, r io.Reader) error {
 	tarr := tar.NewReader(r)

	for {
		header, err := tarr.Next()

		switch {

		// EOF means no more files are found so return
		case err == io.EOF:
			return nil

		case err != nil:
			return err

		case header == nil:
			continue
		}

		target := filepath.Join(dest, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}
		case tar.TypeReg: // TypeReg -> file
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))

			if err != nil {
				return err
			}

			if _, err := io.Copy(f, tarr); err != nil {
				return err
			}

			f.Close()
		}
	}
}


