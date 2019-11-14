package utils

import (
	"archive/zip"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const (
	testTmpZipDirPath  = testTmpRootDirPath + "/zip"
	testTmpZipDirPath1 = testTmpRootDirPath + "/zip/1"
	testTmpZipDirPath2 = testTmpRootDirPath + "/zip/2"
)

func TestZipT_ArchivesRecursive(t *testing.T) {
	t.Helper()

	testTargetPaths := []string{testTmpZipDirPath + `/file`, testTmpZipDirPath1, testTmpZipDirPath2}

	// script->
	testTnitialScript := `
rm -rf ` + testTmpZipDirPath + `
mkdir -p ` + testTmpZipDirPath1 + `/dir
mkdir -p ` + testTmpZipDirPath2 + `/dir
touch ` + testTmpZipDirPath + `/file
touch ` + testTmpZipDirPath1 + `/file
touch ` + testTmpZipDirPath2 + `/file
` // <-script

	initTest := func() {
		out, err := exec.Command("/bin/bash", "-c", testTnitialScript).CombinedOutput()
		if err != nil {
			Log.Fatalfln("TestZipT_ArchivesRecursive(): %v, %v", strings.TrimSuffix(string(out), "\n"), err)
		}
	}

	t.Run("non-normal/filepathAbsZipFn", func(t *testing.T) {
		initTest()
		failureZIP := ZIP
		failureZIP.filepathAbsZipFn = func(path string) (s string, e error) {
			return path, ErrorDummyErrorForTest
		}
		if err := failureZIP.ArchivesRecursive(testTmpZipDirPath+"/failure.zip", testTargetPaths, false); err == nil {
			t.Errorf("TestZipT_ArchivesRecursive(): ZIP.ArchivesRecursive(): err == nil: %v", err)
		}
	})

	t.Run("non-normal/filepathAbsTargetFn", func(t *testing.T) {
		initTest()
		failureZIP := ZIP
		failureZIP.filepathAbsTargetFn = func(path string) (s string, e error) {
			return path, ErrorDummyErrorForTest
		}
		if err := failureZIP.ArchivesRecursive(testTmpZipDirPath+"/failure.zip", testTargetPaths, false); err == nil {
			t.Errorf("TestZipT_ArchivesRecursive(): ZIP.ArchivesRecursive(): err == nil: %v", err)
		}
	})

	t.Run("non-normal/already_exists", func(t *testing.T) {
		initTest()
		failureZIP := ZIP
		if err := failureZIP.ArchivesRecursive("/", testTargetPaths, false); err == nil {
			t.Errorf("TestZipT_ArchivesRecursive(): ZIP.ArchivesRecursive(): err == nil: %v", err)
		}
	})

	t.Run("non-normal/ioCopyFn", func(t *testing.T) {
		initTest()
		failureZIP := ZIP
		failureZIP.ioCopyFn = func(dst io.Writer, src io.Reader) (written int64, err error) {
			return 0, ErrorDummyErrorForTest
		}
		if err := failureZIP.ArchivesRecursive(testTmpZipDirPath+"/failure.zip", testTargetPaths, false); err == nil {
			t.Errorf("TestZipT_ArchivesRecursive(): ZIP.ArchivesRecursive(): err == nil: %v", err)
		}
	})

	t.Run("non-normal/osOpenFn", func(t *testing.T) {
		initTest()
		failureZIP := ZIP
		failureZIP.osOpenFn = func(name string) (*os.File, error) {
			return nil, ErrorDummyErrorForTest
		}
		if err := failureZIP.ArchivesRecursive(testTmpZipDirPath+"/failure.zip", testTargetPaths, false); err == nil {
			t.Errorf("TestZipT_ArchivesRecursive(): ZIP.ArchivesRecursive(): err == nil: %v", err)
		}
	})

	t.Run("non-normal/zipWriterCreateFn", func(t *testing.T) {
		initTest()
		failureZIP := ZIP
		failureZIP.zipWriterCreateFn = func(w *zip.Writer, name string) (io.Writer, error) {
			return nil, ErrorDummyErrorForTest
		}
		if err := failureZIP.ArchivesRecursive(testTmpZipDirPath+"/failure.zip", testTargetPaths, false); err == nil {
			t.Errorf("TestZipT_ArchivesRecursive(): ZIP.ArchivesRecursive(): err == nil: %v", err)
		}
	})

	t.Run("non-normal/osCreateFn", func(t *testing.T) {
		initTest()
		failureZIP := ZIP
		failureZIP.osCreateFn = func(name string) (*os.File, error) {
			return nil, ErrorDummyErrorForTest
		}
		if err := failureZIP.ArchivesRecursive(testTmpZipDirPath+"/failure.zip", testTargetPaths, false); err == nil {
			t.Errorf("TestZipT_ArchivesRecursive(): ZIP.ArchivesRecursive(): err == nil: %v", err)
		}
	})

	t.Run("non-normal/funcForWalkFunc", func(t *testing.T) {
		initTest()
		if err := funcForWalkFunc(nil, nil, "", "", nil, ErrorDummyErrorForTest); err == nil {
			t.Error("TestZipT_ArchivesRecursive(): funcForWalkFunc(): err == nil")
		}
	})

	t.Run("normal/withoutRootDirectory=true", func(t *testing.T) {
		initTest()
		if err := ZIP.ArchivesRecursive(testTmpZipDirPath+"/withRoot.zip", testTargetPaths, true); err != nil {
			t.Errorf("TestZipT_ArchivesRecursive(): ZIP.ArchivesRecursive(): err != nil: %v", err)
		}
	})

	t.Run("normal/withoutRootDirectory=false", func(t *testing.T) {
		initTest()
		if err := ZIP.ArchivesRecursive(testTmpZipDirPath+"/withoutRoot.zip", testTargetPaths, false); err != nil {
			t.Errorf("TestZipT_ArchivesRecursive(): ZIP.ArchivesRecursive(): err != nil: %v", err)
		}
	})
}
