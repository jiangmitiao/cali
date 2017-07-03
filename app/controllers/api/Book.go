package api

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"github.com/revel/revel"
	"net/url"
	"path"
	"path/filepath"
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
	categoryId := rcali.ValueOrDefault(c.Request.FormValue("categoryId"), models.DefaultCaliCategory.Id)
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(bookService.QueryBooksCount(categoryId)))
}

//all books info
func (c Book) Books() revel.Result {
	categoryid := rcali.ValueOrDefault(c.Request.FormValue("categoryId"), models.DefaultCaliCategory.Id)
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	more := c.Request.FormValue("more")

	books := bookService.QueryBooks(limit, start, categoryid)
	booksVos := make([]models.CaliBookVo, 0)
	if more != "" {
		for _, value := range books {
			bookvo := models.CaliBookVo{CaliBook: value}
			bookvo.Formats = formatService.QueryByBookId(bookvo.Id)
			bookvo.Categories = categoryService.QueryByBookIdWithOutDefault(bookvo.Id)
			booksVos = append(booksVos, bookvo)
		}
	} else {
		for _, value := range books {
			bookvo := models.CaliBookVo{CaliBook: value}
			booksVos = append(booksVos, bookvo)
		}
	}
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(booksVos),
	)
}

//book's download
func (c Book) BookDown() revel.Result {
	//bytes := rcali.FILE(bookService.QueryBookFile(bookid))
	formatId := rcali.ValueOrDefault(c.Request.FormValue("formatId"), "0")
	if ok, format := formatService.GetById(formatId); ok {
		if f, err := bookService.QueryBookFile(format.Id); err == nil {
			user, _ := userService.GetLoginUser(c.Request.FormValue("session"))
			c.addDownloadRecord(format, user)
			return c.RenderBinary(f, format.Title+"-"+format.Author+"."+format.Format, revel.Attachment, time.Unix(format.UpdatedAt, 0))
			//return c.RenderFile(f, revel.Attachment)
		}
	}
	return c.RenderText("file is not exit")
}

func (c Book) addDownloadRecord(format models.CaliFormat, user models.UserInfo) {
	// add status to sys status
	key := time.Now().Format("20060102") + "-downsize"
	if status := sysStatusService.Get(key); status.Key != "" {
		value, _ := strconv.ParseInt(status.Value, 10, 0)
		value += format.UncompressedSize
		status.Value = strconv.FormatInt(value, 10)
		sysStatusService.UpdateStatus(status)
	} else {
		status = models.SysStatus{Key: key, Value: strconv.FormatInt(format.UncompressedSize, 10)}
		sysStatusService.AddSysStatus(status)
	}

	//add books download count
	_, book := bookService.QueryBook(format.CaliBook)
	book.DownloadCount += 1
	bookService.UpdateCaliBookDownload(book)

	//add format download count
	format.DownloadCount += 1
	formatService.UpdateCaliFormatDownload(format)

	//user download
	userService.AddDownload(user.Id, format.Id)
}

//query a book by bookid
func (c Book) Book() revel.Result {
	bookId := rcali.ValueOrDefault(c.Request.FormValue("bookId"), "0")
	if has, book := bookService.QueryBook(bookId); has {
		bookvo := models.CaliBookVo{CaliBook: book}
		bookvo.Formats = formatService.QueryByBookId(bookvo.Id)
		bookvo.Categories = categoryService.QueryByBookIdWithOutDefault(bookvo.Id)
		return c.RenderJSONP(
			c.Request.FormValue("callback"),
			models.NewOKApiWithInfo(bookvo),
		)
	} else {
		return c.RenderJSONP(
			c.Request.FormValue("callback"),
			models.NewErrorApiWithMessageAndInfo(c.Message("nofindbook"), nil),
		)
	}
}

//query a book's info from //https://developers.douban.com/wiki/?title=book_v2#get_isbn_book by bookid by bookname

