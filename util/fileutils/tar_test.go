package fileutils

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

var (
	srcDir = "/tmp/tar_test"
	tarFile   = "/tmp/created.tar"
	dest   = "/tmp/created"

	testString1 = "This is a test #1"
	testString2 = "This is a test #2"
)


func Test_Tar_uncompressed(t *testing.T) {

	if err := createDirAndFiles(); err != nil {
		t.Errorf("Failed to create test environment: %v", err.Error())
	}

	defer cleanup(t)

	buf := new(bytes.Buffer)

	//
	// Tar
	//

	err := Tar(srcDir, buf)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	toWrite, err := os.Create(tarFile)

	defer toWrite.Close()

	if err != nil {
		t.Errorf("%v", err.Error())
	}


	_, err = toWrite.Write(buf.Bytes())

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	//
	// Untar
	//

	toRead, err := os.Open(tarFile)

	if err != nil {
		t.Error("Could not open tar file")
	}

	if err := Untar(dest, toRead); err != nil {
		t.Errorf("Failed to untar file: %v", err.Error())
	}

	validate(t)
}

func Test_Tar_gzip_compressed(t *testing.T) {

	// TODO: only execute this flag when integration test flag is set

	if err := createDirAndFiles(); err != nil {
		t.Errorf("Failed to create test environment: %v", err.Error())
	}

	defer cleanup(t)

	if err := TarGzip(srcDir, tarFile); err != nil {
		t.Errorf("Failed to tar directory: %v", err.Error())
	}

	if err := UntarGzip(dest, tarFile); err != nil {
		t.Errorf("Failed to untar directory: %v", err.Error())
	}

	validate(t)
}

func validate(t *testing.T) {
	sub, err := os.Stat(dest + "/subdir")

	if err != nil {
		t.Error("Could not find directory 'sub'")
	}

	if !sub.IsDir() {
		t.Error("'sub' should be directory but wasn't")
	}

	if !checkFileContents(dest + "/testfile1.txt", testString1) {
		t.Error("File 'testfile1.txt' could not be found or has invalid content.")
	}

	if !checkFileContents(dest + "/subdir/testfile2.txt", testString2) {
		t.Error("File 'testfile1.txt' could not be found or has invalid content.")
	}
}


func checkFileContents(path string, expected string) bool {
	_, err := os.Stat(path)

	if err != nil {
		return false
	}

	content, err := ioutil.ReadFile(path)

	if err != nil {
		return false
	}

	if string(content) != expected {
		return false
	}

	return true
}


func createDirAndFiles() error {
	if err := os.Mkdir("/tmp/tar_test", 0775); err != nil {
		return err
	}

	if err := ioutil.WriteFile("/tmp/tar_test/testfile1.txt", []byte(testString1), 0775); err != nil {
		return err
	}

	if err := os.Mkdir("/tmp/tar_test/subdir", 0775); err != nil {
		return err
	}


	if err := ioutil.WriteFile("/tmp/tar_test/subdir/testfile2.txt", []byte(testString2), 0775); err != nil {
		return err
	}

	return nil
}

func cleanup(t *testing.T) {
	if err := os.RemoveAll(srcDir); err != nil {
		t.Fatalf("Error while cleanup: %v", err.Error())
	}

	if err := os.RemoveAll(dest); err != nil {
		t.Fatalf("Error while cleanup: %v", err.Error())
	}

	if err := os.Remove(tarFile); err != nil {
		t.Fatalf("Error while cleanup: %v", err.Error())
	}
}


