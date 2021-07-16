package model

import (
	"errors"
	"gorm.io/gorm"
)

//我的评分表
type Score struct {
	gorm.Model
	//品牌
	CarBand string `gorm:"not null"`
	//车系
	CarSeries string `gorm:"not null"`
	//分数
	Score float32 `gorm:"not null"`
	//评分人id
	UserID uint `gorm:"not null;column:userid"`
}

//avg
type CarAvg struct {
	//品牌
	CarBand string
	//车系
	CarSeries string
	//平均分
	AvgScore float32
}

//打分
func (s *Score)MakeScore() ( score *Score,err error) {

	if err=db.Table("scores").Create(s).Error;err!=nil{
		return &Score{}, err
	}
	return s,nil
}
//修改评分
func (s *Score)ModifyScore() ( score *Score,err error) {
	var temp =new(Score)
	tx:=db.Begin()
	if tx.Table("scores").Where("car_band = ? AND car_series = ?",s.CarBand,s.CarSeries).First(temp).RowsAffected >0 {
		tx.Rollback()
		return &Score{}, errors.New("已经打过分了")
	}
	err=tx.Table("scores").Where("car_band = ? AND car_series = ?",s.CarBand,s.CarSeries).Update("score",s.Score).Error
	if err != nil {
		tx.Rollback()
		return &Score{}, err
	}
	tx.Commit()
	return s, nil
}

//查看我的评分
func FindMYScore(userid int32)(list *[]Score,err error)  {
	err=db.Table("scores").Where("userid = ?",userid).Find(list).Error
	if err != nil {
		return new([]Score), err
	}
	return list, nil
}


//todo 计算平均分
//计算平均分
func GetAVRScore() (cv *CarAvg,err error) {
	err=db.Raw("SELECT car_band,car_series,AVG(score) AS AvgScore FROM scores GROUP BY car_band,car_series").Scan(cv).Error
	if err != nil {
		return new(CarAvg), err
	}
	return cv, nil
}