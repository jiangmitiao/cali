package services

import (
	"errors"
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"github.com/google/uuid"
)

type CaliBookService struct {
	lock *sync.Mutex
}

//all books count
func (service CaliBookService) QueryBooksCount() int64 {
	count, _ := engine.Count(models.CaliBook{})
	return count
}

//all books info
func (service CaliBookService) QueryBooks(limit, start int) []models.CaliBookVo {
	books := make([]models.CaliBook, 0)
	engine.Limit(limit,start).Find(&books)

	bookVos :=make([]models.CaliBookVo,len(books))
	for i,value :=range books{
		bookVos[i].CaliBook = value
	}
	return bookVos
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
		bytes, _ := ioutil.ReadFile(path.Join(bookpath,format.FileName))
		return bytes
	}
	return make([]byte, 0)
}

//book's file
func (service CaliBookService) QueryBookFile(formatid string) (*os.File, error) {
	if ok,format := DefaultFormatService.GetById(formatid);ok{
		if bookpath, ok := rcali.GetBooksPath(); ok {
			f, err := os.Open(path.Join(bookpath,format.FileName))
			return f,err
		}
	}
	return nil, errors.New("no exit")
}

//find a book by bookid
func (service CaliBookService) QueryBook(bookid string) models.CaliBookVo {
	book := models.CaliBook{}
	engine.Where("id = ?", bookid).Get(&book)
	bookVo :=models.CaliBookVo{CaliBook:book}
	return bookVo
}

func (service CaliBookService) SearchBooksCount(searchStr string) int {
	count := 0
	engine.SQL("select count(1) from cali_book where title like ? or author like ?", "%"+searchStr+"%", "%"+searchStr+"%").Get(&count)
	return count
}

func (service CaliBookService) SearchBooks(searchStr string, limit, start int) []models.CaliBookVo {
	books := make([]models.CaliBook, 0)
	engine.SQL("select * from cali_book where title like ? or author like ? limit ?,?", "%"+searchStr+"%", "%"+searchStr+"%", start, limit).Find(&books)

	bookVos :=make([]models.CaliBookVo,len(books))
	for i,value :=range books{
		bookVos[i].CaliBook = value
	}
	return bookVos
}

func (service CaliBookService) UploadBookFormat(filePath string) (bool, models.CaliFormat) {
	service.lock.Lock()
	defer service.lock.Unlock()
	if ebook,pathname := rcali.AddBook(filePath); ebook!=nil {
		format := models.CaliFormat{
			Id:uuid.New().String(),
			Format:ebook.Format(),
			UncompressedSize:ebook.UncompressedSize(),
			FileName:path.Base(pathname),
		}
		return DefaultFormatService.Add(format),format
	}
	return false,models.CaliFormat{}
}
