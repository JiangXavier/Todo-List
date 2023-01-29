package model

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"todo_list/cache"
)

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

func (Task *Task) View() uint64 {
	//增加点击数
	countStr, _ := cache.RedisClient.Get(cache.TaskViewKey(Task.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

//AddView
func (Task *Task) AddView() {
	cache.RedisClient.Incr(cache.TaskViewKey(Task.ID))                      //增加视频点击数
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(Task.ID))) //增加排行点击数
}