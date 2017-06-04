package rcali

import (
	"errors"
	"fmt"
	"github.com/revel/revel"
	"os"
	"path/filepath"
)

func GetBooksPath() (string, bool) {
	return revel.Config.String("books.path")
}

func GetSqliteDbPath() (string, bool) {
	return revel.Config.String("sqlitedb.path")
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return errors.New("nil error")
			}
			if f.IsDir() {
				return errors.New("dir error")
			}
			println(path)
			return nil
		})
		if err != nil {
			return false, err
		}
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

type Debug string

func (d Debug) Debug(a ...interface{}) {
	if d == "dev" {
		fmt.Println("Debug:", a)
	}
}

var DEBUG = Debug("")
