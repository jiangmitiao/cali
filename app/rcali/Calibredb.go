package rcali

import (
	"github.com/google/uuid"
	"github.com/jiangmitiao/ebook-go"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
)

/**
It is to operate cmd
*/

func hasCalibredb() bool {
	_, err := exec.LookPath("calibredb")
	if err == nil {
		return true
	} else {
		return false
	}
}

func calibredbPath() string {
	if hasCalibredb() {
		str, _ := exec.LookPath("calibredb")
		return str
	}
	return ""
}

func AddBook(bookpath string) (books.Ebook, string) {
	ebook := books.GetEBook(bookpath)
	if ebook != nil {
		bookspath, _ := GetBooksPath()
		filename := path.Join(bookspath, uuid.New().String()+filepath.Ext(bookpath))
		if err := CopyFile(bookpath, filename); err == nil {
			return ebook, filename
		}
	}
	return nil, ""
}

func CopyFile(srcName, dstName string) error {
	src, err := os.Open(srcName)
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	return err
}

func DeleteBook(bookid int) bool {
	if hasCalibredb() {
		cmd := exec.Command(calibredbPath(), "remove", strconv.Itoa(bookid))
		err := cmd.Run()
		if err == nil {
			return cmd.ProcessState.Success()
		}
	} else {
		return false
	}
	return false
}
