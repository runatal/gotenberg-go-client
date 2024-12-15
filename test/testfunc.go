// Package test contains useful functions used across tests.
package test

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// HTMLTestFilePath returns the absolute file path of a file in "html" folder in test/testdata.
func HTMLTestFilePath(t *testing.T, filename string) string {
	return abs(t, "html", filename)
}

// MarkdownTestFilePath returns the absolute file path of a file in "markdown" folder in test/testdata.
func MarkdownTestFilePath(t *testing.T, filename string) string {
	return abs(t, "markdown", filename)
}

// OfficeTestFilePath returns the absolute file path of a file in "office" folder in test/testdata.
func OfficeTestFilePath(t *testing.T, filename string) string {
	return abs(t, "office", filename)
}

// PDFTestFilePath returns the absolute file path of a file in "pdf" folder in test/testdata.
func PDFTestFilePath(t *testing.T, filename string) string {
	return abs(t, "pdf", filename)
}

func abs(t *testing.T, kind, filename string) string {
	_, gofilename, _, ok := runtime.Caller(0)
	require.True(t, ok, "got no caller information")

	if filename == "" {
		fpath, err := filepath.Abs(fmt.Sprintf("%s/testdata/%s", path.Dir(gofilename), kind))
		require.NoErrorf(t, err, `getting the absolute path of "%s"`, kind)

		return fpath
	}

	fpath, err := filepath.Abs(fmt.Sprintf("%s/testdata/%s/%s", path.Dir(gofilename), kind, filename))
	require.NoErrorf(t, err, `getting the absolute path of "%s"`, filename)

	return fpath
}

// IsPDF checks if the given file is a PDF file by looking for the PDF header.
func IsPDF(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	buffer := make([]byte, 5)
	_, err = file.Read(buffer)
	if err != nil {
		return false, err
	}

	if bytes.Equal(buffer, []byte("%PDF-")) {
		return true, nil
	}

	return false, nil
}

// IsPDFA checks if the given PDF file is PDF/A compliant.
func IsPDFA(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "pdfaid:part") {
			return true, nil
		}
	}

	return false, nil
}

// IsPDFUA checks if the given PDF file is PDF/UA compliant.
func IsPDFUA(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "pdfuaid:part") {
			return true, nil
		}
	}

	return false, nil
}
