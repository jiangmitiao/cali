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

func (service CaliFormatService) QueryByCaliBook(bookid string)[]models.CaliFormat  {
	formats :=make([]models.CaliFormat,0)
	engine.Where("cali_book = ?",bookid).Find(&formats)
	return formats
}

func (service CaliFormatService) UpdateBookid(formatid, bookid string) bool {
	service.UpdateBookFormatCaliBook(models.CaliFormat{Id: formatid, CaliBook: bookid})
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
