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
	"strings"
	"mime/multipart"
	"io/ioutil"
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

func WriteBook(file multipart.File,bookfilepath string)error  {
	defer file.Close()
	b, _ := ioutil.ReadAll(file)
	if len(b)==0 {
		Logger.Error("==== upload error size is 0 "+bookfilepath)
	}
	err :=ioutil.WriteFile(bookfilepath, b, 0755)
	if err != nil {
		os.Remove(bookfilepath)
	}
	return err
}

func GetRealBookInfo(bookfilepath string) (books.Ebook, bool) {
	ebook := books.GetEBook(bookfilepath)
	if ebook == nil {
		Logger.Error("==== can not parse "+bookfilepath)
		return nil, false
	} else {
		return ebook, true
	}
}

func AddBook(bookfilepath string) (books.Ebook, string) {
	ebook := books.GetEBook(bookfilepath)
	if ebook != nil {
		bookspath, _ := GetBooksPath()
		filename := path.Join(bookspath, uuid.New().String()+filepath.Ext(bookfilepath))
		if err := CopyFile(bookfilepath, filename); err == nil {
			os.Remove(bookfilepath)
			return ebook, filename
		}
	}
	Logger.Error("==== can not parse "+bookfilepath)
	os.Remove(bookfilepath)
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

func DeleteRealBook(bookpath string)  {
	os.Remove(bookpath)
}

func DeleteTmpBook()  {
	uploadbookdir ,_:=GetUploadPath()
	uploadbookdir = path.Join(uploadbookdir)
	filepath.Walk(uploadbookdir, func(path string, info os.FileInfo, err error) error {
		if (strings.ToLower(filepath.Ext(info.Name())) == ".epub" || strings.ToLower(filepath.Ext(info.Name())) == ".mobi") && filepath.Dir(path)==uploadbookdir{
			os.Remove(filepath.Join(uploadbookdir,info.Name()))
		}
		return nil
	})
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
