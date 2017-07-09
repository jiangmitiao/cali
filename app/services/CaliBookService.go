package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"path"
	"sync"
	"time"
)

type CaliBookService struct {
	lock *sync.Mutex
}

//all books count
func (service CaliBookService) QueryBooksCount(categoryId string) int64 {
	count, _ := engine.Where("id in (select cali_book from cali_book_category where cali_category = ?)", categoryId).Count(models.CaliBook{})
	return count
}

//all books info
func (service CaliBookService) QueryBooks(limit, start int, categoryId, desc string) []models.CaliBook {
	books := make([]models.CaliBook, 0)
	engine.Where("id in (select cali_book from cali_book_category where cali_category = ?)", categoryId).Desc(desc).Limit(limit, start).Find(&books)
	return books
}

//find a book by bookid
func (service CaliBookService) GetById(bookId string) (has bool, book models.CaliBook) {
	has, _ = engine.Where("id = ?", bookId).Get(&book)
	return has,book
}

func (service CaliBookService) SearchBooksCount(searchStr, categoryId string) (count int) {
	engine.SQL("select count(1) from cali_book where id in (select cali_book from cali_book_category where cali_category = ?) and ( title like ? or author like ? )", categoryId, "%"+searchStr+"%", "%"+searchStr+"%").Get(&count)
	return
}

func (service CaliBookService) SearchBooks(limit, start int,searchStr, categoryId,desc string) []models.CaliBook {
	books := make([]models.CaliBook, 0)
	engine.SQL("select * from cali_book where id in (select cali_book from cali_book_category where cali_category = ?) and ( title like ? or author like ? ) limit ?,?", categoryId, "%"+searchStr+"%", "%"+searchStr+"%", start, limit).Desc(desc).Find(&books)

	return books
}

func (service CaliBookService) UploadBookFormat(filePath, tag string) (  models.CaliFormat ,bool,error) {
	service.lock.Lock()
	defer service.lock.Unlock()
	if eBook, ok := rcali.GetRealBookInfo(filePath); ok {
		if tmpFormat := DefaultFormatService.GetFormatBySize(eBook.UncompressedSize()); tmpFormat.Id == "" {
			if _, pathname := rcali.AddBook(filePath); pathname != "" {
				format := models.CaliFormat{
					Id:               uuid.New().String(),
					Format:           eBook.Format(),
					UncompressedSize: eBook.UncompressedSize(),
					FileName:         path.Base(pathname),
					Title:            eBook.Title(),
					Author:           eBook.Author(),
					Tag:              tag,
				}
				return format,DefaultFormatService.Add(format), nil
			}
		} else {
			rcali.Logger.Info("has the book id:" + tmpFormat.Id)
			return models.CaliFormat{},false, errors.New("hasthesamebook")
		}
	}

	return models.CaliFormat{},false, errors.New("unknownerror")
}

func (service CaliBookService) GetBookOrInsertByTitleAndAuthor(title, author string) (book models.CaliBook) {
	if ok, _ := engine.Where("title = ?", title).And("author = ?", author).Get(&book); ok {
		return
	} else {
		book.Id = uuid.New().String()
		book.Title = title
		book.Author = author
		book.DownloadCount = 0
		book.CreatedAt = time.Now().Unix()
		book.UpdatedAt = time.Now().Unix()
		engine.InsertOne(book)
		return
	}
}

func (service CaliBookService) GetBookByTitleAndAuthor(title, author string) (book models.CaliBook,err error) {
	 _,err=engine.Where("title = ?", title).And("author = ?", author).Get(&book)
	return
}

func (service CaliBookService) UpdateCaliBook(book models.CaliBook) {
	book.UpdatedAt = time.Now().Unix()
	engine.ID(book.Id).Cols("id", "title", "author", "download_count", "douban_id", "douban_json", "created", "updated").Update(book)
}

func (service CaliBookService) UpdateCaliBookDownload(book models.CaliBook) {
	book.UpdatedAt = time.Now().Unix()
	engine.ID(book.Id).Cols("download_count", "updated").Update(book)
}

func (service CaliBookService) AddBookCategory(bookId, categoryId string) {
	tmp := models.CaliBookCategory{}
	engine.Where("cali_book = ? ", bookId).And("cali_category = ?", categoryId).Get(tmp)
	if tmp.Id == "" {
		bc := models.CaliBookCategory{Id: uuid.New().String(), CaliBook: bookId, CaliCategory: categoryId, CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}
		engine.InsertOne(bc)
	}
}

func (service CaliBookService) DeleteById(bookId string) {
	engine.Where("id = ?", bookId).Delete(models.CaliBook{})
}
