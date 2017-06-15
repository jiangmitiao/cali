package rcali

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/revel/revel"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

func GetBooksPath() (string, bool) {
	return revel.Config.String("books.path")
}

func GetSqliteDbPath() (string, bool) {
	return revel.Config.String("sqlitedb.path")
}

func GetUploadPath() (string, bool) {
	return revel.Config.String("books.uploadpath")
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

// Return value if nonempty, def otherwise.
func ValueOrDefault(value, def string) string {
	if value != "" {
		return value
	}
	return def
}

//for debug
type Debug string

func (d Debug) Debug(a ...interface{}) {
	if d == "dev" {
		fmt.Println("Debug:", a)
	}
}

var DEBUG = Debug("")

// Home returns the home directory for the executing user.
//
// This uses an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// cross compile support

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