//UPLOAD
func (c *Book) UploadBook() revel.Result {
	uploadpath, _ := rcali.GetUploadPath()
	tag := rcali.ValueOrDefault(c.Request.FormValue("tag"), "")
	if file, header, err := c.Request.FormFile("book"); err == nil {
		tmpPath := path.Join(uploadpath, header.Filename)
		if rcali.WriteBook(file, tmpPath) == nil {
			if ok, format := bookService.UploadBookFormat(tmpPath, tag); ok {
				user, _ := userService.GetLoginUser(c.Request.FormValue("session"))
				c.addUploadRecord(format, user)
				return c.RenderJSON(models.NewOKApiWithMessageAndInfo("add book success", format))
			} else {
				return c.RenderJSON(models.NewErrorApiWithMessageAndInfo("add book error", nil))

			}
		} else {
			return c.RenderJSON(models.NewErrorApiWithMessageAndInfo("file upload error", nil))
		}
	} else {
		rcali.Logger.Debug("read file error :", err.Error())
		return c.RenderJSON(models.NewErrorApiWithMessageAndInfo(err.Error(), nil))
	}
	return c.RenderJSON(models.NewErrorApiWithMessageAndInfo("file read error", nil))
}

func (c Book) addUploadRecord(format models.CaliFormat, user models.UserInfo) {
	// add status to sys status
	key := time.Now().Format("20060102") + "-uploadsize"

	if status := sysStatusService.Get(key); status.Key != "" {
		value, _ := strconv.ParseInt(status.Value, 10, 0)
		value += format.UncompressedSize
		status.Value = strconv.FormatInt(value, 10)
		sysStatusService.UpdateStatus(status)
	} else {
		status = models.SysStatus{Key: key, Value: strconv.FormatInt(format.UncompressedSize, 10)}
		sysStatusService.AddSysStatus(status)
	}

	//user upload
	userService.AddUpload(user.Id, format.Id)
}

func (c *Book) UploadBookConfirm() revel.Result {
	//book
	book := bookService.GetBookOrInsertByTitleAndAuthor(rcali.ValueOrDefault(c.Request.FormValue("title"), ""), rcali.ValueOrDefault(c.Request.FormValue("author"), ""))
	book.DoubanId = rcali.ValueOrDefault(book.DoubanId, c.Request.FormValue("douban_id"))
	book.DoubanJson = rcali.GetDoubanInfoById(book.DoubanId)
	bookService.UpdateCaliBook(book)

	//category
	//categoryId := rcali.ValueOrDefault(c.Request.FormValue("categoryid"), models.DefaultCaliCategory.Id)
	//
	//bookService.AddBookCategory(book.Id, categoryId)
	bookService.AddBookCategory(book.Id, models.DefaultCaliCategory.Id)

	//format
	formatId := rcali.ValueOrDefault(c.Request.FormValue("formatId"), "")
	formatService.UpdateBookid(formatId, book.Id)

	return c.RenderJSON(models.NewOKApi())
}

//SEARCH
func (c *Book) SearchCount() revel.Result {
	q, _ := url.QueryUnescape(c.Request.FormValue("q"))
	q = rcali.ValueOrDefault(q, "")
	categoryId := rcali.ValueOrDefault(c.Request.FormValue("categoryId"), models.DefaultCaliCategory.Id)
	if q == "" {
		return c.RenderJSONP(c.Request.FormValue("callback"), models.NewErrorApi())
	} else {
		return c.RenderJSONP(c.Request.FormValue("callback"), models.NewOKApiWithInfo(bookService.SearchBooksCount(q, categoryId)))
	}
}

