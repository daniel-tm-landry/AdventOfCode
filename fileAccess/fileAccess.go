package fileAccess

import (
	//"bufio"
	//"fmt"
	//"io"
	"os"
	"path/filepath"
	"strings"
)

func FileToArrayOnNewLine(filename string) ([]string, error) {
	return FileToArray(filename, "\r\n")
}

func FileToArray(filename string, splitOn string) ([]string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(wd, filename)
	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(dat), splitOn), nil

}
