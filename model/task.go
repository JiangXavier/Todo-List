package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	User User `gorm:"ForeignKey:Uid"`
	Uid uint `gorm:"not null"`
	Title string `gorm:"index;not null"`
	Status int `gorm:"default:0"` //0未完成 1已完成
	Content string`gorm:"type:longtext"`
	StartTime int64 //备忘录的开始时间
	EndTime int64 `gorm:"default:0"`//备忘录的完成时间
}