func (c *Book) Search() revel.Result {
	q, _ := url.QueryUnescape(c.Request.FormValue("q"))
	q = rcali.ValueOrDefault(q, "")
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	categoryId := rcali.ValueOrDefault(c.Request.FormValue("categoryId"), models.DefaultCaliCategory.Id)
	more := c.Request.FormValue("more")
	if q == "" {
		return c.RenderJSONP(c.Request.FormValue("callback"), models.NewErrorApi())
	} else {
		books := bookService.SearchBooks(q, categoryId, limit, start)

		booksVos := make([]models.CaliBookVo, 0)
		if more != "" {
			for _, value := range books {
				bookvo := models.CaliBookVo{CaliBook: value}
				bookvo.Formats = formatService.QueryByBookId(bookvo.Id)
				bookvo.Categories = categoryService.QueryByBookIdWithOutDefault(bookvo.Id)
				booksVos = append(booksVos, bookvo)
			}
		} else {
			for _, value := range books {
				bookvo := models.CaliBookVo{CaliBook: value}
				booksVos = append(booksVos, bookvo)
			}
		}
		return c.RenderJSONP(c.Request.FormValue("callback"), models.NewOKApiWithInfo(booksVos))
	}
}

//tmp del
func (c *Book) DelJSON() revel.Result {
	formats := formatService.GetNoBookLink()
	bookpathdir, _ := rcali.GetBooksPath()
	for _, format := range formats {
		rcali.DeleteRealBook(filepath.Join(bookpathdir, format.FileName))
		formatService.DeleteById(format.Id)
	}
	rcali.DeleteTmpBook()
	return c.RenderJSON("ok")
}

func (c *Book) Delete() revel.Result {
	bookId := rcali.ValueOrDefault(c.Request.FormValue("bookId"), "0")
	if has, book := bookService.QueryBook(bookId); has {
		bookService.DeleteById(book.Id)
		categoryService.DeleteBookCategoryByBookId(book.Id)
		formats := formatService.QueryByBookId(book.Id)
		for _, value := range formats {
			formatService.DeleteUserUploadDownload(value.Id)
		}
		formatService.DeleteByBookId(book.Id)
		return c.RenderJSON(models.NewOKApi())
	} else {
		return c.RenderJSON(models.NewErrorApi())
	}
}
func (c *Book) Update() revel.Result {
	bookId := rcali.ValueOrDefault(c.Request.FormValue("bookId"), "0")
	bookTitle := rcali.ValueOrDefault(c.Request.FormValue("bookTitle"), "0")
	bookAuthor := rcali.ValueOrDefault(c.Request.FormValue("bookAuthor"), "0")
	bookDoubanId := rcali.ValueOrDefault(c.Request.FormValue("bookDoubanId"), "")
	bookCategoryId := rcali.ValueOrDefault(c.Request.FormValue("bookCategoryId"), "0")
	if has, book := bookService.QueryBook(bookId); has {
		if newBook := bookService.GetBookByTitleAndAuthor(bookTitle, bookAuthor); newBook.Id != "" { //has
			formats := formatService.QueryByBookId(book.Id)
			for _, value := range formats {
				formatService.UpdateBookid(value.Id, newBook.Id)
			}
			categoryService.DeleteBookCategoryByBookId(bookId)
			categoryService.DeleteBookCategoryByBookId(newBook.Id)

			newBook.DoubanId = bookDoubanId
			newBook.DoubanJson = rcali.GetDoubanInfoById(bookDoubanId)
			newBook.DownloadCount += book.DownloadCount
			bookService.UpdateCaliBook(newBook)
			bookService.AddBookCategory(newBook.Id, models.DefaultCaliCategory.Id)
			bookService.AddBookCategory(newBook.Id, bookCategoryId)
			return c.RenderJSON(models.NewOKApi())
		} else {
			book.Title = bookTitle
			book.Author = bookAuthor
			book.DoubanId = bookDoubanId
			book.DoubanJson = rcali.GetDoubanInfoById(bookDoubanId)
			bookService.UpdateCaliBook(book)

			categoryService.DeleteBookCategoryByBookId(bookId)
			bookService.AddBookCategory(book.Id, models.DefaultCaliCategory.Id)
			bookService.AddBookCategory(book.Id, bookCategoryId)
			return c.RenderJSON(models.NewOKApi())
		}

		return c.RenderJSON(models.NewOKApi())
	} else {
		return c.RenderJSON(models.NewErrorApi())
	}
}
