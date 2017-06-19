package api

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"github.com/revel/revel"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
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

//rating books info
func (c Book) RatingBooks() revel.Result {
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(bookService.QueryRatingBooks(limit, start)),
	)
}

//new books info
func (c Book) NewBooks() revel.Result {
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(bookService.QueryNewBooks(limit, start)),
	)
}

//discover books info
func (c Book) DiscoverBooks() revel.Result {
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(bookService.QueryDiscoverBooks(limit, start)),
	)
}

//tag books info
func (c Book) TagBooksCount() revel.Result {
	tagid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("tagid"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(bookService.QueryTagBooksCount(tagid)),
	)
}

//tag books info
func (c Book) TagBooks() revel.Result {
	tagid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("tagid"), "0"))
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(bookService.QueryTagBooks(tagid, limit, start)),
	)
}

//author books info
func (c Book) AuthorBooksCount() revel.Result {
	authorid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("authorid"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(bookService.QueryAuthorBooksCount(authorid)),
	)
}

//author books info
func (c Book) AuthorBooks() revel.Result {
	authorid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("authorid"), "0"))
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(bookService.QueryAuthorBooks(authorid, limit, start)),
	)
}

//language books info
func (c Book) LanguageBooksCount() revel.Result {
	lang_code, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("lang_code"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(bookService.QueryLanguageBooksCount(lang_code)),
	)
}

//language books info
func (c Book) LanguageBooks() revel.Result {
	lang_code, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("lang_code"), "0"))
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(bookService.QueryLanguageBooks(lang_code, limit, start)),
	)
}

//book's rating
func (c Book) BookRating() revel.Result {
	bookid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("bookid"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(bookService.QueryBookRating(bookid)),
	)
}

//book's img
func (c Book) BookImage() revel.Result {
	bookid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("bookid"), "0"))
	bytes := rcali.IMGJPG(bookService.QueryCoverImg(bookid))
	return bytes
}

//book's download
func (c Book) BookDown() revel.Result {
	//bytes := rcali.FILE(bookService.QueryBookFile(bookid))
	bookid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("bookid"), "0"))
	if f, err := bookService.QueryBookFile(bookid); err == nil {
		user, _ := userService.GetLoginUser(c.Request.FormValue("session"))
		if addOk := userService.AddDownload(user.Id, bookid); addOk {
			return c.RenderFile(f, revel.Attachment)
		} else {
			return c.RenderText("database error")
		}
	}
	return c.RenderText("file is not exit")
}

//query a book by bookid
func (c Book) Book() revel.Result {
	bookid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("bookid"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(bookService.QueryBook(bookid)),
	)
}

//query a book's info from //https://developers.douban.com/wiki/?title=book_v2#get_isbn_book by bookid by bookname
func (c Book) DoubanBook() revel.Result {
	bookid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("bookid"), "0"))
	callback := c.Request.FormValue("callback")

	bookVo := bookService.QueryBook(bookid)
	rcali.Logger.Debug("https://api.douban.com/v2/book/search?q=" + bookVo.Title)
	resp, err := http.Get("https://api.douban.com/v2/book/search?q=" + bookVo.Title)
	if err != nil {
		// handle error
		return c.RenderJSONP(callback, models.NewErrorApi())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return c.RenderJSONP(callback, models.NewErrorApi())
	}
	return c.RenderJSONP(callback, models.NewOKApiWithInfo(string(body)))
}

func (c *Book) UploadBook() revel.Result {
	uploadpath, _ := rcali.GetUploadPath()
	file, header, err := c.Request.FormFile("book")
	if err == nil {
		defer file.Close()
		b, _ := ioutil.ReadAll(file)
		ioutil.WriteFile(path.Join(uploadpath, header.Filename), b, 0755)
		//ok := rcali.AddBook(path.Join(uploadpath, header.Filename))
		ok, bookid := bookService.UploadBook(path.Join(uploadpath, header.Filename))
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
