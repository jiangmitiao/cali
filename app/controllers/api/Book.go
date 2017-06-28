package api

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"github.com/revel/revel"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"
)

type Book struct {
	*revel.Controller
}

func (c Book) Index() revel.Result {
	return c.RenderJSONP(c.Request.FormValue("callback"), models.NewOKApi())
}

//all books count
func (c Book) BooksCount() revel.Result {
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(bookService.QueryBooksCount()))
}

//all books info
func (c Book) Books() revel.Result {
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(bookService.QueryBooks(limit, start)),
	)
}

//book's download
func (c Book) BookDown() revel.Result {
	//bytes := rcali.FILE(bookService.QueryBookFile(bookid))
	formatid := rcali.ValueOrDefault(c.Request.FormValue("formatid"), "0")
	if ok,format := formatService.GetById(formatid);ok{
		if f, err := bookService.QueryBookFile(format.Id); err == nil {
			user, _ := userService.GetLoginUser(c.Request.FormValue("session"))
			book := bookService.QueryBook(format.CaliBook)
			if addOk := userService.AddDownload(user.Id, book.Id); addOk {

				// add status to sys status
				key := time.Now().Format("20060102") + "-downsize"
				if finfo, err := f.Stat(); err == nil {
					if status := sysStatusService.Get(key); status.Key != "" {
						value, _ := strconv.ParseInt(status.Value, 10, 0)
						value += finfo.Size()
						status.Value = strconv.FormatInt(value, 10)
						sysStatusService.UpdateStatus(status)

					} else {
						status = models.SysStatus{Key: key, Value: strconv.FormatInt(finfo.Size(), 10)}
						sysStatusService.AddSysStatus(status)
					}
				}

				return c.RenderFile(f, revel.Attachment)
			} else {
				return c.RenderText("database error")
			}
		}
	}

	return c.RenderText("file is not exit")
}

//query a book by bookid
func (c Book) Book() revel.Result {
	bookid:= rcali.ValueOrDefault(c.Request.FormValue("bookid"), "0")
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(bookService.QueryBook(bookid)),
	)
}

//query a book's info from //https://developers.douban.com/wiki/?title=book_v2#get_isbn_book by bookid by bookname
//func (c Book) DoubanBook() revel.Result {
//	bookid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("bookid"), "0"))
//	callback := c.Request.FormValue("callback")
//
//	bookVo := bookService.QueryBook(bookid)
//	rcali.Logger.Debug("https://api.douban.com/v2/book/search?q=" + bookVo.Title)
//	resp, err := http.Get("https://api.douban.com/v2/book/search?q=" + bookVo.Title)
//	if err != nil {
//		// handle error
//		return c.RenderJSONP(callback, models.NewErrorApi())
//	}
//
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		// handle error
//		return c.RenderJSONP(callback, models.NewErrorApi())
//	}
//	return c.RenderJSONP(callback, models.NewOKApiWithInfo(string(body)))
//}

//UPLOAD
func (c *Book) UploadBook() revel.Result {
	uploadpath, _ := rcali.GetUploadPath()
	file, header, err := c.Request.FormFile("book")
	if err == nil {
		defer file.Close()
		b, _ := ioutil.ReadAll(file)
		ioutil.WriteFile(path.Join(uploadpath, header.Filename), b, 0755)

		// add status to sys status
		key := time.Now().Format("20060102") + "-uploadsize"
		f, _ := os.Open(path.Join(uploadpath, header.Filename))
		defer f.Close()
		if finfo, err := f.Stat(); err == nil {
			if status := sysStatusService.Get(key); status.Key != "" {
				value, _ := strconv.ParseInt(status.Value, 10, 0)
				value += finfo.Size()
				status.Value = strconv.FormatInt(value, 10)
				sysStatusService.UpdateStatus(status)

			} else {
				status = models.SysStatus{Key: key, Value: strconv.FormatInt(finfo.Size(), 10)}
				sysStatusService.AddSysStatus(status)
			}
		}

		//ok := rcali.AddBook(path.Join(uploadpath, header.Filename))
		ok, bookid := bookService.UploadBookFormat(path.Join(uploadpath, header.Filename))
		if !ok {
			return c.RenderJSON(models.NewErrorApiWithInfo("add book error"))
		} else {
			user, _ := userService.GetLoginUser(c.Request.FormValue("session"))
			userService.AddUpload(user.Id, bookid)
			return c.RenderJSON(models.NewOKApiWithInfo("add book success"))
		}
	} else {
		rcali.Logger.Debug("read file error :", err.Error())
		return c.RenderJSON(models.NewErrorApiWithInfo(err))
	}
	return c.RenderJSON(models.NewOKApi())
}

func (c *Book) UploadBookConfirm() revel.Result {
	return c.RenderJSON(models.NewOKApi())
}


//SEARCH
func (c *Book) SearchCount() revel.Result {
	q, _ := url.QueryUnescape(c.Request.FormValue("q"))
	if q == "" {
		return c.RenderJSONP(c.Request.FormValue("callback"), models.NewErrorApi())
	} else {
		return c.RenderJSONP(c.Request.FormValue("callback"), models.NewOKApiWithInfo(bookService.SearchBooksCount(q)))
	}
}

func (c *Book) Search() revel.Result {
	q, _ := url.QueryUnescape(c.Request.FormValue("q"))
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	if q == "" {
		return c.RenderJSONP(c.Request.FormValue("callback"), models.NewErrorApi())
	} else {
		return c.RenderJSONP(c.Request.FormValue("callback"), models.NewOKApiWithInfo(bookService.SearchBooks(q, limit, start)))
	}
}
