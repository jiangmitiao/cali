package services

import "github.com/jiangmitiao/cali/app/models"

type CaliFormatService struct {

}

func (service CaliFormatService)Add (format models.CaliFormat)bool  {
	if _,err:=engine.Insert(format)err!=nil{
		return true
	}else {
		return false
	}
}

func (service CaliFormatService)GetById(formatid string)(ok bool,format models.CaliFormat)  {
	ok,_ =engine.ID(formatid).Get(&format)
	return
}

func (service CaliFormatService) UpdateBookid(formatid,bookid string)bool  {
	if _,err:= engine.ID(formatid).Cols("cali_book").Update(models.CaliFormat{CaliBook:bookid});
		err==nil{
		return true
	}else {
		return false
	}
}

func (service CaliFormatService) UpdateTag(formatid,tag string)bool  {
	if _,err:= engine.ID(formatid).Cols("tag").Update(models.CaliFormat{Tag:tag});
		err==nil{
		return true
	}else {
		return false
	}
}

