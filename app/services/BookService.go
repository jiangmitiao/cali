package services

import (
	"errors"
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

//all books count
func QueryBooksCount() int64 {
	count, _ := engine.Count(models.Book{})
	return count
}

//all books info
func QueryBooks(limit, start int) []models.BookVo {
	bookVos := make([]models.BookVo, 0)
	engine.SQL("select books.* ,ratings.rating,authors.name  from books,ratings,books_ratings_link,authors,books_authors_link  where books.id=books_ratings_link.book and ratings.id=books_ratings_link.rating and books.id=books_authors_link.book and authors.id=books_authors_link.author limit " + strconv.Itoa(start) + " , " + strconv.Itoa(limit)).Find(&bookVos)
	return bookVos
}

//rating books info
func QueryRatingBooks(limit, start int) []models.BookVo {
	bookVos := make([]models.BookVo, 0)
	engine.SQL("select books.*,ratings.rating,authors.name  from books,ratings,books_ratings_link,authors,books_authors_link  where books.id=books_ratings_link.book and ratings.id=books_ratings_link.rating and books.id=books_authors_link.book and authors.id=books_authors_link.author order by ratings.rating desc limit " + strconv.Itoa(start) + " , " + strconv.Itoa(limit)).Find(&bookVos)
	return bookVos
}

//new books info
func QueryNewBooks(limit, start int) []models.BookVo {
	bookVos := make([]models.BookVo, 0)
	engine.SQL("select books.*,ratings.rating,authors.name  from books,ratings,books_ratings_link,authors,books_authors_link  where books.id=books_ratings_link.book and ratings.id=books_ratings_link.rating and books.id=books_authors_link.book and authors.id=books_authors_link.author order by books.timestamp desc limit " + strconv.Itoa(start) + " , " + strconv.Itoa(limit)).Find(&bookVos)
	return bookVos
}

//discover books info
func QueryDiscoverBooks(limit, start int) []models.BookVo {
	bookVos := make([]models.BookVo, 0)
	engine.SQL("select books.*,ratings.rating,authors.name  from books,ratings,books_ratings_link,authors,books_authors_link  where books.id=books_ratings_link.book and ratings.id=books_ratings_link.rating and books.id=books_authors_link.book and authors.id=books_authors_link.author order by books.title,books.last_modified desc limit " + strconv.Itoa(start) + " , " + strconv.Itoa(limit)).Find(&bookVos)
	return bookVos
}

//tag books info
func QueryTagBooksCount(tagid int) int {
	count := 0
	engine.SQL("select count(1) from books,ratings,books_ratings_link,authors,books_authors_link  where books.id=books_ratings_link.book and ratings.id=books_ratings_link.rating and books.id=books_authors_link.book and authors.id=books_authors_link.author and books.id in (select book from books_tags_link where tag=" + strconv.Itoa(tagid) + ") ").Get(&count)
	return count
}

//tag books info
func QueryTagBooks(tagid, limit, start int) []models.BookVo {
	bookVos := make([]models.BookVo, 0)
	engine.SQL("select books.*,ratings.rating,authors.name  from books,ratings,books_ratings_link,authors,books_authors_link  where books.id=books_ratings_link.book and ratings.id=books_ratings_link.rating and books.id=books_authors_link.book and authors.id=books_authors_link.author and books.id in (select book from books_tags_link where tag=" + strconv.Itoa(tagid) + ") limit " + strconv.Itoa(start) + " , " + strconv.Itoa(limit)).Find(&bookVos)
	return bookVos
}

//author books info
func QueryAuthorBooksCount(authorid int) int {
	count := 0
	engine.SQL("select count(1) from books,ratings,books_ratings_link,authors,books_authors_link  where books.id=books_ratings_link.book and ratings.id=books_ratings_link.rating and books.id=books_authors_link.book and authors.id=books_authors_link.author and books.id in (select book from books_authors_link where author=" + strconv.Itoa(authorid) + ") ").Get(&count)
	return count
}

//author books info
func QueryAuthorBooks(authorid, limit, start int) []models.BookVo {
	bookVos := make([]models.BookVo, 0)
	engine.SQL("select books.*,ratings.rating,authors.name  from books,ratings,books_ratings_link,authors,books_authors_link  where books.id=books_ratings_link.book and ratings.id=books_ratings_link.rating and books.id=books_authors_link.book and authors.id=books_authors_link.author and books.id in (select book from books_authors_link where author=" + strconv.Itoa(authorid) + ") limit " + strconv.Itoa(start) + " , " + strconv.Itoa(limit)).Find(&bookVos)
	return bookVos
}

//languages books info
func QueryLanguageBooksCount(lang_code int) int {
	count := 0
	engine.SQL("select count(1) from books,ratings,books_ratings_link,authors,books_authors_link  where books.id=books_ratings_link.book and ratings.id=books_ratings_link.rating and books.id=books_authors_link.book and authors.id=books_authors_link.author and books.id in (select book from books_languages_link where lang_code=" + strconv.Itoa(lang_code) + ") ").Get(&count)
	return count
}

//languages books info
func QueryLanguageBooks(lang_code, limit, start int) []models.BookVo {
	bookVos := make([]models.BookVo, 0)
	engine.SQL("select books.*,ratings.rating,authors.name  from books,ratings,books_ratings_link,authors,books_authors_link  where books.id=books_ratings_link.book and ratings.id=books_ratings_link.rating and books.id=books_authors_link.book and authors.id=books_authors_link.author and books.id in (select book from books_languages_link where lang_code=" + strconv.Itoa(lang_code) + ") limit " + strconv.Itoa(start) + " , " + strconv.Itoa(limit)).Find(&bookVos)
	return bookVos
}

//book's rating
func QueryBookRating(bookid int) int {
	bookrating := 0
	engine.SQL("select ratings.rating from books,ratings,books_ratings_link where books.id=books_ratings_link.book and ratings.id=books_ratings_link.rating and books.id=" + strconv.Itoa(bookid)).Get(&bookrating)
	return bookrating
}

//book's img
func QueryCoverImg(bookid int) []byte {
	book := models.Book{}
	engine.Where("id=?", bookid).Get(&book)
	if book.HasCover == 1 {
		if path, ok := rcali.GetBooksPath(); ok {
			//fmt.Println(path + book.Path + string(filepath.Separator) + "cover.jpg")
			bytes, _ := ioutil.ReadFile(path + book.Path + string(filepath.Separator) + "cover.jpg")
			return bytes
		}
	}
	return make([]byte, 0)
}

//book's file
func QueryBookFileByte(bookid int) []byte {
	data := models.Data{}
	engine.Where("book=?", bookid).Get(&data)
	book := models.Book{}
	engine.Where("id=?", bookid).Get(&book)
	if data.Format == "EPUB" {
		if path, ok := rcali.GetBooksPath(); ok {
			//fmt.Println(path + book.Path + string(filepath.Separator) + "cover.jpg")
			bytes, _ := ioutil.ReadFile(path + book.Path + string(filepath.Separator) + data.Name + ".epub")
			return bytes
		}
	}
	rcali.DEBUG.Debug(data, book)
	return make([]byte, 0)
}

//book's file
func QueryBookFile(bookid int) (*os.File, error) {
	data := models.Data{}
	engine.Where("book=?", bookid).Get(&data)
	book := models.Book{}
	engine.Where("id=?", bookid).Get(&book)
	if data.Format == "EPUB" {
		if path, ok := rcali.GetBooksPath(); ok {
			//fmt.Println(path + book.Path + string(filepath.Separator) + "cover.jpg")
			//bytes, _ := ioutil.ReadFile(path + book.Path + string(filepath.Separator) +data.Name+ ".epub")
			f, _ := os.Open(path + book.Path + string(filepath.Separator) + data.Name + ".epub")
			return f, nil
		}
	}
	rcali.DEBUG.Debug(data, book)
	return nil, errors.New("no exit")
}

func QueryBook(bookid int) models.BookVo {
	book := models.BookVo{}
	engine.SQL("select books.* ,ratings.rating,authors.name from books,ratings,books_ratings_link,authors,books_authors_link where books.id=books_ratings_link.book and ratings.id=books_ratings_link.rating and books.id=books_authors_link.book and authors.id=books_authors_link.author and books.id=" + strconv.Itoa(bookid)).Get(&book)
	return book
}
