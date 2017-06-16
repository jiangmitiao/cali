package api

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"github.com/jiangmitiao/cali/app/services"
	"github.com/revel/revel"
	"io/ioutil"
	"net/http"
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
		models.NewOKApiWithInfo(services.QueryBooksCount()))
}

//all books info
func (c Book) Books() revel.Result {
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(services.QueryBooks(limit, start)),
	)
}

//rating books info
func (c Book) RatingBooks() revel.Result {
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(services.QueryRatingBooks(limit, start)),
	)
}

//new books info
func (c Book) NewBooks() revel.Result {
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(services.QueryNewBooks(limit, start)),
	)
}

//discover books info
func (c Book) DiscoverBooks() revel.Result {
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(services.QueryDiscoverBooks(limit, start)),
	)
}

//tag books info
func (c Book) TagBooksCount() revel.Result {
	tagid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("tagid"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(services.QueryTagBooksCount(tagid)),
	)
}

//tag books info
func (c Book) TagBooks() revel.Result {
	tagid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("tagid"), "0"))
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(services.QueryTagBooks(tagid, limit, start)),
	)
}

//author books info
func (c Book) AuthorBooksCount() revel.Result {
	authorid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("authorid"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(services.QueryAuthorBooksCount(authorid)),
	)
}

//author books info
func (c Book) AuthorBooks() revel.Result {
	authorid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("authorid"), "0"))
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(services.QueryAuthorBooks(authorid, limit, start)),
	)
}

//language books info
func (c Book) LanguageBooksCount() revel.Result {
	lang_code, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("lang_code"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(services.QueryLanguageBooksCount(lang_code)),
	)
}

//language books info
func (c Book) LanguageBooks() revel.Result {
	lang_code, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("lang_code"), "0"))
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(services.QueryLanguageBooks(lang_code, limit, start)),
	)
}

//book's rating
func (c Book) BookRating() revel.Result {
	bookid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("bookid"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(services.QueryBookRating(bookid)),
	)
}

//book's img
func (c Book) BookImage() revel.Result {
	bookid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("bookid"), "0"))
	bytes := rcali.IMGJPG(services.QueryCoverImg(bookid))
	return bytes
}

//book's download
func (c Book) BookDown() revel.Result {
	//bytes := rcali.FILE(services.QueryBookFile(bookid))
	bookid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("bookid"), "0"))
	if f, err := services.QueryBookFile(bookid); err == nil {
		return c.RenderFile(f, revel.Attachment)
	}
	return c.RenderText("file is not exit")
}

//query a book by bookid
func (c Book) Book() revel.Result {
	bookid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("bookid"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(services.QueryBook(bookid)),
	)
}

//query a book's info from //https://developers.douban.com/wiki/?title=book_v2#get_isbn_book by bookid by bookname
func (c Book) DoubanBook() revel.Result {
	bookid, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("bookid"), "0"))
	callback := c.Request.FormValue("callback")

	bookVo := services.QueryBook(bookid)
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
		ok := rcali.AddBook(path.Join(uploadpath, header.Filename))
		if !ok {
			return c.RenderJSON(models.NewErrorApiWithInfo("add book error"))
		} else {
			return c.RenderJSON(models.NewOKApiWithInfo("add book success"))
		}
	} else {
		rcali.Logger.Debug("read file error :", err.Error())
		return c.RenderJSON(models.NewErrorApiWithInfo(err))
	}
	return c.RenderJSON(models.NewOKApi())
}
