package rcali

import (
	"os/exec"
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

func AddBook(bookpath string) bool {
	if hasCalibredb() {
		cmd := exec.Command(calibredbPath(), "add", bookpath)
		err := cmd.Run()
		if err == nil {
			return cmd.ProcessState.Success()
		}
	} else {
		return false
	}
	return false
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
