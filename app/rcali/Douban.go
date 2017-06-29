package rcali

import (
	"net/http"
	"io/ioutil"
)

func GetDoubanInfoById(id string)string  {
	if id == "" {
		return ""
	}

	resp, err := http.Get("https://api.douban.com/v2/book/" + id)
	if err != nil {
		// handle error
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return ""
	}
	return string(body)
}