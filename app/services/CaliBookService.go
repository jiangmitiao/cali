package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"time"
)

type CaliBookService struct {
	lock *sync.Mutex
}

//all books count
func (service CaliBookService) QueryBooksCount(categoryid string) int64 {
	count, _ := engine.Where("id in (select cali_book from cali_book_category where cali_category = ?)", categoryid).Count(models.CaliBook{})
	return count
}

//all books info
func (service CaliBookService) QueryBooks(limit, start int, categoryid string) []models.CaliBook {
	books := make([]models.CaliBook, 0)
	engine.Where("id in (select cali_book from cali_book_category where cali_category = ?)", categoryid).Desc("updated").Limit(limit, start).Find(&books)
	return books
}

//book's img
//func (service CaliBookService) QueryCoverImg(bookid int) []byte {
//	book := models.Book{}
//	engine.Where("id=?", bookid).Get(&book)
//	if book.HasCover == 1 {
//		if basepath, ok := rcali.GetBooksPath(); ok {
//			//fmt.Println(path + book.Path + string(filepath.Separator) + "cover.jpg")
//			//bytes, _ := ioutil.ReadFile(path + book.Path + string(filepath.Separator) + "cover.jpg")
//			pathtmp := path.Join(basepath, book.Path, "cover.jpg")
//			if f, err := os.Open(pathtmp); err == nil {
//				img, _, _ := image.Decode(f)
//				result := rcali.JpegImage2Bytes(rcali.ResizeImage(200, 300, img))
//				return result
//			}
//		}
//	}
//	return rcali.JpegImage2Bytes(rcali.EmptyIamge(200, 300))
//}

//book's file
func (service CaliBookService) QueryBookFileByte(formatid string) []byte {
	format := models.CaliFormat{}
	engine.Where("id=?", formatid).Get(&format)
	if bookpath, ok := rcali.GetBooksPath(); ok {
		bytes, _ := ioutil.ReadFile(path.Join(bookpath, format.FileName))
		return bytes
	}
	return make([]byte, 0)
}

//book's file
func (service CaliBookService) QueryBookFile(formatid string) (*os.File, error) {
	if ok, format := DefaultFormatService.GetById(formatid); ok {
		if bookpath, ok := rcali.GetBooksPath(); ok {
			f, err := os.Open(path.Join(bookpath, format.FileName))
			return f, err
		}
	}
	return nil, errors.New("no exit")
}

//find a book by bookid
func (service CaliBookService) QueryBook(bookid string) (has bool, book models.CaliBook) {
	if has, _ := engine.Where("id = ?", bookid).Get(&book); has {
		return true, book
	} else {
		return false, book
	}
}

func (service CaliBookService) SearchBooksCount(searchStr, categoryId string) int {
	count := 0
	engine.SQL("select count(1) from cali_book where id in (select cali_book from cali_book_category where cali_category = ?) and ( title like ? or author like ? )", categoryId, "%"+searchStr+"%", "%"+searchStr+"%").Get(&count)
	return count
}

func (service CaliBookService) SearchBooks(searchStr, categoryId string, limit, start int) []models.CaliBook {
	books := make([]models.CaliBook, 0)
	engine.SQL("select * from cali_book where id in (select cali_book from cali_book_category where cali_category = ?) and ( title like ? or author like ? ) limit ?,?", categoryId, "%"+searchStr+"%", "%"+searchStr+"%", start, limit).Desc("updated").Find(&books)

	return books
}

func (service CaliBookService) UploadBookFormat(filePath, tag string) (bool, models.CaliFormat) {
	service.lock.Lock()
	defer service.lock.Unlock()
	if ebook, ok := rcali.GetRealBookInfo(filePath); ok {
		if tmpFormat := DefaultFormatService.GetFormatBySize(ebook.UncompressedSize()); tmpFormat.Id == "" {
			if _, pathname := rcali.AddBook(filePath); pathname != "" {
				format := models.CaliFormat{
					Id:               uuid.New().String(),
					Format:           ebook.Format(),
					UncompressedSize: ebook.UncompressedSize(),
					FileName:         path.Base(pathname),
					Title:            ebook.Title(),
					Author:           ebook.Author(),
					Tag:              tag,
				}
				return DefaultFormatService.Add(format), format
			}
		} else {
			rcali.Logger.Info("has the book id:" + tmpFormat.Id)
		}
	}

	return false, models.CaliFormat{}
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

func (service CaliBookService) GetBookByTitleAndAuthor(title, author string) (book models.CaliBook) {
	engine.Where("title = ?", title).And("author = ?", author).Get(&book)
	return
}

func (service CaliBookService) UpdateCaliBook(book models.CaliBook) {
	book.UpdatedAt = time.Now().Unix()
	engine.ID(book.Id).Cols("id","title","author","download_count","douban_id","douban_json","created","updated").Update(book)
}

func (service CaliBookService) UpdateCaliBookDownload(book models.CaliBook) {
	book.UpdatedAt = time.Now().Unix()
	engine.ID(book.Id).Cols("download_count", "updated").Update(book)
}

func (service CaliBookService) AddBookCategory(bookid, categoryid string) {
	tmp := models.CaliBookCategory{}
	engine.Where("cali_book ? ", bookid).And("cali_category = ?", categoryid).Get(tmp)
	if tmp.Id == "" {
		bc := models.CaliBookCategory{Id: uuid.New().String(), CaliBook: bookid, CaliCategory: categoryid, CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}
		engine.InsertOne(bc)
	}
}

func (service CaliBookService) DeleteById(bookId string) {
	engine.Where("id = ?", bookId).Delete(models.CaliBook{})
}
