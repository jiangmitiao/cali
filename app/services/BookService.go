package services

import (
	"errors"
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"image"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

type BookService struct {
	lock *sync.Mutex
}

//all books count
func (service BookService) QueryBooksCount() int64 {
	count, _ := engine.Count(models.Book{})
	return count
}

//all books info
func (service BookService) QueryBooks(limit, start int) []models.BookVo {
	bookVos := make([]models.BookVo, 0)
	engine.SQL("select books.* ,ratings.rating,authors.name  from books,authors,books_authors_link left join (select books_ratings_link.book,ratings.rating from ratings,books_ratings_link where ratings.id=books_ratings_link.rating) as ratings on books.id=ratings.book where books.id=books_authors_link.book and authors.id=books_authors_link.author limit ?,?", start, limit).Find(&bookVos)
	return bookVos
}

//rating books info
func (service BookService) QueryRatingBooks(limit, start int) []models.BookVo {
	bookVos := make([]models.BookVo, 0)
	engine.SQL("select books.*,ratings.rating,authors.name from books,authors,books_authors_link left join (select books_ratings_link.book,ratings.rating from ratings,books_ratings_link where ratings.id=books_ratings_link.rating) as ratings on books.id=ratings.book where books.id=books_authors_link.book and authors.id=books_authors_link.author order by ratings.rating desc limit ?,?", start, limit).Find(&bookVos)
	return bookVos
}

//new books info
func (service BookService) QueryNewBooks(limit, start int) []models.BookVo {
	bookVos := make([]models.BookVo, 0)
	engine.SQL("select books.*,ratings.rating,authors.name  from books,authors,books_authors_link left join (select books_ratings_link.book,ratings.rating from ratings,books_ratings_link where ratings.id=books_ratings_link.rating) as ratings on books.id=ratings.book  where books.id=books_authors_link.book and authors.id=books_authors_link.author order by books.timestamp desc limit ?,?", start, limit).Find(&bookVos)
	return bookVos
}

//discover books info
func (service BookService) QueryDiscoverBooks(limit, start int) []models.BookVo {
	bookVos := make([]models.BookVo, 0)
	engine.SQL("select books.*,ratings.rating,authors.name  from books,authors,books_authors_link left join (select books_ratings_link.book,ratings.rating from ratings,books_ratings_link where ratings.id=books_ratings_link.rating) as ratings on books.id=ratings.book where books.id=books_authors_link.book and authors.id=books_authors_link.author order by books.title,books.last_modified desc limit ?,?", start, limit).Find(&bookVos)
	return bookVos
}

//tag books info
func (service BookService) QueryTagBooksCount(tagid int) int {
	count := 0
	engine.SQL("select count(1) from books,authors,books_authors_link left join (select books_ratings_link.book,ratings.rating from ratings,books_ratings_link where ratings.id=books_ratings_link.rating) as ratings on books.id=ratings.book where books.id=books_authors_link.book and authors.id=books_authors_link.author and books.id in (select book from books_tags_link where tag= ? ) ", tagid).Get(&count)
	return count
}

//tag books info
func (service BookService) QueryTagBooks(tagid, limit, start int) []models.BookVo {
	bookVos := make([]models.BookVo, 0)
	engine.SQL("select books.*,ratings.rating,authors.name  from books,authors,books_authors_link left join (select books_ratings_link.book,ratings.rating from ratings,books_ratings_link where ratings.id=books_ratings_link.rating) as ratings on books.id=ratings.book where books.id=books_authors_link.book and authors.id=books_authors_link.author and books.id in (select book from books_tags_link where tag= ? ) limit ?,?", tagid, start, limit).Find(&bookVos)
	return bookVos
}

//author books info
func (service BookService) QueryAuthorBooksCount(authorid int) int {
	count := 0
	engine.SQL("select count(1) from books,authors,books_authors_link left join (select books_ratings_link.book,ratings.rating from ratings,books_ratings_link where ratings.id=books_ratings_link.rating) as ratings on books.id=ratings.book where books.id=books_authors_link.book and authors.id=books_authors_link.author and books.id in (select book from books_authors_link where author= ? ) ", authorid).Get(&count)
	return count
}

//author books info
func (service BookService) QueryAuthorBooks(authorid, limit, start int) []models.BookVo {
	bookVos := make([]models.BookVo, 0)
	engine.SQL("select books.*,ratings.rating,authors.name from books,authors,books_authors_link left join (select books_ratings_link.book,ratings.rating from ratings,books_ratings_link where ratings.id=books_ratings_link.rating) as ratings on books.id=ratings.book where books.id=books_authors_link.book and authors.id=books_authors_link.author and books.id in (select book from books_authors_link where author= ? ) limit ?,?", authorid, start, limit).Find(&bookVos)
	return bookVos
}

//languages books info
func (service BookService) QueryLanguageBooksCount(lang_code int) int {
	count := 0
	engine.SQL("select count(1) from books,authors,books_authors_link left join (select books_ratings_link.book,ratings.rating from ratings,books_ratings_link where ratings.id=books_ratings_link.rating) as ratings on books.id=ratings.book where books.id=books_authors_link.book and authors.id=books_authors_link.author and books.id in (select book from books_languages_link where lang_code= ? ) ", lang_code).Get(&count)
	return count
}

//languages books info
func (service BookService) QueryLanguageBooks(lang_code, limit, start int) []models.BookVo {
	bookVos := make([]models.BookVo, 0)
	engine.SQL("select books.*,ratings.rating,authors.name  from books,authors,books_authors_link left join (select books_ratings_link.book,ratings.rating from ratings,books_ratings_link where ratings.id=books_ratings_link.rating) as ratings on books.id=ratings.book where books.id=books_authors_link.book and authors.id=books_authors_link.author and books.id in (select book from books_languages_link where lang_code= ? ) limit ?,?", lang_code, start, limit).Find(&bookVos)
	return bookVos
}

//book's rating
func (service BookService) QueryBookRating(bookid int) int {
	bookrating := 0
	engine.SQL("select ratings.rating from books,ratings,books_ratings_link where books.id=books_ratings_link.book and ratings.id=books_ratings_link.rating and books.id= ? ", bookid).Get(&bookrating)
	return bookrating
}

//book's img
func (service BookService) QueryCoverImg(bookid int) []byte {
	book := models.Book{}
	engine.Where("id=?", bookid).Get(&book)
	if book.HasCover == 1 {
		if basepath, ok := rcali.GetBooksPath(); ok {
			//fmt.Println(path + book.Path + string(filepath.Separator) + "cover.jpg")
			//bytes, _ := ioutil.ReadFile(path + book.Path + string(filepath.Separator) + "cover.jpg")
			pathtmp := path.Join(basepath, book.Path, "cover.jpg")
			if f, err := os.Open(pathtmp); err == nil {
				img, _, _ := image.Decode(f)
				result := rcali.JpegImage2Bytes(rcali.ResizeImage(200, 300, img))
				return result
			}
		}
	}
	return rcali.JpegImage2Bytes(rcali.EmptyIamge(200, 300))
}

//book's file
func (service BookService) QueryBookFileByte(bookid int) []byte {
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
	rcali.Logger.Debug(data, book)
	return make([]byte, 0)
}

//book's file
func (service BookService) QueryBookFile(bookid int) (*os.File, error) {
	data := models.Data{}
	engine.Where("book=?", bookid).Get(&data)
	book := models.Book{}
	engine.Where("id=?", bookid).Get(&book)
	if data.Format != "" {
		if path, ok := rcali.GetBooksPath(); ok {
			//fmt.Println(path + book.Path + string(filepath.Separator) + "cover.jpg")
			//bytes, _ := ioutil.ReadFile(path + book.Path + string(filepath.Separator) +data.Name+ ".epub")
			f, _ := os.Open(path + book.Path + string(filepath.Separator) + data.Name + "." + strings.ToLower(data.Format))
			return f, nil
		}
	}
	rcali.Logger.Debug(data, book)
	return nil, errors.New("no exit")
}

//find a book by bookid
func (service BookService) QueryBook(bookid int) models.BookVo {
	book := models.BookVo{}
	engine.SQL("select books.* ,ratings.rating,authors.name,comments.comments from books,authors,books_authors_link left join (select books_ratings_link.book,ratings.rating from ratings,books_ratings_link where ratings.id=books_ratings_link.rating) as ratings on books.id=ratings.book left join (select book,text as comments from comments) as comments on comments.book=books.id where books.id=books_authors_link.book and authors.id=books_authors_link.author and books.id= ? ", bookid).Get(&book)
	data := models.Data{}
	engine.Where("book=?", bookid).Get(&data)
	book.Format = data.Format
	return book
}

func (service BookService) SearchBooksCount(searchStr string) int {
	count := 0
	engine.SQL("select count(1)  from books,authors,books_authors_link left join (select books_ratings_link.book,ratings.rating from ratings,books_ratings_link where ratings.id=books_ratings_link.rating) as ratings on books.id=ratings.book where books.id=books_authors_link.book and authors.id=books_authors_link.author and (books.title like ? or authors.name like ?) ", "%"+searchStr+"%", "%"+searchStr+"%").Get(&count)
	return count
}

func (service BookService) SearchBooks(searchStr string, limit, start int) []models.BookVo {
	bookVos := make([]models.BookVo, 0)
	engine.SQL("select books.* ,ratings.rating,authors.name  from books,authors,books_authors_link left join (select books_ratings_link.book,ratings.rating from ratings,books_ratings_link where ratings.id=books_ratings_link.rating) as ratings on books.id=ratings.book where books.id=books_authors_link.book and authors.id=books_authors_link.author and (books.title like ? or authors.name like ?) limit ?,?", "%"+searchStr+"%", "%"+searchStr+"%", start, limit).Find(&bookVos)
	return bookVos
}

func (service BookService) UploadBook(filePath string) (uploadOk bool, bookid int) {
	service.lock.Lock()
	defer service.lock.Unlock()
	uploadOk = false
	bookid = 0
	oldBookId := 0
	if _, err := localEngine.SQL("select seq from sqlite_sequence where name='books'").Get(&oldBookId); err == nil {
		if uploadOk = rcali.AddBook(filePath); uploadOk {
			if _, err := localEngine.SQL("select seq from sqlite_sequence where name='books'").Get(&bookid); err == nil && bookid == oldBookId+1 {
				return
			} else {
				return false, 0
			}
		}
	}
	return
}
