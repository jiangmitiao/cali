package services

import (
	"github.com/jiangmitiao/cali/app/models"
	"time"
	"github.com/jiangmitiao/cali/app/rcali"
	"io/ioutil"
	"path"
	"os"
	"errors"
)

type CaliFormatService struct {
}

func (service CaliFormatService) Add(format models.CaliFormat) bool {
	format.CreatedAt = time.Now().Unix()
	format.UpdatedAt = time.Now().Unix()
	if _, err := engine.Insert(format); err == nil {
		return true
	} else {
		return false
	}
}

func (service CaliFormatService) GetById(formatid string) (ok bool, format models.CaliFormat) {
	ok, _ = engine.ID(formatid).Get(&format)
	return
}

func (service CaliFormatService) QueryByBookId(bookId string) []models.CaliFormat {
	formats := make([]models.CaliFormat, 0)
	engine.Where("cali_book = ?", bookId).Find(&formats)
	return formats
}

func (service CaliFormatService) UpdateBookId(formatId, bookId string) bool {
	return service.UpdateBookFormatCaliBook(models.CaliFormat{Id: formatId, CaliBook: bookId})
}

func (service CaliFormatService) UpdateTag(formatId, tag string) bool {
	if _, err := engine.ID(formatId).Cols("tag").Update(models.CaliFormat{Tag: tag}); err == nil {
		return true
	} else {
		return false
	}
}

func (service CaliFormatService) UpdateCaliFormatDownload(format models.CaliFormat) {
	format.UpdatedAt = time.Now().Unix()
	engine.ID(format.Id).Cols("download_count", "updated").Update(format)
}

func (service CaliFormatService) UpdateBookFormatCaliBook(format models.CaliFormat)bool {
	format.UpdatedAt = time.Now().Unix()
	if _,err :=engine.ID(format.Id).Cols("cali_book", "updated").Update(format);
		err==nil{
		return true
	}else {
		return false
	}
}

func (service CaliFormatService) GetFormatBySize(size int64) (format models.CaliFormat) {
	engine.Where("uncompressed_size = ? ", size).And("cali_book != ?", "").Get(&format)
	return
}

func (service CaliFormatService) GetNoBookLink() (formats []models.CaliFormat) {
	formats = make([]models.CaliFormat, 0)
	engine.Where("cali_book = ?", "").Find(&formats)
	return formats
}

func (service CaliFormatService) DeleteById(formatId string) {
	engine.Where("id = ?", formatId).Delete(models.CaliFormat{})
}

func (service CaliFormatService) DeleteByBookId(bookId string) {
	engine.Where("cali_book = ?", bookId).Delete(models.CaliFormat{})
}

func (service CaliFormatService) DeleteUserUploadDownload(formatId string)error {
	session := engine.NewSession()
	session.Begin()
	defer session.Close()
	if _,err :=session.Where("cali_format = ?", formatId).Delete(models.UserInfoBookUploadLink{});
	err==nil{
		if _,err:=engine.Where("cali_format = ?", formatId).Delete(models.UserInfoBookDownloadLink{});
		err==nil{
			session.Commit()
			return nil
		}else {
			session.Rollback()
			return err
		}
	}else {
		session.Rollback()
		return err
	}
}

//book's file
func (service CaliFormatService) QueryFormatFileByte(formatId string)([]byte,error) {
	if has,format :=DefaultFormatService.GetById(formatId);has{
		if bookpath, ok := rcali.GetBooksPath(); ok {
			return ioutil.ReadFile(path.Join(bookpath, format.FileName))
		}else {
			return make([]byte, 0),errors.New("books path is error")
		}
	}else {
		return make([]byte, 0),errors.New("not has this format")
	}
}

//book's file
func (service CaliFormatService) QueryFormatFile(formatId string) (*os.File, error) {
	if ok, format := DefaultFormatService.GetById(formatId); ok {
		if bookpath, ok := rcali.GetBooksPath(); ok {
			return os.Open(path.Join(bookpath, format.FileName))
		}else {
			return nil,errors.New("books path is error")
		}
	}else {
		return nil,errors.New("not has this format")
	}
}