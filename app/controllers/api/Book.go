package api

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"github.com/jiangmitiao/cali/app/services"
	"github.com/revel/revel"
	"io/ioutil"
	"net/http"
	"path"
)

type Book struct {
	*revel.Controller
}

func (c Book) Index(callback string) revel.Result {
	return c.RenderJSONP(callback, models.NewOKApi())
}

//all books count
func (c Book) BooksCount(callback string) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryBooksCount()))
}

//all books info
func (c Book) Books(callback string, limit, start int) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryBooks(limit, start)),
	)
}

//rating books info
func (c Book) RatingBooks(callback string, limit, start int) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryRatingBooks(limit, start)),
	)
}

//new books info
func (c Book) NewBooks(callback string, limit, start int) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryNewBooks(limit, start)),
	)
}

//discover books info
func (c Book) DiscoverBooks(callback string, limit, start int) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryDiscoverBooks(limit, start)),
	)
}

//tag books info
func (c Book) TagBooksCount(callback string, tagid int) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryTagBooksCount(tagid)),
	)
}

//tag books info
func (c Book) TagBooks(callback string, tagid, limit, start int) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryTagBooks(tagid, limit, start)),
	)
}

//author books info
func (c Book) AuthorBooksCount(callback string, authorid int) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryAuthorBooksCount(authorid)),
	)
}

//author books info
func (c Book) AuthorBooks(callback string, authorid, limit, start int) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryAuthorBooks(authorid, limit, start)),
	)
}

//language books info
func (c Book) LanguageBooksCount(callback string, lang_code int) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryLanguageBooksCount(lang_code)),
	)
}

//language books info
func (c Book) LanguageBooks(callback string, lang_code, limit, start int) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryLanguageBooks(lang_code, limit, start)),
	)
}

//book's rating
func (c Book) BookRating(callback string, bookid int) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryBookRating(bookid)),
	)
}

//book's img
func (c Book) BookImage(bookid int) revel.Result {
	bytes := rcali.IMGJPG(services.QueryCoverImg(bookid))
	return bytes
}

//book's download
func (c Book) BookDown(bookid int) revel.Result {
	//bytes := rcali.FILE(services.QueryBookFile(bookid))
	if f, err := services.QueryBookFile(bookid); err == nil {
		return c.RenderFile(f, revel.Attachment)
	}
	return c.RenderText("file is not exit")
}

func (c Book) Book(callback string, bookid int) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryBook(bookid)),
	)
}

func (c Book) DoubanBook(callback string, bookid int) revel.Result {
	bookVo := services.QueryBook(bookid)
	rcali.DEBUG.Debug("https://api.douban.com/v2/book/search?q=" + bookVo.Title)
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
		rcali.DEBUG.Debug("read file error :", err.Error())
		return c.RenderJSON(models.NewErrorApiWithInfo(err))
	}
	return c.RenderJSON(models.NewOKApi())
}
