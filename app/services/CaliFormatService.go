package services

import (
	"github.com/jiangmitiao/cali/app/models"
	"time"
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

func (service CaliFormatService) UpdateBookid(formatId, bookId string) bool {
	service.UpdateBookFormatCaliBook(models.CaliFormat{Id: formatId, CaliBook: bookId})
	return true
}

func (service CaliFormatService) UpdateTag(formatid, tag string) bool {
	if _, err := engine.ID(formatid).Cols("tag").Update(models.CaliFormat{Tag: tag}); err == nil {
		return true
	} else {
		return false
	}
}

func (service CaliFormatService) UpdateCaliFormatDownload(format models.CaliFormat) {
	format.UpdatedAt = time.Now().Unix()
	engine.ID(format.Id).Cols("download_count", "updated").Update(format)
}

func (service CaliFormatService) UpdateBookFormatCaliBook(format models.CaliFormat) {
	format.UpdatedAt = time.Now().Unix()
	engine.ID(format.Id).Cols("cali_book", "updated").Update(format)
}

func (service CaliFormatService) GetFormatBySize(size int64) (format models.CaliFormat) {
	engine.Where("uncompressed_size = ? ", size).And("cali_book != ?", "").Get(&format)
	return
}

func (service CaliFormatService)GetNoBookLink()(formats []models.CaliFormat)  {
	formats = make([]models.CaliFormat,0)
	engine.Where("cali_book = ?","").Find(&formats)
	return formats
}

func (service CaliFormatService) DeleteById(formatId string)  {
	engine.Where("id = ?",formatId).Delete(models.CaliFormat{})
}

func (service CaliFormatService) DeleteByBookId(bookId string)  {
	engine.Where("cali_book = ?",bookId).Delete(models.CaliFormat{})
}

func (service CaliFormatService) DeleteUserUploadDownload(formatId string)  {
	engine.Where("cali_format = ?",formatId).Delete(models.UserInfoBookUploadLink{})
	engine.Where("cali_format = ?",formatId).Delete(models.UserInfoBookDownloadLink{})
}
